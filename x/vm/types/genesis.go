package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ContractsList: []Contracts{},
		ProgramList:   []Program{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in contracts
	contractsIdMap := make(map[uint64]bool)
	contractsCount := gs.GetContractsCount()
	for _, elem := range gs.ContractsList {
		if _, ok := contractsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for contracts")
		}
		if elem.Id >= contractsCount {
			return fmt.Errorf("contracts id should be lower or equal than the last id")
		}
		contractsIdMap[elem.Id] = true
	}
	// Check for duplicated index in program
	programIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProgramList {
		index := string(ProgramKey(elem.Name))
		if _, ok := programIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for program")
		}
		programIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
