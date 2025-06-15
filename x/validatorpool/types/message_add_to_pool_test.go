package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/qubetics/qubetics-blockchain/v2/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddToPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddToPool
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgAddToPool{
				Creator: "invalid_address",
				Amount:  "1000000stake",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty creator address",
			msg: MsgAddToPool{
				Creator: "",
				Amount:  "1000000stake",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty amount",
			msg: MsgAddToPool{
				Creator: sample.AccAddress(),
				Amount:  "",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid amount format",
			msg: MsgAddToPool{
				Creator: sample.AccAddress(),
				Amount:  "invalid-amount",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "multiple coins not allowed",
			msg: MsgAddToPool{
				Creator: sample.AccAddress(),
				Amount:  "1000000stake,2000000token",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "zero amount",
			msg: MsgAddToPool{
				Creator: sample.AccAddress(),
				Amount:  "0stake",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "negative amount",
			msg: MsgAddToPool{
				Creator: sample.AccAddress(),
				Amount:  "-1000000stake",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "valid message",
			msg: MsgAddToPool{
				Creator: sample.AccAddress(),
				Amount:  "1000000stake",
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
