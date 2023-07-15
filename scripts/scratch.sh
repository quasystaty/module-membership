#!/bin/bash

ADDRESS_ME=mm1dmh80jwx0mv5khvqdz9sj28dmuhvems97wq628
PROPOSAL_TEXT="proposal_text.json"
GAS_PRICE_DENOM="ucrd"
GAS_PRICE="0.0025"

# Scratch area of commands used during membership and governance testing

# Become a member of the denom

#membershipd tx membership enroll \
#    --from me \
#    --gas-prices 0.0025ucrd \
#    --gas auto \
#    --gas-adjustment 1.5 \
#    --yes

# Get me account membership
#membershipd query membership get-member $ADDRESS_ME \
#    --output json | jq --color-output

# submit a text proposal
echo "Submitting proposal"
membershipd tx gov submit-proposal \
    $PROPOSAL_TEXT \
    --from val1 \
    --gas-prices $GAS_PRICE$GAS_PRICE_DENOM \
    --gas auto \
    --gas-adjustment 1.5 \
    --yes

sleep 2

#membershipd query gov proposals \
#    --output json | jq --color-output

# Get the ID of the  latest proposal
PROPOSAL_ID=$(membershipd query gov proposals --output json --reverse --limit 1 | jq -r '.proposals[].id')

# deposit to a proposal
echo "Depositing to proposal $PROPOSAL_ID"
membershipd tx gov deposit $PROPOSAL_ID 1000000unoria \
    --from val1 \
    --gas-prices 0.0025ucrd \
    --gas auto \
    --gas-adjustment 1.5 \
    --yes

sleep 2

# vote on a proposal
echo "Voting on proposal $PROPOSAL_ID"
membershipd tx gov vote \
    $PROPOSAL_ID Yes \
    --from val1 \
    --gas-prices 0.0025ucrd \
    --gas auto \
    --gas-adjustment 1.5 \
    --yes

sleep 2

# params-change - add guardian wallet to params
# membershipd 