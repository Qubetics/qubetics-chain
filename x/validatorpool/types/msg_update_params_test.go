package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateParams(t *testing.T) {
	// Test data
	validAuthority := "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"
	invalidAuthority := "invalid-address"

	validParams := Params{
		// Add your params fields here
	}

	tests := []struct {
		name        string
		msg         *MsgUpdateParams
		expectError bool
	}{
		{
			name:        "valid message",
			msg:         NewMsgUpdateParams(validAuthority, validParams),
			expectError: false,
		},
		{
			name:        "invalid authority address",
			msg:         NewMsgUpdateParams(invalidAuthority, validParams),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Route
			require.Equal(t, RouterKey, tt.msg.Route())

			// Test Type
			require.Equal(t, TypeMsgUpdateParams, tt.msg.Type())

			// Test GetSigners
			if !tt.expectError {
				signers := tt.msg.GetSigners()
				require.Len(t, signers, 1)
				expectedSigner, err := sdk.AccAddressFromBech32(validAuthority)
				require.NoError(t, err)
				require.Equal(t, expectedSigner, signers[0])
			}

			// Test GetSignBytes
			signBytes := tt.msg.GetSignBytes()
			require.NotNil(t, signBytes)

			// Test ValidateBasic
			err := tt.msg.ValidateBasic()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNewMsgUpdateParams(t *testing.T) {
	authority := "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"
	params := Params{
		// Add your params fields here
	}

	msg := NewMsgUpdateParams(authority, params)
	require.NotNil(t, msg)
	require.Equal(t, authority, msg.Authority)
	require.Equal(t, params, msg.Params)
}
