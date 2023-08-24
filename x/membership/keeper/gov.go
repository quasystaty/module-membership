package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// IterateInactiveProposalsQueue iterates over the proposals in the inactive proposal queue
// and performs a callback function
func (k Keeper) IterateInactiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal govtypes_v1.Proposal) (stop bool)) {
	k.govKeeper.IterateInactiveProposalsQueue(ctx, endTime, cb)
}

// IterateActiveProposalsQueue cycle through proposals that have ended their voting period
func (k Keeper) IterateActiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal govtypes_v1.Proposal) (stop bool)) {
	// Pass-through to Gov keeper
	k.govKeeper.IterateActiveProposalsQueue(ctx, endTime, cb)
}

// IterateVotes iterates over the all the proposals votes and performs a callback function
func (k Keeper) IterateVotes(ctx sdk.Context, proposalID uint64, cb func(vote govtypes_v1.Vote) (stop bool)) {
	k.govKeeper.IterateVotes(ctx, proposalID, cb)
}

// RefundAndDeleteDeposits refunds and deletes all the deposits on a specific proposal.
func (k Keeper) RefundAndDeleteDeposits(ctx sdk.Context, proposalID uint64) {
	k.govKeeper.RefundAndDeleteDeposits(ctx, proposalID)
}

// IsLegitimateProposal returns true if this proposal exists and was created by an electorate member
func (k Keeper) IsLegitimateProposal(ctx sdk.Context, proposal govtypes_v1.Proposal) bool {
	p, proposalExists := k.govKeeper.GetProposal(ctx, proposal.Id)
	if !proposalExists {
		return false
	}

	return k.IsMember(ctx, sdk.MustAccAddressFromBech32(p.Proposer))
}

// SetProposal writes the updated proposal to the store
func (k Keeper) SetProposal(ctx sdk.Context, proposal govtypes_v1.Proposal) {
	k.govKeeper.SetProposal(ctx, proposal)
}

// RemoveFromActiveProposalQueue removes a proposalID from the Active Proposal Queue
func (k Keeper) RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	k.govKeeper.RemoveFromActiveProposalQueue(ctx, proposalID, endTime)
}

// DeleteProposal deletes a proposal from store.
// Panics if the proposal doesn't exist.
func (k Keeper) DeleteProposal(ctx sdk.Context, proposalID uint64) {
	k.govKeeper.DeleteProposal(ctx, proposalID)
}

// DeleteAndBurnDeposits deletes and burns all the deposits on a specific proposal.
func (k Keeper) DeleteAndBurnDeposits(ctx sdk.Context, proposalID uint64) {
	k.govKeeper.DeleteAndBurnDeposits(ctx, proposalID)
}

// GetGovParams gets the governance parameters from the global param store
func (k Keeper) GetGovParams(ctx sdk.Context) (params govtypes_v1.Params) {
	return k.govKeeper.GetParams(ctx)
}

// Hooks gets the hooks for governance Keeper
func (k *Keeper) GovHooks() govtypes.GovHooks {
	return k.govKeeper.Hooks
}

// Router returns the gov keeper's router
func (k *Keeper) GovRouter() *baseapp.MsgServiceRouter {
	return k.govKeeper.Router()
}
