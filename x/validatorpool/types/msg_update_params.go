// In your types/msg_update_params.go file (or wherever you implement message methods)
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateParams = "update_params"

var _ sdk.Msg = &MsgUpdateParams{}

// Route returns the name of the module
func (msg *MsgUpdateParams) Route() string {
	return RouterKey
}

// Type returns the action
func (msg *MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}

// GetSigners returns the expected signers for the message
func (msg *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

// GetSignBytes returns the message bytes to sign over
func (msg *MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validates the message
func (msg *MsgUpdateParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	return msg.Params.Validate()
}

// NewMsgUpdateParams creates a new MsgUpdateParams instance
func NewMsgUpdateParams(authority string, params Params) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: authority,
		Params:    params,
	}
}
