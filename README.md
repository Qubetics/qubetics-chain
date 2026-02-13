
<!--
parent:
  order: false
-->

<div align="center">
  <h1> Qubetics Chain </h1>
</div>

## About

Qubetics is a next-generation blockchain platform that combines the best features of Cosmos and Ethereum ecosystems. Built on the Cosmos SDK and CometBFT consensus engine, Qubetics offers:

- **EVM Compatibility**: Full Ethereum Virtual Machine support for seamless integration with Ethereum dApps
- **Cosmos Interoperability**: Seamless cross-chain communication using IBC
- **High Performance**: Optimized for high throughput and low latency
- **Proof-of-Stake Security**: Secure consensus mechanism with validator staking

## Features

- **Smart Contract Execution**
  - Ethereum Virtual Machine (EVM) compatibility
  - Cross-chain contract interoperability

- **Interoperability**
  - Inter-Blockchain Communication (IBC) protocol
  - Ethereum bridge
  - Multi-chain asset transfers

- **Security & Governance**
  - Proof-of-Stake consensus
  - On-chain governance
  - Validator staking and slashing

## Installation

### Prerequisites

- Go 1.22.7 or later
- Docker (optional, for containerized deployment)
- Git
- Make

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/Qubetics/qubetics-chain.git
cd qubetics-chain
```

2. Build the binaries:
```bash
make build
```

3. Install the tools:
```bash
make install
```


## Getting Started

### Initialize the Node

```bash
qubeticsd init <moniker>
```

### Configure the Genesis File

```bash
qubeticsd init-genesis
```

### Start the Node

```bash
qubeticsd start
```

### Run a Local Node

```bash
./local_node.sh
```

## Repository Layout

The layout follows [cosmos/evm](https://github.com/cosmos/evm): reusable EVM
libraries live at the repository root, and the Qubetics chain application lives
in `qubeticsd/` (the equivalent of `evmd/`).

| Path | Contents |
|------|----------|
| `qubeticsd/` | Chain application (`app.go`, keepers wiring, genesis, export) |
| `qubeticsd/cmd/qubeticsd/` | `qubeticsd` binary entrypoint and CLI commands |
| `qubeticsd/cmd/qubeticsd/config/` | Bech32 prefixes and denominations |
| `qubeticsd/eips/` | Custom EIP activators |
| `qubeticsd/post/` | Post handlers |
| `qubeticsd/upgrades/` | Upgrade handlers |
| `ante/` | Ante handlers (`ante/evm`, `ante/cosmos`, `ante/utils`) |
| `x/vm/` | EVM module (module name `evm`) |
| `x/erc20/`, `x/feemarket/`, `x/ibc/` | EVM-related modules |
| `x/vpn/`, `x/node/`, `x/provider/`, `x/plan/`, `x/session/`, `x/subscription/`, `x/lease/`, `x/deposit/` | dVPN modules |
| `x/mint/`, `x/validatorpool/` | Chain-specific modules |
| `precompiles/` | EVM precompiled contracts |
| `rpc/`, `server/`, `indexer/` | JSON-RPC server and transaction indexer |
| `crypto/`, `encoding/`, `ethereum/eip712/`, `wallets/` | Keys, codecs, EIP-712 and hardware wallets |
| `types/`, `utils/`, `version/` | Shared types and helpers |
| `proto/` | Protobuf definitions (`ethermint/*`, `qubetics/*`) |
| `client/docs/` | Swagger UI and embedded (statik) docs |
| `testutil/`, `tests/` | Test helpers, integration, e2e and nix tests |

## Development

| Command | Description |
|---------|-------------|
| `make build` | Build `qubeticsd` into `./build` |
| `make install` | Install `qubeticsd` into `$GOPATH/bin` |
| `make test-unit` | Run unit tests |
| `make lint` | Run `golangci-lint` |
| `make proto-gen` | Regenerate Go code from `proto/` (requires Docker) |
| `make proto-swagger-gen` | Regenerate `client/docs/swagger-ui/swagger.json` (requires Docker) |
| `make update-swagger-docs` | Re-embed the Swagger UI into `client/docs/statik` |
| `make swagger` | `proto-swagger-gen` + `update-swagger-docs` |
| `make proto-all` | Format, lint and regenerate protobuf code and Swagger |

## Security

Please report security vulnerabilities to our security team at security@qubetics.io.

