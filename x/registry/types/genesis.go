package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		WhoisList:   []Whois{},
		WhatisList:  []Whatis{},
		ThereisList: []Thereis{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in whois
	whoisIndexMap := make(map[string]struct{})

	for _, elem := range gs.WhoisList {
		index := string(WhoisKey(elem.Index))
		if _, ok := whoisIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for whois")
		}
		whoisIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in whatis
	whatisIndexMap := make(map[string]struct{})

	for _, elem := range gs.WhatisList {
		index := string(WhatisKey(elem.Index))
		if _, ok := whatisIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for whatis")
		}
		whatisIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in thereis
	thereisIndexMap := make(map[string]struct{})

	for _, elem := range gs.ThereisList {
		index := string(ThereisKey(elem.Index))
		if _, ok := thereisIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for thereis")
		}
		thereisIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
