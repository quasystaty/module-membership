package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/noria-net/module-membership/x/membership/types"
)

func ParseDirectDemocracyUpdateProposal(cdc codec.JSONCodec, proposalPath string) (proposal types.DirectDemocracyUpdateProposal, err error) {
	proposal = types.DirectDemocracyUpdateProposal{}
	contents, err := os.ReadFile(proposalPath)
	if err != nil {
		return proposal, err
	}

	err = cdc.UnmarshalJSON(contents, &proposal)
	return proposal, err
}
