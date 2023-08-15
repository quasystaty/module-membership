package types

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDirectDemocracy = "update_direct_democracy"

var _ sdk.Msg = &MsgUpdateDirectDemocracy{}

func NewMsgUpdateDirectDemocracy(creator string, addGuardians []string, removeGuardians []string, totalVotingWeight string) *MsgUpdateDirectDemocracy {
	if totalVotingWeight != "" {
		tvw := math.LegacyMustNewDecFromStr(totalVotingWeight)
		return &MsgUpdateDirectDemocracy{
			Creator:           creator,
			AddGuardians:      addGuardians,
			RemoveGuardians:   removeGuardians,
			TotalVotingWeight: &tvw,
		}
	} else {
		return &MsgUpdateDirectDemocracy{
			Creator:         creator,
			AddGuardians:    addGuardians,
			RemoveGuardians: removeGuardians,
		}
	}
}

func (msg *MsgUpdateDirectDemocracy) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDirectDemocracy) Type() string {
	return TypeMsgUpdateDirectDemocracy
}

func (msg *MsgUpdateDirectDemocracy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDirectDemocracy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDirectDemocracy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Make sure total voting weight is between 0 and 1, inclusively
	if msg.TotalVotingWeight != nil {
		if msg.TotalVotingWeight.LT(sdk.ZeroDec()) || msg.TotalVotingWeight.GT(sdk.OneDec()) {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "total voting weight must be between 0 and 1, inclusively")
		}
	}

	// There must be at least *one* guardian to add or remove, or total voting weight set
	if len(msg.AddGuardians) == 0 && len(msg.RemoveGuardians) == 0 && msg.TotalVotingWeight == nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "nothing to do")
	}

	// All guardian addresses must be valid
	if msg.AddGuardians != nil {
		for _, add := range msg.AddGuardians {
			_, err := sdk.AccAddressFromBech32(add)
			if err != nil {
				return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid add guardian address (%s)", err)
			}
		}
	}
	if msg.RemoveGuardians != nil {
		for _, remove := range msg.RemoveGuardians {
			_, err := sdk.AccAddressFromBech32(remove)
			if err != nil {
				return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid remove guardian address (%s)", err)
			}
		}
	}

	// Cannot have an address in both the add and the remove list
	if msg.AddGuardians != nil && msg.RemoveGuardians != nil {
		for _, add := range msg.AddGuardians {
			for _, remove := range msg.RemoveGuardians {
				if add == remove {
					return errors.Wrapf(sdkerrors.ErrInvalidRequest, "cannot add and remove the same guardian (%s)", add)
				}
			}
		}
	}

	return nil
}
