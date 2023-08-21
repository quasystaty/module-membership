package types_test

import (
	"errors"
	"testing"

	"cosmossdk.io/math"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/stretchr/testify/require"
)

const (
	valid_1 = "cosmos1l0znsvddllw9knha3yx2svnlxny676d8ns7uys"
	valid_2 = "cosmos1j8pp7zvcu9z8vd882m284j29fn2dszh05cqvf9"
	invalid = "invalid_address"
)

func TestProposalValidation(t *testing.T) {
	validVotingWeight := math.LegacyMustNewDecFromStr("0.5")
	// invalidVotingWeight := math.LegacyMustNewDecFromStr("1.5")

	for _, tc := range []struct {
		desc     string
		proposal types.DirectDemocracyUpdateProposal
		err      error
	}{
		{
			desc: "add a guardian",
			proposal: types.DirectDemocracyUpdateProposal{
				GuardiansToAdd: []string{
					valid_1,
				},
			},
		}, {
			desc: "remove a guardian",
			proposal: types.DirectDemocracyUpdateProposal{
				GuardiansToRemove: []string{
					valid_1,
				},
			},
		}, {
			desc: "set the total voting weight",
			proposal: types.DirectDemocracyUpdateProposal{
				TotalVotingWeight: &validVotingWeight,
			},
		}, {
			desc:     "empty proposal is invalid",
			proposal: types.NewEmptyDirectDemocracyUpdateProposal(),
			err:      errors.New("nothing to do"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.proposal.ValidateBasic()
			if tc.err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
