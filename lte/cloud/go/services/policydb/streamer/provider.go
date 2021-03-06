/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package streamer

import (
	"os"
	"sort"

	"magma/lte/cloud/go/lte"
	protos2 "magma/lte/cloud/go/protos"
	"magma/lte/cloud/go/services/policydb"
	"magma/lte/cloud/go/services/policydb/obsidian/models"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/services/magmad"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

type PoliciesProvider struct{}

func (provider *PoliciesProvider) GetStreamName() string {
	return "policydb"
}

func (provider *PoliciesProvider) GetUpdates(gatewayId string, extraArgs *any.Any) ([]*protos.DataUpdate, error) {
	migrated := os.Getenv(orc8r.UseConfiguratorEnv)
	if migrated == "1" {
		gwEnt, err := configurator.LoadEntityForPhysicalID(gatewayId, configurator.EntityLoadCriteria{})
		if err != nil {
			return nil, err
		}

		ruleEnts, err := configurator.LoadAllEntitiesInNetwork(gwEnt.NetworkID, lte.PolicyRuleEntityType, configurator.EntityLoadCriteria{LoadConfig: true})
		if err != nil {
			return nil, err
		}

		ruleProtos := make([]*protos2.PolicyRule, 0, len(ruleEnts))
		for _, rule := range ruleEnts {
			ruleConfig := rule.Config.(*models.PolicyRule)
			ruleProto := &protos2.PolicyRule{}
			err = ruleConfig.ToProto(ruleProto)
			if err != nil {
				return nil, err
			}
			ruleProtos = append(ruleProtos, ruleProto)
		}
		return rulesToUpdates(ruleProtos)
	} else {
		networkId, err := magmad.FindGatewayNetworkId(gatewayId)
		if err != nil {
			return nil, err
		}
		policies, err := policydb.GetAllRules(networkId)
		if err != nil {
			return nil, err
		}
		return rulesToUpdates(policies)
	}
}

func rulesToUpdates(rules []*protos2.PolicyRule) ([]*protos.DataUpdate, error) {
	ret := make([]*protos.DataUpdate, 0, len(rules))
	for _, policy := range rules {
		marshaledPolicy, err := proto.Marshal(policy)
		if err != nil {
			return nil, err
		}
		ret = append(ret, &protos.DataUpdate{Key: policy.Id, Value: marshaledPolicy})
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].Key < ret[j].Key })
	return ret, nil
}

type BaseNamesProvider struct{}

func (provider *BaseNamesProvider) GetStreamName() string {
	return "base_names"
}

func (provider *BaseNamesProvider) GetUpdates(gatewayId string, extraArgs *any.Any) ([]*protos.DataUpdate, error) {
	migrated := os.Getenv(orc8r.UseConfiguratorEnv)
	if migrated == "1" {
		gwEnt, err := configurator.LoadEntityForPhysicalID(gatewayId, configurator.EntityLoadCriteria{})
		if err != nil {
			return nil, err
		}

		bnEnts, err := configurator.LoadAllEntitiesInNetwork(
			gwEnt.NetworkID,
			lte.BaseNameEntityType,
			configurator.EntityLoadCriteria{LoadConfig: true, LoadAssocsFromThis: true},
		)
		if err != nil {
			return nil, err
		}

		bnProtos := make([]*protos2.ChargingRuleBaseNameRecord, 0, len(bnEnts))
		for _, bn := range bnEnts {
			bnConfig := bn.Config.(*models.BaseNameRecord)
			ruleNames := make([]string, 0, len(bn.Associations))
			for _, assoc := range bn.Associations {
				ruleNames = append(ruleNames, assoc.Key)
			}
			bnProto := &protos2.ChargingRuleBaseNameRecord{
				Name:         string(bnConfig.Name),
				RuleNamesSet: &protos2.ChargingRuleNameSet{RuleNames: ruleNames},
			}
			bnProtos = append(bnProtos, bnProto)
		}
		return bnsToUpdates(bnProtos)
	} else {
		networkId, err := magmad.FindGatewayNetworkId(gatewayId)
		if err != nil {
			return nil, err
		}
		baseNameRecords, err := policydb.GetAllBaseNames(networkId)
		if err != nil {
			return nil, err
		}
		return bnsToUpdates(baseNameRecords)
	}
}

func bnsToUpdates(bns []*protos2.ChargingRuleBaseNameRecord) ([]*protos.DataUpdate, error) {
	ret := make([]*protos.DataUpdate, 0, len(bns))
	for _, bn := range bns {
		marshaledBN, err := proto.Marshal(bn)
		if err != nil {
			return nil, err
		}
		ret = append(ret, &protos.DataUpdate{Key: bn.Name, Value: marshaledBN})
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].Key < ret[j].Key })
	return ret, nil
}
