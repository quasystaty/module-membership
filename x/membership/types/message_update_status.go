package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateStatus = "update_status"

var _ sdk.Msg = &MsgUpdateStatus{}

func NewMsgUpdateStatus(creator string, address string, status MembershipStatus) *MsgUpdateStatus {
	return &MsgUpdateStatus{
		Creator: creator,
		Address: address,
		Status:  status,
	}
}

func (msg *MsgUpdateStatus) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStatus) Type() string {
	return TypeMsgUpdateStatus
}

func (msg *MsgUpdateStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStatus) ValidateBasic() error {
	// Validate creator address
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// Validate target address
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid target address (%s)", err)
	}
	// Validate status
	if !msg.Status.IsValid() {
		return errors.Wrapf(ErrInvalidMembershipStatus, "invalid status (%s)", msg.Status)
	}
	return nil
}
