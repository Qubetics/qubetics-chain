package contracts

import (
	contractutils "github.com/qubetics/qubetics-blockchain/v2/contracts/utils"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/vm/types"
)

func LoadReverterContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("Reverter.json")
}
