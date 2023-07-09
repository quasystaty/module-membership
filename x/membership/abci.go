package membership

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/noria-net/module-membership/x/membership/keeper"
	"github.com/noria-net/module-membership/x/membership/types"
)

func EndBlocker(ctx sdk.Context, keeper *keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// fetch active proposals whose voting periods have ended (are passed the block time)
	keeper.IterateActiveProposalsQueue(ctx, ctx.BlockHeader().Time, func(proposal v1.Proposal) (stop bool) {
		return processActiveProposal(ctx, keeper, proposal)
	})
}

func processActiveProposal(ctx sdk.Context, keeper *keeper.Keeper, proposal v1.Proposal) (stop bool) {
	logger := keeper.Logger(ctx)
	var tagValue, logMsg string

	passes, burnDeposits, tallyResults := keeper.Tally(ctx, proposal)

	if burnDeposits {
		keeper.DeleteAndBurnDeposits(ctx, proposal.Id)
	} else {
		keeper.RefundAndDeleteDeposits(ctx, proposal.Id)
	}

	if passes {
		var (
			idx    int
			events sdk.Events
			msg    sdk.Msg
		)

		// attempt to execute all messages within the passed proposal
		// Messages may mutate state thus we use a cached context. If one of
		// the handlers fails, no state mutation is written and the error
		// message is logged.
		cacheCtx, writeCache := ctx.CacheContext()
		messages, err := proposal.GetMsgs()
		if err == nil {
			for idx, msg = range messages {
				handler := keeper.GovRouter().Handler(msg)

				var res *sdk.Result
				res, err = handler(cacheCtx, msg)
				if err != nil {
					break
				}

				events = append(events, res.GetEvents()...)
			}
		}

		// `err == nil` when all handlers passed.
		// Or else, `idx` and `err` are populated with the msg index and error.
		if err == nil {
			proposal.Status = v1.StatusPassed
			tagValue = govtypes.AttributeValueProposalPassed
			logMsg = "passed"

			// write state to the underlying multi-store
			writeCache()

			// propagate the msg events to the current context
			ctx.EventManager().EmitEvents(events)
		} else {
			proposal.Status = v1.StatusFailed
			tagValue = govtypes.AttributeValueProposalFailed
			logMsg = fmt.Sprintf("passed, but msg %d (%s) failed on execution: %s", idx, sdk.MsgTypeURL(msg), err)
		}
	} else {
		proposal.Status = v1.StatusRejected
		tagValue = govtypes.AttributeValueProposalRejected
		logMsg = "rejected"
	}

	proposal.FinalTallyResult = &tallyResults

	keeper.SetProposal(ctx, proposal)
	keeper.RemoveFromActiveProposalQueue(ctx, proposal.Id, *proposal.VotingEndTime)

	// when proposal become active
	keeper.GovHooks().AfterProposalVotingPeriodEnded(ctx, proposal.Id)

	logger.Info(
		"proposal tallied",
		"proposal", proposal.Id,
		"results", logMsg,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			govtypes.EventTypeActiveProposal,
			sdk.NewAttribute(govtypes.AttributeKeyProposalID, fmt.Sprintf("%d", proposal.Id)),
			sdk.NewAttribute(govtypes.AttributeKeyProposalResult, tagValue),
		),
	)
	return false

}
