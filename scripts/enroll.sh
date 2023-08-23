#! /bin/bash

# This script is used to enroll a new member to the membershipd module
# It accepts the user's address book key as an argument
# It accepts a second argument called "ignore" which will ignore the error if the address is already enrolled

# Some presets
DAEMON_BINARY=membershipd

#  Input argument
KEY=$1
IGNORE_ALREADY_ENROLLED=$2

# Validate the input and print a help message if the address is missing
if [ -z "$KEY" ]; then
    echo "Usage: $0 <key> --ignore"
    echo "Example: $0 me"
    echo "Example: $0 val1 --ignore"
    exit 1
fi

# Get the address from the address book and save it to ADDRESS
ADDRESS=$($DAEMON_BINARY keys show $KEY --address)
if [ $? -eq 1 ]; then
    echo "ERROR: $KEY not found in address book"
    exit 1
fi

# Try and enroll the member
# If the address is already enrolled, the script will fail unless the IGNORE_ALREADY_ENROLLED argument is set
$DAEMON_BINARY query membership member $ADDRESS --output json | grep "member not found" > /dev/null 2>&1
# Only exit if the return code is non-zero and --ignore is not set
if [ $? -eq 1 ]; then
    if [ -z "$IGNORE_ALREADY_ENROLLED" ]; then
        echo "ERROR: $ADDRESS is already enrolled"
        echo "Use the '--ignore' argument to enroll anyway"
        exit 1
    fi
    # Exit silently, nothing to do
    exit 0
fi

# Enroll the address
echo "Enrolling $ADDRESS"
./tx-gas.sh $DAEMON_BINARY tx membership enroll --from $KEY --yes
