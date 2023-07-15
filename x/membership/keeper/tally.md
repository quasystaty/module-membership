# Notes about Tallying votes with guardians

## Discussion with DrCisco

* Guardian voting power as % of total voting power -> GuardianPower = GuardianWeight / NumberOfGuardians
* Member voting power as % of total voting power -> MemberPower = (1-GuardianWeight) / (NumberOfMembers - NumberOfGuardians) 
* NumberOfMembers is all electorate members, including guardians


## Notes

NB: Guardians cannot be any membership status except Electorate. Their guardianship must be removed before their status can be changed!

NB: Tally Results must be stored in the Membership keeper too, because
they won't make sense in the normal gov sense.

### Calculating the Result

Discord Link: https://discord.com/channels/@me/1043251283575967835/1119365159455047812

```
NormalMember_Amount = Members_Amount - Guardian_Amount
Guardian_Power = GuardianWeight / Guardian_Amount
NormalMember_Power = (1-GuardianWeight) / NormalMember_Amount

VotePortion = Guardian_Votes * Guardian_Power + NormalMember_Votes * NormalMember_Power
CombinedVotingPower = VotePortion

YesPortion = (Guardian_YesVotes / Guardian_Votes) * Guardian_Power + (NormalMember_YesVotes / NormalMember_Votes) * NormalMember_Power

VetoPortion = (Guardian_NoWithVetoVotes / (Guardian_Votes - Guardian_AbstainVotes)) * Guardian_Power + (NormalMember_NoWithVetoVotes / (NormalMember_Votes - NormalMember_AbstainVotes)) * NormalMember_Power

ChorumReached = VotePortion > 1/3
or, ChorumReached = CombinedVotingPower > 1/3

ProposalPassed = YesPortion > 0.5 && VetoPortion < 1/3
```

```
NormalMember_Amount = Members_Amount - Guardian_Amount
Guardian_Power = GuardianWeight / Guardian_Amount
NormalMember_Power = (1-GuardianWeight) / NormalMember_Amount

VotePortion = Guardian_Votes * Guardian_Power + NormalMember_Votes * NormalMember_Power

YesPortion = (Guardian_YesVotes / Guardian_Votes) * Guardian_Power + (NormalMember_YesVotes / NormalMember_Votes) * NormalMember_Power

VetoPortion = (Guardian_NoWithVetoVotes / (Guardian_Votes - Guardian_AbstainVotes)) * Guardian_Power + (NormalMember_NoWithVetoVotes / (NormalMember_Votes - NormalMember_AbstainVotes)) * NormalMember_Power

ChorumReached = VotePortion > 1/3

ProposalPassed = YesPortion > 0.5 && VetoPortion < 1/3
```


## ChatGPT's auto-generated pseudocode

```
initialize YesVotes to 0
initialize NoVotes to 0
initialize NoWithVetoVotes to 0
initialize AbstainVotes to 0
initialize GuardianYesVotes to 0
initialize GuardianNoVotes to 0
initialize GuardianNoWithVetoVotes to 0
initialize GuardianAbstainVotes to 0
initialize totalNumberOfVoters to the total number of voters in the electorate
initialize totalNumberOfGuardians to the total number of guardians in the electorate
initialize totalVotingWeight to the pre-determined fixed voting weight
initialize guardianVotingWeight to totalVotingWeight / totalNumberOfGuardians
initialize quorumPercentage to the pre-determined quorum percentage
initialize result to "Quorum Not Met"

for each voter in the list of voters
    if voter is a guardian
        if voter vote is "Yes"
            increment GuardianYesVotes by 1
        else if voter vote is "No"
            increment GuardianNoVotes by 1
        else if voter vote is "No With Veto"
            increment GuardianNoWithVetoVotes by 1
        else
            increment GuardianAbstainVotes by 1
    else
        if voter vote is "Yes"
            increment YesVotes by 1
        else if voter vote is "No"
            increment NoVotes by 1
        else if voter vote is "No With Veto"
            increment NoWithVetoVotes by 1
        else
            increment AbstainVotes by 1

calculate totalYesVotes as YesVotes + (GuardianYesVotes * guardianVotingWeight)
calculate totalNoVotes as NoVotes + (GuardianNoVotes * guardianVotingWeight) + NoWithVetoVotes + (GuardianNoWithVetoVotes * guardianVotingWeight)
calculate totalVotes as totalNumberOfVoters - AbstainVotes - GuardianAbstainVotes
calculate totalVetoVotes as NoWithVetoVotes + (GuardianNoWithVetoVotes * guardianVotingWeight)
calculate totalParticipation as totalYesVotes + totalNoVotes

if totalParticipation / totalVotingWeight >= quorumPercentage
    if (totalYesVotes / totalVotes) > 0.5 and (totalVetoVotes / totalVotes) < 1/3
        set result to "Proposal Passed"
    else if (totalVetoVotes / totalVotes) >= 1/3
        set result to "Proposal Vetoed"
    else
        set result to "Proposal Failed"

print result
```

### Explanation

The voting system is used to determine if a proposal will pass or not. Each member of an electorate has the ability to vote. They can cast one of four votes: Yes, No, No with Veto, or Abstain.

However, there is a special class of members known as guardians who have more influence in the voting process. Their votes carry more weight than a normal member's vote. The weight of their votes is calculated based on a predetermined total voting weight (which is a percentage) and the number of guardians.

This system allows for majority rule, but with a twist: guardians can use their amplified voting power to help guide the process, especially in the early stages. Eventually, they may choose to step back and let the process continue with one vote per member.

In the pseudocode:

We first initialize several variables to track the number of each type of vote (Yes, No, No with Veto, and Abstain), both for regular members and for guardians.

We then loop over each voter in the electorate. Depending on the voter's type and their vote, we increment the appropriate vote count.

After all votes have been counted, we calculate the total voting power of Yes votes and No votes, considering the extra weight of guardian votes.

We also calculate the total number of votes excluding Abstain votes, and the total voting power for the "No with Veto" votes.

The voting system also requires that a certain minimum percentage of the total voting power (the quorum) participates in the vote for it to be valid. We calculate this total participation and check against the quorum.

If the quorum is met, we then determine the result of the vote:

If more than half of the non-abstaining votes were Yes votes and less than one third were "No with Veto" votes, then the proposal passes.
If one third or more of the non-abstaining votes were "No with Veto" votes, then the proposal is vetoed.
If neither of these conditions are met, then the proposal fails.
The result is then printed out.

This system allows for nuanced decision making, with the power to veto proposals, and special weighted voting power for certain members. It also ensures that decisions are only made when a sufficient proportion of the total voting power participates, providing democratic legitimacy.



