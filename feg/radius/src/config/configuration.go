/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package config

import (
	"encoding/json"
	"fbc/cwf/radius/modules"
	"fbc/lib/go/oc"
	"io/ioutil"
)

// LiveTier name
const LiveTier = "live"

type (
	// ModuleDescriptor a descriptor for loading a single module
	ModuleDescriptor struct {
		Name   string               `json:"name"`
		Config modules.ModuleConfig `json:"config"`
	}

	// ListenerConfig for a single listener (server has a listerner per each port)
	ListenerConfig struct {
		Name    string             `json:"name"`
		Port    int                `json:"port"`
		Type    string             `json:"type"`
		Modules []ModuleDescriptor `json:"modules"`
	}

	// ServiceTier represents a uniquely identifiable named set of upstream hosts
	ServiceTier struct {
		Name          string   `json:"name"`
		UpstreamHosts []string `json:"upstreamHosts"`
	}

	// ListenerRoute maps a listener to a ServiceTier name
	ListenerRoute struct {
		Listener    string `json:"listener"`
		ServiceTier string `json:"serviceTier"`
	}

	// TierRouting a set of ListenerRoute representing a map from Listener names to their respective ServiceTier
	TierRouting struct {
		Routes []ListenerRoute `json:"tierRoutes"`
	}

	// Canary represents a definition of a canary
	Canary struct {
		Name                string      `json:"name"`
		TrafficSlicePercent int         `json:"trafficSlicePercent"`
		Routing             TierRouting `json:"routing"`
	}

	// LoadBalanceConfig holds the complete configuration for a server load balancer
	LoadBalanceConfig struct {
		ServiceTiers []ServiceTier `json:"serviceTiers"`
		LiveTier     TierRouting   `json:"liveTier"`
		Canaries     []Canary      `json:"canaries"`
		DefaultTier  string        `json:"defaultTier"`
	}

	// ServerConfig Encapsulates the configuration of a radius server
	ServerConfig struct {
		Secret      string            `json:"secret"`
		DedupWindow Duration          `json:"dedupWindow"`
		LoadBalance LoadBalanceConfig `json:"loadBalance"`
		Listeners   []ListenerConfig  `json:"listeners"`
		Filters     []string          `json:"filters"`
	}

	// RadiusConfig the configuration file format
	RadiusConfig struct {
		Counters oc.Config    `json:"counters"`
		Server   ServerConfig `json:"server"`
	}
)

// Read reads and parses a configuration file into a RadiusConfig
func Read(filename string) (*RadiusConfig, error) {
	configBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config RadiusConfig
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
