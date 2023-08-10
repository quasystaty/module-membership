package types_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/noria-net/module-membership/testutil/sample"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	knownGuardianAddress := sample.AccAddress()

	tests := []struct {
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
			desc: "valid genesis state: one guardian",
			genState: &types.GenesisState{
				DirectDemocracy: types.DirectDemocracy{
					TotalVotingWeight: math.LegacyZeroDec(),
					Guardians: []string{
						sample.AccAddress(),
					},
				},

				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "invalid genesis state: bad guardian address",
			genState: &types.GenesisState{
				DirectDemocracy: types.DirectDemocracy{
					TotalVotingWeight: math.LegacyZeroDec(),
					Guardians: []string{
						"bad address",
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid genesis state: duplicate guardian address",
			genState: &types.GenesisState{
				DirectDemocracy: types.DirectDemocracy{
					TotalVotingWeight: math.LegacyZeroDec(),
					Guardians: []string{
						knownGuardianAddress,
						knownGuardianAddress,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid genesis state: total voting weight > 100",
			genState: &types.GenesisState{
				DirectDemocracy: types.DirectDemocracy{
					TotalVotingWeight: math.LegacyMustNewDecFromStr("101"),
					Guardians: []string{
						sample.AccAddress(),
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid genesis state: total voting weight < 0",
			genState: &types.GenesisState{
				DirectDemocracy: types.DirectDemocracy{
					TotalVotingWeight: math.LegacyMustNewDecFromStr("-1"),
					Guardians: []string{
						sample.AccAddress(),
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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
