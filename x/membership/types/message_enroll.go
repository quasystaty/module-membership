package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnroll = "enroll"

var _ sdk.Msg = &MsgEnroll{}

func NewMsgEnroll(creator string, nickname string) *MsgEnroll {
	return &MsgEnroll{
		Creator:  creator,
		Nickname: nickname,
	}
}

func (msg *MsgEnroll) Route() string {
	return RouterKey
}

func (msg *MsgEnroll) Type() string {
	return TypeMsgEnroll
}

func (msg *MsgEnroll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnroll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnroll) ValidateBasic() error {
	// Validate creator address
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// Nickname cannot be longer than 30 characters
	if len(msg.Nickname) > NicknameMaxLength {
		return errors.Wrapf(ErrInvalidNickname, "nickname cannot be longer than %d characters", NicknameMaxLength)
	}
	return nil
}
