/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"magma/orc8r/cloud/go/datastore"
	"magma/orc8r/cloud/go/sqorc"
	"magma/orc8r/cloud/go/storage"

	sq "github.com/Masterminds/squirrel"
	"github.com/thoas/go-funk"
)

type sqlConfigStorage struct {
	db      *sql.DB
	builder sqorc.StatementBuilder
}

func NewSqlConfigurationStorage(db *sql.DB, sqlBuilder sqorc.StatementBuilder) ConfigurationStorage {
	return &sqlConfigStorage{db: db, builder: sqlBuilder}
}

const (
	tableName = "configurations"
	typeCol   = "type"
	keyCol    = "\"key\""
	valCol    = "value"
	verCol    = "version"
)

func GetConfigTableName(networkId string) string {
	return datastore.GetTableName(networkId, tableName)
}

// This is a mega-hack before our inter-service streaming architecture is in
// place. Execute a CREATE TABLE IF NOT EXISTS on every query so we don't
// query a non-existent table.
func (store *sqlConfigStorage) initTable(tx *sql.Tx, table string) error {
	_, err := store.builder.CreateTable(table).
		IfNotExists().
		Column(typeCol).Type(sqorc.ColumnTypeText).NotNull().EndColumn().
		Column(keyCol).Type(sqorc.ColumnTypeText).NotNull().EndColumn().
		Column(valCol).Type(sqorc.ColumnTypeBytes).EndColumn().
		Column(verCol).Type(sqorc.ColumnTypeInt).NotNull().Default(0).EndColumn().
		PrimaryKey(typeCol, keyCol).
		RunWith(tx).
		Exec()
	return err
}

func (store *sqlConfigStorage) GetConfig(networkId string, configType string, key string) (*ConfigValue, error) {
	txFn := func(tx *sql.Tx) (interface{}, error) {
		value, version, err := store.getConfig(tx, networkId, configType, key)
		if err == sql.ErrNoRows {
			return &ConfigValue{}, nil
		}
		if err != nil {
			return nil, err
		}
		return &ConfigValue{Value: value, Version: version}, nil
	}

	ret, err := sqorc.ExecInTx(store.db, store.getInitFn(networkId), txFn)
	if err != nil {
		return nil, err
	}
	return ret.(*ConfigValue), err
}

func (store *sqlConfigStorage) GetConfigs(networkId string, criteria *FilterCriteria) (map[storage.TypeAndKey]*ConfigValue, error) {
	err := validateFilterCriteria(criteria)
	if err != nil {
		return nil, err
	}

	txFn := func(tx *sql.Tx) (interface{}, error) {
		tableName := GetConfigTableName(networkId)

		rows, err := store.builder.Select(typeCol, keyCol, valCol, verCol).
			From(tableName).
			Where(getWhereConditionFromFilterCriteria(criteria)).
			RunWith(tx).
			Query()
		if err != nil {
			return nil, err
		}
		defer sqorc.CloseRowsLogOnError(rows, "GetConfigs")

		scannedRows := map[storage.TypeAndKey]*ConfigValue{}
		for rows.Next() {
			var configType, key string
			var value []byte
			var version uint64

			err = rows.Scan(&configType, &key, &value, &version)
			if err != nil {
				return nil, err
			}
			scannedRows[storage.TypeAndKey{Type: configType, Key: key}] = &ConfigValue{Value: value, Version: version}
		}
		return scannedRows, nil
	}

	ret, err := sqorc.ExecInTx(store.db, store.getInitFn(networkId), txFn)
	if err != nil {
		return nil, err
	}
	return ret.(map[storage.TypeAndKey]*ConfigValue), err
}

func (store *sqlConfigStorage) ListKeysForType(networkId string, configType string) ([]string, error) {
	txFn := func(tx *sql.Tx) (interface{}, error) {
		tableName := GetConfigTableName(networkId)
		rows, err := store.builder.Select(keyCol).
			From(tableName).
			Where(sq.Eq{typeCol: configType}).
			RunWith(tx).
			Query()
		if err != nil {
			return nil, err
		}
		defer sqorc.CloseRowsLogOnError(rows, "ListKeysForType")

		scannedRows := make([]string, 0)
		for rows.Next() {
			var key string
			err = rows.Scan(&key)
			if err != nil {
				return nil, err
			}
			scannedRows = append(scannedRows, key)
		}
		return scannedRows, nil
	}

	ret, err := sqorc.ExecInTx(store.db, store.getInitFn(networkId), txFn)
	if err != nil {
		return nil, err
	}
	return ret.([]string), err
}

func (store *sqlConfigStorage) CreateConfig(networkId string, configType string, key string, value []byte) error {
	txFn := func(tx *sql.Tx) (interface{}, error) {
		// Check for existing record
		tableName := GetConfigTableName(networkId)
		exists, err := store.doesConfigExist(tx, networkId, configType, key)
		if err != nil {
			return nil, err
		}
		if exists {
			err = fmt.Errorf("Creating already existing config with type %s and key %s", configType, key)
			return nil, err
		}

		_, err = store.builder.Insert(tableName).
			Columns(typeCol, keyCol, valCol).
			Values(configType, key, value).
			RunWith(tx).
			Exec()
		return nil, err
	}

	_, err := sqorc.ExecInTx(store.db, store.getInitFn(networkId), txFn)
	return err
}

func (store *sqlConfigStorage) UpdateConfig(networkId string, configType string, key string, newValue []byte) error {
	txFn := func(tx *sql.Tx) (interface{}, error) {
		tableName := GetConfigTableName(networkId)
		// Get current generation number and update row
		_, version, err := store.getConfig(tx, networkId, configType, key)
		if err != nil {
			if err == sql.ErrNoRows {
				err = fmt.Errorf("Updating nonexistent config with type %s and key %s", configType, key)
			}
			return nil, err
		}

		_, err = store.builder.Update(tableName).
			Set(valCol, newValue).
			Set(verCol, version+1).
			Where(getExactWhereCondition(configType, key)).
			RunWith(tx).
			Exec()
		return nil, err
	}

	_, err := sqorc.ExecInTx(store.db, store.getInitFn(networkId), txFn)
	return err
}

func (store *sqlConfigStorage) DeleteConfig(networkId string, configType string, key string) error {
	txFn := func(tx *sql.Tx) (interface{}, error) {
		tableName := GetConfigTableName(networkId)
		exists, err := store.doesConfigExist(tx, networkId, configType, key)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, fmt.Errorf("Deleting nonexistent config with type %s and key %s", configType, key)
		}
		_, err = store.builder.Delete(tableName).
			Where(getExactWhereCondition(configType, key)).
			RunWith(tx).
			Exec()
		return nil, err
	}

	_, err := sqorc.ExecInTx(store.db, store.getInitFn(networkId), txFn)
	return err
}

func (store *sqlConfigStorage) DeleteConfigs(networkId string, criteria *FilterCriteria) error {
	err := validateFilterCriteria(criteria)
	if err != nil {
		return err
	}

	txFn := func(tx *sql.Tx) (interface{}, error) {
		_, err := store.builder.Delete(GetConfigTableName(networkId)).
			Where(getWhereConditionFromFilterCriteria(criteria)).
			RunWith(tx).
			Exec()
		return nil, err
	}

	_, err = sqorc.ExecInTx(store.db, store.getInitFn(networkId), txFn)
	return err
}

func (store *sqlConfigStorage) DeleteConfigsForNetwork(networkId string) error {
	txFn := func(tx *sql.Tx) (interface{}, error) {
		tableName := GetConfigTableName(networkId)
		_, err := tx.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
		return nil, err
	}

	_, err := sqorc.ExecInTx(store.db, func(*sql.Tx) error { return nil }, txFn)
	return err
}

func (store *sqlConfigStorage) doesConfigExist(tx *sql.Tx, networkId string, configType string, key string) (bool, error) {
	tableName := GetConfigTableName(networkId)

	var result uint64
	err := store.builder.Select("1").From(tableName).
		Where(sq.And{sq.Eq{typeCol: configType}, sq.Eq{keyCol: key}}).
		RunWith(tx).
		QueryRow().Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, err
	}
}

func (store *sqlConfigStorage) getConfig(tx *sql.Tx, networkId string, configType string, key string) ([]byte, uint64, error) {
	tableName := GetConfigTableName(networkId)
	var value []byte
	var version uint64

	// Explicit sq.And to preserve ordering of Eq clauses
	err := store.builder.Select(valCol, verCol).
		From(tableName).
		Where(sq.And{sq.Eq{typeCol: configType}, sq.Eq{keyCol: key}}).
		RunWith(tx).
		QueryRow().Scan(&value, &version)
	return value, version, err
}

func (store *sqlConfigStorage) getInitFn(networkID string) func(*sql.Tx) error {
	return func(tx *sql.Tx) error {
		return store.initTable(tx, GetConfigTableName(networkID))
	}
}

func getWhereConditionFromFilterCriteria(criteria *FilterCriteria) sq.And {
	andClause := sq.And{}
	if !funk.IsEmpty(criteria.Type) {
		andClause = append(andClause, sq.Eq{typeCol: criteria.Type})
	}
	if !funk.IsEmpty(criteria.Key) {
		andClause = append(andClause, sq.Eq{keyCol: criteria.Key})
	}
	return andClause
}

func getExactWhereCondition(configType string, key string) sq.And {
	return sq.And{sq.Eq{typeCol: configType}, sq.Eq{keyCol: key}}
}

func validateFilterCriteria(criteria *FilterCriteria) error {
	if len(criteria.Key) == 0 && len(criteria.Type) == 0 {
		return errors.New("At least one field of filter criteria must be specified")
	}
	return nil
}
