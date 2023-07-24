package keeper

import (
	"errors"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/noria-net/module-membership/x/membership/types"
)

// VoteOptions is a map of vote options to the number of votes for that option
type voteOptions map[govtypes_v1.VoteOption]math.Int

// combinedTallyResults is a struct containing the results of a tally of both
// guardian and regular member votes
type combinedTallyResults struct {
	results          map[govtypes_v1.VoteOption]math.LegacyDec
	votingPower      math.LegacyDec
	numGuardianVotes math.Int
	numMemberVotes   math.Int
	totalVotes       math.Int
}

// Tally iterates over the votes and updates the tally of a proposal based on the voting power of the
// voters
func (k Keeper) Tally(ctx sdk.Context, proposal govtypes_v1.Proposal) (passes bool, burnDeposits bool, tallyResults govtypes_v1.TallyResult) {

	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	guardians := k.GetGuardians(ctx)

	memberPower, guardianPower := calculateVotePower(
		int64(k.GetMemberStatusCount(ctx, types.MembershipStatus_MemberElectorate)),
		int64(len(guardians)),
		k.GetTotalVotingWeight(ctx),
	)

	k.IterateVotes(ctx, proposal.Id, func(vote govtypes_v1.Vote) (stop bool) {
		// Create a custom logger for this voter
		voterLogger := ctx.WithLogger(ctx.Logger().With("voter", vote.Voter))

		voterAddress := sdk.MustAccAddressFromBech32(vote.Voter)
		member, found := k.GetMemberAccount(ctx, voterAddress)

		err := processSingleVote(vote,
			&member,
			found,
			memberResults,
			guardianResults,
		)

		// Getting an error doesn't stop us iterating through the votes
		if err != nil {
			voterLogger.Logger().Error(fmt.Sprintf("Error processing vote: %s", err.Error()))
		}

		// Delete this vote, now that its been processed
		k.deleteVote(ctx, vote.ProposalId, voterAddress)

		return false
	})

	govParams := k.GetGovParams(ctx)
	passes, burnDeposits, tallyResults = calculateVoteResults(proposal,
		govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	return passes, burnDeposits, tallyResults
}

// processSingleVote processes a single vote, updating the tally results
func processSingleVote(vote govtypes_v1.Vote,
	member *types.Member,
	found bool,
	memberResults voteOptions,
	guardianResults voteOptions) error {

	// voter must be a denom member
	if !found {
		return errors.New("voter is not a member of the electorate")
	}

	// member must be eligible to vote
	if member.Status != types.MembershipStatus_MemberElectorate {
		return errors.New("member is not eligible to vote")
	}

	// member's vote weight must be valid
	ok, weightingError := isValidVoteWeighting(vote.Options)
	if !ok {
		return fmt.Errorf("invalid voting weight: %s", weightingError.Error())
	}

	choice := getVoterChoice(vote.Options)
	if member.IsGuardian {
		guardianResults[choice] = guardianResults[choice].Add(math.NewInt(1))
	} else {
		memberResults[choice] = memberResults[choice].Add(math.NewInt(1))
	}

	return nil
}

// isValidVoteWeighting checks if the vote has been made on a single option, and not spread across more than one
func isValidVoteWeighting(options []*govtypes_v1.WeightedVoteOption) (bool, error) {

	totalWeight := sdk.NewDec(0)
	for _, option := range options {
		weight := sdk.MustNewDecFromStr(option.Weight)
		// Cannot have any other weighting besides 0 or 1
		if !weight.IsZero() && !weight.Equal(sdk.NewDec(1)) {
			return false, fmt.Errorf("option %s's weight is invalid: %s", option.Option, option.Weight)
		}
		totalWeight = totalWeight.Add(weight)

		// Cannot have a total weight of more than 1
		if !totalWeight.Equal(sdk.NewDec(1)) {
			return false, fmt.Errorf("vote is spoilt, total weighting of %s exceeds 1", option.Weight)
		}
	}

	return true, nil
}

// getVoterChoice returns the vote option that the voter has chosen
func getVoterChoice(options []*govtypes_v1.WeightedVoteOption) govtypes_v1.VoteOption {
	for _, option := range options {
		optionDec, err := sdk.NewDecFromStr(option.Weight)
		if err != nil {
			panic(err)
		}

		if optionDec.Equal(sdk.NewDec(1)) {
			return option.Option
		}
	}
	return govtypes_v1.OptionEmpty
}

func calculateVoteResults(proposal govtypes_v1.Proposal,
	govParams govtypes_v1.Params,
	memberResults voteOptions,
	guardianResults voteOptions,
	memberPower sdk.Dec,
	guardianPower sdk.Dec) (passes bool, burnDeposits bool, tallyResults govtypes_v1.TallyResult) {

	// Get relevant gov params
	quorum := sdk.MustNewDecFromStr(govParams.Quorum).MulInt64(100)
	vetoThreshold := sdk.MustNewDecFromStr(govParams.VetoThreshold).MulInt64(100)

	// Calculate total votes counted
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)
	tallyResults = govtypes_v1.NewTallyResultFromMap(combined.results)

	// If there is not enough voting power to reach quorum, proposal fails
	// combinedVotingPower = (Guardian_Votes * Guardian_Power) + (NormalMember_Votes * NormalMember_Power)
	if combined.votingPower.LT(quorum) {
		return false, govParams.BurnVoteQuorum, tallyResults
	}

	// If no one votes (everyone abstains), proposal fails
	if combined.votingPower.Sub(combined.results[govtypes_v1.OptionAbstain]).Equal(math.LegacyZeroDec()) {
		return false, false, tallyResults
	}

	// If more than 1/3 of voters veto, proposal fails
	// VetoPortion = (Guardian_NoWithVetoVotes / (Guardian_Votes - Guardian_AbstainVotes)) * Guardian_Power
	// + (NormalMember_NoWithVetoVotes / (NormalMember_Votes - NormalMember_AbstainVotes)) * NormalMember_Power
	guardianVeto := calculateVeto(guardianResults, combined.numGuardianVotes, guardianPower)
	memberVeto := calculateVeto(memberResults, combined.numMemberVotes, memberPower)
	combinedVeto := guardianVeto.Add(memberVeto)
	if combinedVeto.GT(vetoThreshold) {
		return false, govParams.BurnVoteVeto, tallyResults
	}

	// If combined Yes votes exceeds threshold, proposal passes
	// YesPortion = (Guardian_YesVotes / Guardian_Votes) * Guardian_Power + (NormalMember_YesVotes / NormalMember_Votes) * NormalMember_Power
	guardianYes := calculateWeightedOptionVote(guardianResults[govtypes_v1.OptionYes], combined.numGuardianVotes, guardianPower)
	memberYes := calculateWeightedOptionVote(memberResults[govtypes_v1.OptionYes], combined.numMemberVotes, memberPower)
	combinedYes := guardianYes.Add(memberYes)
	yesThreshold := sdk.MustNewDecFromStr(govParams.Threshold).MulInt64(100)
	if combinedYes.GT(yesThreshold) {
		return true, false, tallyResults
	}

	// If more than {threshold} of non-abstaining voters vote No, proposal fails
	return false, false, tallyResults
}

func (k Keeper) deleteVote(ctx sdk.Context, proposalID uint64, voter sdk.AccAddress) {
	// TODO: ignore this call - we need Gov to open up deleteVote in their keeper first
	// https://github.com/cosmos/cosmos-sdk/blob/v0.45.2/x/gov/keeper/vote.go#L129

	// One option to handle this is to delete votes on an upgrade using a migration step
}

// calculateVotePower calculates the voting power of members and guardians
func calculateVotePower(numElectorateMembers int64, numGuardians int64, totalVotingWeight math.LegacyDec) (memberPower math.LegacyDec, guardianPower math.LegacyDec) {

	// Member count excludes guardians
	numMembers := numElectorateMembers - numGuardians
	memberPower = sdk.NewDec(100).Sub(totalVotingWeight).QuoInt64(numMembers)
	guardianPower = totalVotingWeight.QuoInt64(numGuardians)

	return memberPower, guardianPower
}

// makeResultMap returns a map with all the vote options set to 0
func NewEmptyVoteOptions() voteOptions {
	results := make(map[govtypes_v1.VoteOption]math.Int)
	results[govtypes_v1.OptionYes] = math.ZeroInt()
	results[govtypes_v1.OptionAbstain] = math.ZeroInt()
	results[govtypes_v1.OptionNo] = math.ZeroInt()
	results[govtypes_v1.OptionNoWithVeto] = math.ZeroInt()
	return results
}

// calculateCombinedTallyResults combines the results of the member and guardian votes,
// and uses the voting power of each group to calculate the total voting power of each option
func calculateCombinedTallyResults(memberResults,
	guardianResults map[govtypes_v1.VoteOption]math.Int,
	memberPower sdk.Dec,
	guardianPower sdk.Dec,
) combinedTallyResults {
	combined := combinedTallyResults{
		results:          make(map[govtypes_v1.VoteOption]math.LegacyDec),
		votingPower:      sdk.ZeroDec(),
		numGuardianVotes: math.ZeroInt(),
		numMemberVotes:   math.ZeroInt(),
		totalVotes:       math.ZeroInt(),
	}

	for option, guardianVoteCount := range guardianResults {
		combined.results[option] = guardianPower.Mul(math.LegacyNewDecFromInt(guardianVoteCount)).Add(
			memberPower.Mul(math.LegacyNewDecFromInt(memberResults[option])))
		combined.votingPower = combined.votingPower.Add(combined.results[option])
		combined.numGuardianVotes = combined.numGuardianVotes.Add(guardianVoteCount)
		combined.numMemberVotes = combined.numMemberVotes.Add(memberResults[option])
		combined.totalVotes = combined.totalVotes.Add(guardianVoteCount).Add(memberResults[option])
	}

	return combined
}

// calculateVeto calculates the weighted veto of a group of voters
func calculateVeto(results voteOptions, numVotes math.Int, power math.LegacyDec) math.LegacyDec {
	// Cannot calculate weighted veto if there are no votes
	if numVotes.IsZero() {
		return math.LegacyNewDec(0)
	}
	// (NoWithVetoVotes / (NumVotes - AbstainVotes)) * Power
	veto := math.LegacyNewDecFromInt(results[govtypes_v1.OptionNoWithVeto]).Quo(
		math.LegacyNewDecFromInt(numVotes).Sub(math.LegacyNewDecFromInt(results[govtypes_v1.OptionAbstain]))).Mul(power)
	return veto
}

// calculateWeightedOptionVote calculates the weighted vote of a group of voters
func calculateWeightedOptionVote(numOptionVotes math.Int, numVotes math.Int, power math.LegacyDec) math.LegacyDec {
	// Cannot calculate the vote's weighted option if there are no votes
	if numVotes.IsZero() {
		return math.LegacyNewDec(0)
	}
	return math.LegacyNewDecFromInt(numOptionVotes).Quo(math.LegacyNewDecFromInt(numVotes)).Mul(power)
}
