/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package view_factory_test

import (
	"encoding/json"
	"testing"

	"magma/orc8r/cloud/go/services/magmad/obsidian/handlers/view_factory"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/stretchr/testify/assert"
)

func TestJSONMapToProtobufStruct(t *testing.T) {
	jsonMap := map[string]interface{}{
		"nil":    nil,
		"number": 1.0,
		"string": "str",
		"struct": map[string]interface{}{
			"a": 2.0,
		},
		"list": []interface{}{1.0, "foo"},
	}
	marshaled, err := json.Marshal(jsonMap)
	assert.NoError(t, err)
	expectedProtobufStruct := &structpb.Struct{}
	err = jsonpb.UnmarshalString(string(marshaled), expectedProtobufStruct)
	assert.NoError(t, err)

	actualProtobufStruct, err := view_factory.JSONMapToProtobufStruct(jsonMap)

	assert.NoError(t, err)
	assert.Equal(t, expectedProtobufStruct, actualProtobufStruct)
}

func TestProtobufStructToJSONMap(t *testing.T) {
	expectedJsonMap := map[string]interface{}{
		"nil":    nil,
		"number": 1.0,
		"string": "str",
		"struct": map[string]interface{}{
			"a": 2.0,
		},
		"list": []interface{}{1.0, "foo"},
	}
	marshaled, err := json.Marshal(expectedJsonMap)
	assert.NoError(t, err)
	protobufStruct := &structpb.Struct{}
	err = jsonpb.UnmarshalString(string(marshaled), protobufStruct)
	assert.NoError(t, err)

	actualJsonMap, err := view_factory.ProtobufStructToJSONMap(protobufStruct)

	assert.NoError(t, err)
	assert.Equal(t, expectedJsonMap, actualJsonMap)
}
