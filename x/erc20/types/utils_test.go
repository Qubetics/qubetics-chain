package types_test

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"
	"github.com/stretchr/testify/require"
)

func TestSanitizeERC20Name(t *testing.T) {
	testCases := []struct {
		name         string
		erc20Name    string
		expErc20Name string
		expectPass   bool
	}{
		{"name contains 'Special Characters'", "*Special _ []{}||*¼^%  &Token", "SpecialToken", true},
		{"name contains 'Special Numbers'", "*20", "20", false},
		{"name contains 'Spaces'", "   Spaces   Token", "SpacesToken", true},
		{"name contains 'Leading Numbers'", "12313213  Number     Coin", "NumberCoin", true},
		{"name contains 'Numbers in the middle'", "  Other    Erc20 Coin ", "OtherErc20Coin", true},
		{"name contains '/'", "USD/Coin", "USD/Coin", true},
		{"name contains '/'", "/SlashCoin", "SlashCoin", true},
		{"name contains '/'", "O/letter", "O/letter", true},
		{"name contains '/'", "Ot/2letters", "Ot/2letters", true},
		{"name contains '/'", "ibc/valid", "valid", true},
		{"name contains '/'", "erc20/valid", "valid", true},
		{"name contains '/'", "ibc/erc20/valid", "valid", true},
		{"name contains '/'", "ibc/erc20/ibc/valid", "valid", true},
		{"name contains '/'", "ibc/erc20/ibc/20invalid", "20invalid", false},
		{"name contains '/'", "123/leadingslash", "leadingslash", true},
		{"name contains '-'", "Dash-Coin", "Dash-Coin", true},
		{"really long word", strings.Repeat("a", 150), strings.Repeat("a", 128), true},
		{"single word name: Token", "Token", "Token", true},
		{"single word name: Coin", "Coin", "Coin", true},
	}

	for _, tc := range testCases {
		name := types.SanitizeERC20Name(tc.erc20Name)
		require.Equal(t, tc.expErc20Name, name, tc.name)
		err := sdk.ValidateDenom(name)
		if tc.expectPass {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}

func TestEqualMetadata(t *testing.T) {
	testCases := []struct {
		name      string
		metadataA banktypes.Metadata
		metadataB banktypes.Metadata
		expError  bool
	}{
		{
			"equal metadata",
			banktypes.Metadata{
				Base:        "tics",
				Display:     "qubetics",
				Name:        "Qubetics",
				Symbol:      "QUBETICS",
				Description: "EVM, staking and governance denom of Qubetics",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "tics",
						Exponent: 0,
						Aliases:  []string{"atto qubetics"},
					},
					{
						Denom:    "qubetics",
						Exponent: 18,
					},
				},
			},
			banktypes.Metadata{
				Base:        "tics",
				Display:     "qubetics",
				Name:        "Qubetics",
				Symbol:      "QUBETICS",
				Description: "EVM, staking and governance denom of Qubetics",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "tics",
						Exponent: 0,
						Aliases:  []string{"atto qubetics"},
					},
					{
						Denom:    "qubetics",
						Exponent: 18,
					},
				},
			},
			false,
		},
		{
			"different base field",
			banktypes.Metadata{
				Base: "tics",
			},
			banktypes.Metadata{
				Base: "ttics",
			},
			true,
		},
		{
			"different denom units length",
			banktypes.Metadata{
				Base:        "tics",
				Display:     "qubetics",
				Name:        "Qubetics",
				Symbol:      "QUBETICS",
				Description: "EVM, staking and governance denom of Qubetics",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "tics",
						Exponent: 0,
						Aliases:  []string{"atto qubetics"},
					},
					{
						Denom:    "qubetics",
						Exponent: 18,
					},
				},
			},
			banktypes.Metadata{
				Base:        "tics",
				Display:     "qubetics",
				Name:        "Qubetics",
				Symbol:      "QUBETICS",
				Description: "EVM, staking and governance denom of Qubetics",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "tics",
						Exponent: 0,
						Aliases:  []string{"atto qubetics"},
					},
				},
			},
			true,
		},
		{
			"different denom units",
			banktypes.Metadata{
				Base:        "tics",
				Display:     "qubetics",
				Name:        "Qubetics",
				Symbol:      "QUBETICS",
				Description: "EVM, staking and governance denom of Qubetics",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "tics",
						Exponent: 0,
						Aliases:  []string{"atto qubetics"},
					},
					{
						Denom:    "uqubetics",
						Exponent: 12,
						Aliases:  []string{"micro qubetics"},
					},
					{
						Denom:    "qubetics",
						Exponent: 18,
					},
				},
			},
			banktypes.Metadata{
				Base:        "tics",
				Display:     "qubetics",
				Name:        "Qubetics",
				Symbol:      "QUBETICS",
				Description: "EVM, staking and governance denom of Qubetics",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "tics",
						Exponent: 0,
						Aliases:  []string{"atto qubetics"},
					},
					{
						Denom:    "Uqubetics",
						Exponent: 12,
						Aliases:  []string{"micro qubetics"},
					},
					{
						Denom:    "qubetics",
						Exponent: 18,
					},
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		err := types.EqualMetadata(tc.metadataA, tc.metadataB)
		if tc.expError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestEqualAliases(t *testing.T) {
	testCases := []struct {
		name     string
		aliasesA []string
		aliasesB []string
		expEqual bool
	}{
		{
			"empty",
			[]string{},
			[]string{},
			true,
		},
		{
			"different lengths",
			[]string{},
			[]string{"atto qubetics"},
			false,
		},
		{
			"different values",
			[]string{"attoqubetics"},
			[]string{"atto qubetics"},
			false,
		},
		{
			"same values, unsorted",
			[]string{"atto qubetics", "tics"},
			[]string{"tics", "atto qubetics"},
			false,
		},
		{
			"same values, sorted",
			[]string{"tics", "atto qubetics"},
			[]string{"tics", "atto qubetics"},
			true,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expEqual, types.EqualStringSlice(tc.aliasesA, tc.aliasesB), tc.name)
	}
}
