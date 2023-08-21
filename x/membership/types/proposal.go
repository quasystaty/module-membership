package types

import (
	"fmt"
	"strings"

	"cosmossdk.io/math"
	gov_v1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// ProposalTypeDirectDemocracyUpdate defines the type for a DirectDemocracyUpdateProposal
	ProposalTypeDirectDemocracyUpdate = "DirectDemocracyUpdate"
)

// Ensure all proposals implement govtypes.Content at compile time
var (
	_ gov_v1beta1.Content = &DirectDemocracyUpdateProposal{}
)

////////
// Direct Democracy Proposal
////////

// DirectDemocracyUpdateProposal creates an empty proposal instance
func NewEmptyDirectDemocracyUpdateProposal() DirectDemocracyUpdateProposal {
	return DirectDemocracyUpdateProposal{
		GuardiansToAdd:    []string{},
		GuardiansToRemove: []string{},
	}
}

// GetTitle returns the title of a direct democracy update proposal.
func (p *DirectDemocracyUpdateProposal) GetTitle() string { return p.Title }

// GetDescription returns the description of a direct democracy update proposal.
func (p *DirectDemocracyUpdateProposal) GetDescription() string { return p.Description }

// ProposalRoute ensures this proposal will be handled by the Membership Module
func (p *DirectDemocracyUpdateProposal) ProposalRoute() string { return ModuleName }

func (p *DirectDemocracyUpdateProposal) ProposalType() string {
	return ProposalTypeDirectDemocracyUpdate
}

// Validate performs basic validation on the proposal
func (p *DirectDemocracyUpdateProposal) ValidateBasic() error {

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

// String implements fmt.Stringer
func (p *DirectDemocracyUpdateProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Direct Democracy Update Proposal:
  Title:              %s
  Description:        %s
  Guardians to Add:   %s
  Guardians to Remove:%s
  Total Voting Weight:%s
`, p.Title, p.Description, p.GuardiansToAdd, p.GuardiansToRemove, p.TotalVotingWeight))
	return b.String()
}
