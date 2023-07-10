package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
)

func TestMsgUpdateStatus_ValidateBasic(t *testing.T) {

	valid_1 := "cosmos1l0znsvddllw9knha3yx2svnlxny676d8ns7uys"
	valid_2 := "cosmos1j8pp7zvcu9z8vd882m284j29fn2dszh05cqvf9"
	invalid := "invalid_address"

	t.Run("invalid creator address", func(t *testing.T) {
		msg := MsgUpdateStatus{
			Creator: invalid,
		}
		err := msg.ValidateBasic()
		assert.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("invalid Target address", func(t *testing.T) {
		msg := MsgUpdateStatus{
			Creator: valid_1,
			Address: invalid,
		}
		err := msg.ValidateBasic()
		assert.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("invalid membership status", func(t *testing.T) {
		msg := MsgUpdateStatus{
			Creator: valid_1,
			Address: valid_2,
			Status:  MembershipStatus(100),
		}
		err := msg.ValidateBasic()
		assert.ErrorIs(t, err, ErrInvalidMembershipStatus)
	})

	t.Run("membership status cannot be undefined / zero", func(t *testing.T) {
		msg := MsgUpdateStatus{
			Creator: valid_1,
			Address: valid_2,
			Status:  MembershipStatus(0),
		}
		err := msg.ValidateBasic()
		assert.ErrorIs(t, err, ErrInvalidMembershipStatus)
	})
}
