/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package servicers

import (
	"testing"

	fegprotos "magma/feg/cloud/go/protos"
	"magma/lte/cloud/go/lte"
	cellular "magma/lte/cloud/go/services/cellular/config"
	cellular_protos "magma/lte/cloud/go/services/cellular/protos"
	utils "magma/lte/cloud/go/services/eps_authentication/servicers/test_utils"
	"magma/lte/cloud/go/services/subscriberdb/storage"
	orc8rprotos "magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/serde"
	"magma/orc8r/cloud/go/services/config"
	"magma/orc8r/cloud/go/services/config/test_init"
	"magma/orc8r/cloud/go/test_utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
)

// EpsAuthTestSuite is a test suite that will setup a test EPS auth server
// and config service pre-populated with a cellular config.
type EpsAuthTestSuite struct {
	suite.Suite
	Server *EPSAuthServer
}

func (suite *EpsAuthTestSuite) AuthenticationInformation(air *fegprotos.AuthenticationInformationRequest) (*fegprotos.AuthenticationInformationAnswer, error) {
	return suite.Server.AuthenticationInformation(getTestContext(), air)
}

func (suite *EpsAuthTestSuite) UpdateLocation(ulr *fegprotos.UpdateLocationRequest) (*fegprotos.UpdateLocationAnswer, error) {
	return suite.Server.UpdateLocation(getTestContext(), ulr)
}

func (suite *EpsAuthTestSuite) PurgeUE(purge *fegprotos.PurgeUERequest) (*fegprotos.PurgeUEAnswer, error) {
	return suite.Server.PurgeUE(getTestContext(), purge)
}

func (suite *EpsAuthTestSuite) SetupTest() {
	store, err := storage.NewSubscriberDBStorage(test_utils.NewMockDatastore())
	suite.NoError(err)

	for _, subscriber := range utils.GetTestSubscribers() {
		_, err := store.AddSubscriber(subscriber)
		suite.NoError(err)
	}

	server, err := NewEPSAuthServer(store)
	suite.NoError(err)
	suite.Server = server
}

func TestEpsAuthSuite(t *testing.T) {
	test_init.StartTestService(t)
	err := serde.RegisterSerdes(&cellular.CellularNetworkConfigManager{})
	assert.NoError(t, err)

	configProto := &cellular_protos.CellularNetworkConfig{
		Ran: &cellular_protos.NetworkRANConfig{},
		Epc: &cellular_protos.NetworkEPCConfig{
			Mcc:        "123",
			Mnc:        "123",
			Tac:        1,
			LteAuthOp:  []byte("\xcd\xc2\x02\xd5\x12> \xf6+mgj\xc7,\xb3\x18"),
			LteAuthAmf: []byte("\x80\x00"),
			SubProfiles: map[string]*cellular_protos.NetworkEPCConfig_SubscriptionProfile{
				"default": &cellular_protos.NetworkEPCConfig_SubscriptionProfile{
					MaxUlBitRate: 1000,
					MaxDlBitRate: 2000,
				},
				"test_profile": &cellular_protos.NetworkEPCConfig_SubscriptionProfile{
					MaxUlBitRate: 7000,
					MaxDlBitRate: 5000,
				},
			},
		},
	}
	err = config.CreateConfig("test", lte.CellularNetworkType, "test", configProto)
	assert.NoError(t, err)

	testSuite := &EpsAuthTestSuite{}
	suite.Run(t, testSuite)
}

func getTestContext() context.Context {
	return orc8rprotos.NewGatewayIdentity("test", "test", "test").NewContextWithIdentity(context.Background())
}
