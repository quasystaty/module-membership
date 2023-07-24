#!/bin/bash

###############################
# Usage: tx-gas.sh <cosmos tx...>
###############################
# This script is used to submit a transaction to the blockchain
# and wait for it to be included in a block.
#
# It also makes assumptions about the gas prices and gas adjustment
# for your convenience
###############################

if [ $# -eq 0 ]; then
  echo "Usage: gas-tx.sh <cosmos tx...>"
  exit 1
fi

GAS_PRICE="0.0025"
GAS_PRICE_DENOM="ucrd"

./tx.sh "$@" --gas-prices $GAS_PRICE$GAS_PRICE_DENOM --gas auto --gas-adjustment 2.5

# Exit with the same exit code as tx.sh
if [ $? -ne 0 ]; then
  exit 1
fi
