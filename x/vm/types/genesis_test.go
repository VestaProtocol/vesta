package types_test

import (
	"testing"

	"github.com/TheMarstonConnell/vesta/x/vm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				ContractsList: []types.Contracts{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ContractsCount: 2,
				ProgramList: []types.Program{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				RomdataList: []types.Romdata{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated contracts",
			genState: &types.GenesisState{
				ContractsList: []types.Contracts{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid contracts count",
			genState: &types.GenesisState{
				ContractsList: []types.Contracts{
					{
						Id: 1,
					},
				},
				ContractsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated program",
			genState: &types.GenesisState{
				ProgramList: []types.Program{
					{
						Name: "0",
					},
					{
						Name: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated romdata",
			genState: &types.GenesisState{
				RomdataList: []types.Romdata{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
