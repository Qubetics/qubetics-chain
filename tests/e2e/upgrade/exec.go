package upgrade

import (
	"fmt"
)

// E2ETxArgs contains the arguments to build a CLI transaction command from.
type E2ETxArgs struct {
	ModuleName string
	SubCommand string
	Args       []string
	ChainID    string
	From       string
}

// CreateModuleTxExec creates the execution command for an Qubetics transaction.
func (m *Manager) CreateModuleTxExec(txArgs E2ETxArgs) (string, error) {
	cmd := []string{
		"qubeticsd",
		"tx",
		txArgs.ModuleName,
		txArgs.SubCommand,
	}
	cmd = append(cmd, txArgs.Args...)
	cmd = append(cmd,
		fmt.Sprintf("--chain-id=%s", txArgs.ChainID),
		"--keyring-backend=test",
		"--output=json",
		"--fees=500000000000tics",
		"--gas=auto",
		fmt.Sprintf("--from=%s", txArgs.From),
		"--yes",
	)
	return m.CreateExec(cmd, m.ContainerID())
}
