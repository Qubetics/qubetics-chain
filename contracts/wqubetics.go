package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
)

var (
	//go:embed compiled_contracts/WQUBETICS.json
	WQUBETICSJSON []byte

	// WQUBETICSContract is the compiled contract of WQUBETICS
	WQUBETICSContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WQUBETICSJSON, &WQUBETICSContract)
	if err != nil {
		panic(err)
	}

	if len(WQUBETICSContract.Bin) == 0 {
		panic("failed to load WQUBETICS smart contract")
	}
}
