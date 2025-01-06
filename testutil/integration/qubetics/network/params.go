package network

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
	feemarketypes "github.com/qubetics/qubetics-blockchain/v2/x/feemarket/types"
)

func (n *IntegrationNetwork) UpdateEvmParams(params evmtypes.Params) error {
	return n.app.EvmKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateFeeMarketParams(params feemarketypes.Params) error {
	return n.app.FeeMarketKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateMintParams(params minttypes.Params) error {
	return n.app.MintKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateGovParams(params govtypes.Params) error {
	return n.app.GovKeeper.SetParams(n.ctx, params)
}
