package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/qubetics/qubetics-blockchain/v2/x/erc20/client/cli"
)

var (
	RegisterERC20ProposalHandler         = govclient.NewProposalHandler(cli.NewRegisterERC20ProposalCmd)
	ToggleTokenConversionProposalHandler = govclient.NewProposalHandler(cli.NewToggleTokenConversionProposalCmd)
)
