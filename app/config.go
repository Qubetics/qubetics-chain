package app

import (
	"github.com/qubetics/qubetics-blockchain/v2/app/eips"
	evmconfig "github.com/qubetics/qubetics-blockchain/v2/x/evm/config"
	"github.com/qubetics/qubetics-blockchain/v2/x/evm/core/vm"
)

// The init function of the config file allows to setup the global
// configuration for the EVM, modifying the custom ones defined in evmOS.
func init() {
	err := evmconfig.NewEVMConfigurator().
		WithExtendedEips(qubeticsActivators).
		Configure()
	if err != nil {
		panic(err)
	}
}

// QubeticsActivators defines a map of opcode modifiers associated
// with a key defining the corresponding EIP.
var qubeticsActivators = map[string]func(*vm.JumpTable){
	"qubetics_0": eips.Enable0000,
	"qubetics_1": eips.Enable0001,
	"qubetics_2": eips.Enable0002,
}
