package contracts

import (
	contractutils "github.com/qubetics/qubetics-blockchain/v2/contracts/utils"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
)

func LoadInterchainSenderCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("InterchainSenderCaller.json")
}
