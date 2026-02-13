package post

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrKeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
)

// HandlerOptions are the options required for constructing a PostHandler.
type HandlerOptions struct {
	FeeCollectorName string
	BankKeeper       bankkeeper.Keeper
	DistrKeeper      distrKeeper.Keeper
}

func (h HandlerOptions) Validate() error {
	if h.FeeCollectorName == "" {
		return errors.New("fee collector name cannot be empty")
	}

	if h.BankKeeper == nil {
		return errors.New("bank keeper cannot be nil")
	}

	return nil
}

// NewPostHandler returns a new PostHandler decorators chain.
func NewPostHandler(ho HandlerOptions) sdk.PostHandler {
	postDecorators := []sdk.PostDecorator{
		NewBurnDecorator(ho.FeeCollectorName, ho.BankKeeper, ho.DistrKeeper),
	}

	return sdk.ChainPostDecorators(postDecorators...)
}
