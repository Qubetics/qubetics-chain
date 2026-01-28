package v3

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
)

func NewGenesisState(sessions []Session, params Params) *GenesisState {
	anySession := make([]types.Any, len(sessions))
	for i, session := range sessions {
		bz, err := types.NewAnyWithValue(session)
		if err != nil {
			panic(err)
		}
		anySession[i] = *bz
	}
	return &GenesisState{
		Sessions: anySession,
		Params:   params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	// Validate sessions
	if len(state.Sessions) == 0 {
		return nil // No sessions to validate
	}

	// =====Check each session
	for i, anySession := range state.Sessions {
		// The sessions are wrapped in protobuf Any types. Their complete validation
		// (ID uniqueness, field validity, constraints) is performed when they are
		// unmarshaled and persisted by the keeper's InitGenesis function.
		// Here we perform basic structural validation of the Any wrapper.
		if anySession.TypeUrl == "" {
			return fmt.Errorf("session at index %d has empty TypeUrl", i)
		}

		if len(anySession.Value) == 0 {
			return fmt.Errorf("session at index %d has empty value", i)
		}
	}

	return nil
}
