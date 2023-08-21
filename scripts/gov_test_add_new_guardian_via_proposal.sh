#!/bin/bash

echo "GOV TEST: Everyone votes yes, Proposal passes"

# Me is a regular member
ADDRESS_ME=mm1dmh80jwx0mv5khvqdz9sj28dmuhvems97wq628
# VAL1 is a guardian
ADDRESS_VAL1=mm1e7gp56hf85nk0qtg0542gmmmwq753ww2tg7dws
PROPOSAL_TEXT="proposal_text.json"

// Create a proposal JSON file and write it to a temporary file
cat > $PROPOSAL_TEXT <<EOF
{
  "title": "Test Proposal",
  "description": "This is a test proposal",
  "type": "Text",
  "deposit": "1000000unoria"
}
EOF
