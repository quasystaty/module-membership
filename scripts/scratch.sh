#!/bin/bash

ADDRESS_ME=mm1dmh80jwx0mv5khvqdz9sj28dmuhvems97wq628
PROPOSAL_TEXT="proposal_text.json"
GAS_PRICE_DENOM="ucrd"
GAS_PRICE="0.0025"

# Scratch area of commands used during membership and governance testing

# Become a member of the denom
membershipd tx membership enroll \
    --from me \
    --gas-prices 0.0025ucrd \
    --gas auto \
    --gas-adjustment 1.5 \
    --yes

# Get me account membership
#membershipd query membership get-member $ADDRESS_ME \
#    --output json | jq --color-output

# submit a text proposal
membershipd tx gov submit-proposal \
    $PROPOSAL_TEXT \
    --from me \
    --gas-prices $GAS_PRICE$GAS_PRICE_DENOM \
    --gas auto \
    --gas-adjustment 1.5 \
    --yes

# list all proposals
membershipd query gov proposals \
    --output json | jq --color-output

PROPOSAL_ID=1

# deposit to a proposal
membershipd tx gov deposit \
    $PROPOSAL_ID 1000000unoria \ 
    --from me \
    --gas-prices 0.0025ucrd \
    --gas auto \
    --gas-adjustment 1.5 \
    --yes

# vote on a proposal
membershpid tx gov vote \
    $PROPOSAL_ID Yes \
    --from me \
    --gas-prices 0.0025ucrd \
    --gas auto \
    --gas-adjustment 1.5 \
    --yes

# params-change - add guardian wallet to params
membershipd 