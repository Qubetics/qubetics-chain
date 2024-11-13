package testdata

import (
	contractutils "github.com/qubetics/qubetics-blockchain/v2/contracts/utils"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
)

func LoadStakingCallerTwoContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("StakingCallerTwo.json")
}
