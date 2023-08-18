package types

import (
	"fmt"

	"cosmossdk.io/math"
)

// DirectDemocracyUpdateProposal creates an empty proposal instance
func NewEmptyDirectDemocracyUpdateProposal() DirectDemocracyUpdateProposal {
	return DirectDemocracyUpdateProposal{
		GuardiansToAdd:    []string{},
		GuardiansToRemove: []string{},
	}
}

// Validate performs basic validation on the proposal
func (p *DirectDemocracyUpdateProposal) Validate() error {

	// Cannot add and remove the same guardian
	for _, addGuardian := range p.GuardiansToAdd {
		for _, removeGuardian := range p.GuardiansToRemove {
			if addGuardian == removeGuardian {
				return fmt.Errorf("cannot add and remove the same guardian: %s", addGuardian)
			}
		}
	}

	// Cannot have empty guardian lists AND an empty total voting weight
	if len(p.GuardiansToAdd) == 0 &&
		len(p.GuardiansToRemove) == 0 &&
		p.TotalVotingWeight == nil {
		return fmt.Errorf("nothing to do")
	}

	// Total voting weight must be between 0 and 1, inclusive
	if p.TotalVotingWeight != nil {
		if p.TotalVotingWeight.LT(math.LegacyMustNewDecFromStr(MINIMUM_TOTAL_VOTING_WEIGHT)) ||
			p.TotalVotingWeight.GT(math.LegacyMustNewDecFromStr(MAXIMUM_TOTAL_VOTING_WEIGHT)) {
			return fmt.Errorf("total voting weight must be between %s and %s, inclusive", MINIMUM_TOTAL_VOTING_WEIGHT, MAXIMUM_TOTAL_VOTING_WEIGHT)
		}
	}

	return nil
}
