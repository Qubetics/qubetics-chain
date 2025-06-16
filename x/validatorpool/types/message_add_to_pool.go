package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddToPool = "add_to_pool"

var _ sdk.Msg = &MsgAddToPool{}

func NewMsgAddToPool(creator string, amount string) *MsgAddToPool {
	return &MsgAddToPool{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgAddToPool) Route() string {
	return RouterKey
}

func (msg *MsgAddToPool) Type() string {
	return TypeMsgAddToPool
}

func (msg *MsgAddToPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddToPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddToPool) ValidateBasic() error {
	// Validate creator address
	if msg.Creator == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "creator address cannot be empty")
	}
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Validate amount
	if msg.Amount == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "amount cannot be empty")
	}
	coins, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount format: %s", err)
	}
	if len(coins) != 1 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "expected exactly one coin, got %d", len(coins))
	}
	if coins[0].Amount.IsNil() || coins[0].Amount.IsZero() || coins[0].Amount.IsNegative() {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %s", coins[0].Amount)
	}

	return nil
}
