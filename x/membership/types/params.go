package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/noria-net/module-membership/util"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyGuardians                    = []byte("Guardians")
	DefaultGuardians         string = ""
	KeyTotalVotingWeight            = []byte("TotalVotingWeight")
	DefaultTotalVotingWeight        = sdk.ZeroDec()
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(guardians string, totalVotingWeight sdk.Dec) Params {
	return Params{
		Guardians:         guardians,
		TotalVotingWeight: totalVotingWeight,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultGuardians, DefaultTotalVotingWeight)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyGuardians, &p.Guardians, validateGuardians),
		paramtypes.NewParamSetPair(KeyTotalVotingWeight, &p.TotalVotingWeight, validateTotalVotingWeight),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateGuardians(p.Guardians); err != nil {
		return err
	}
	if err := validateTotalVotingWeight(p.TotalVotingWeight); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateGuardians(v interface{}) error {
	// validate the data type
	guardians, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// validate the list of guardians
	_, err := util.SplitStringIntoAddresses(guardians)
	if err != nil {
		return fmt.Errorf("invalid guardian addresses")
	}

	return nil
}

// validate the total voting weight
func validateTotalVotingWeight(i interface{}) error {
	// validate the data type
	totalVotingWeight, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// totalVotingWeight must be between 0 and 100, inclusive
	if totalVotingWeight.LT(sdk.ZeroDec()) || totalVotingWeight.GT(sdk.NewDec(100)) {
		return fmt.Errorf("total voting weight must be between 0 and 100, inclusive: %s", totalVotingWeight)
	}

	return nil
}
