package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here

	// Check if an account exists in the store.
	HasAccount(sdk.Context, sdk.AccAddress) bool
	// Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(sdk.Context, sdk.AccAddress) types.AccountI
	// Set an account in the store.
	SetAccount(sdk.Context, types.AccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// internalGovKeeper implements everything except Hooks(), which expects a pointer receiver
type internalGovKeeper interface {
	// IterateInactiveProposalsQueue iterates over the proposals in the inactive proposal queue
	// and performs a callback function
	IterateInactiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal govtypes_v1.Proposal) (stop bool))
	// IterateActiveProposalsQueue iterates over the proposals in the active proposal queue
	// and performs a callback function
	IterateActiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal govtypes_v1.Proposal) (stop bool))
	// IterateVotes iterates over the all the proposals votes and performs a callback function
	IterateVotes(ctx sdk.Context, proposalID uint64, cb func(vote govtypes_v1.Vote) (stop bool))
	// RefundAndDeleteDeposits refunds and deletes all the deposits on a specific proposal.
	RefundAndDeleteDeposits(ctx sdk.Context, proposalID uint64)
	// Router returns the gov Keeper's Router
	Router() *baseapp.MsgServiceRouter
	// GetProposal gets a proposal from store by ProposalID.
	GetProposal(ctx sdk.Context, proposalID uint64) (govtypes_v1.Proposal, bool)
	// SetProposal set a proposal to store
	SetProposal(ctx sdk.Context, proposal govtypes_v1.Proposal)
	// RemoveFromActiveProposalQueue removes a proposalID from the Active Proposal Queue
	RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time)
	// DeleteProposal deletes a proposal from store.
	// Panics if the proposal doesn't exist.
	DeleteProposal(ctx sdk.Context, proposalID uint64)
	// GetParams gets the gov module's parameters.
	GetParams(clientCtx sdk.Context) (params govtypes_v1.Params)
	// DeleteAndBurnDeposits deletes and burns all the deposits on a specific proposal.
	DeleteAndBurnDeposits(ctx sdk.Context, proposalID uint64)
}

// GovKeeper implements our expected contract as well as exposes Gov's hooks
// Note: these hooks are captured at initialisation and stored in GovKeeper.Hooks
// Note: Please define all expected functions in the interface above, and not in this struct
type GovKeeper struct {
	internalGovKeeper
	Hooks govtypes.GovHooks
}

// NewExtendedGovKeeper creates a new instance of GovKeeper with hooks
func NewExtendedGovKeeper(gk govkeeper.Keeper) GovKeeper {
	return GovKeeper{
		internalGovKeeper: gk,
		Hooks:             gk.Hooks(),
	}
}
