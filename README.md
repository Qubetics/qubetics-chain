
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

The EVM implementation is [Cosmos EVM](https://github.com/cosmos/evm) v0.1.0,
backported from Cosmos SDK v0.50 to Cosmos SDK v0.47.12 so it can run alongside
the Qubetics dVPN modules. Reusable EVM libraries live at the repository root,
and the Qubetics chain application lives in `qubeticsd/` (the equivalent of
Cosmos EVM's `example_chain/`).

| Path | Contents |
|------|----------|
| `qubeticsd/` | Chain application (`app.go`, keepers wiring, genesis, export, precompiles) |
| `qubeticsd/ante/` | Ante handler wiring for Cosmos and Ethereum transactions |
| `qubeticsd/cmd/qubeticsd/` | `qubeticsd` binary entrypoint and CLI commands |
| `qubeticsd/cmd/qubeticsd/config/` | Bech32 prefixes and denominations |
| `ante/` | Cosmos EVM ante handlers (`ante/evm`, `ante/cosmos`) |
| `x/vm/` | EVM module (module name `evm`) |
| `x/erc20/`, `x/feemarket/`, `x/ibc/` | EVM-related modules |
| `x/vpn/`, `x/node/`, `x/provider/`, `x/plan/`, `x/session/`, `x/subscription/`, `x/lease/`, `x/deposit/` | dVPN modules |
| `x/mint/`, `x/validatorpool/` | Chain-specific modules |
| `precompiles/` | EVM precompiled contracts |
| `rpc/`, `server/`, `indexer/` | JSON-RPC server and transaction indexer |
| `crypto/`, `encoding/`, `ethereum/eip712/`, `wallets/` | Keys, codecs, EIP-712 and hardware wallets |
| `types/`, `utils/`, `version/` | Shared types and helpers |
| `proto/` | Protobuf definitions (`cosmos/evm/*`, `qubetics/*`) |
| `client/docs/` | Swagger UI, generated `swagger.yaml` and embedded (statik) docs |
| `testutil/` | Test helpers |

## Development

| Command | Description |
|---------|-------------|
| `make build` | Build `qubeticsd` into `./build` |
| `make install` | Install `qubeticsd` into `$GOPATH/bin` |
| `make test-unit` | Run unit tests |
| `make lint` | Run linters |
| `make proto-gen` | Regenerate Go code from `proto/` (requires Docker) |
| `make proto-swagger-gen` | Regenerate `client/docs/swagger-ui/swagger.yaml` and re-embed it (requires Docker) |
| `make update-swagger-docs` | Re-embed `client/docs/swagger-ui` into `client/docs/statik` |
| `make swagger` | Alias for `make proto-swagger-gen` |
| `make proto-all` | Format, lint and regenerate protobuf code and the Swagger spec |

The REST API serves the Swagger UI at `/swagger/` when `swagger = true` is set in the
`[api]` section of `app.toml` (enabled by `local_node.sh`).

## License

The EVM implementation in this repository is based on
[Cosmos EVM](https://github.com/cosmos/evm) v0.1.0, licensed under the
[Apache License 2.0](LICENSE). It was imported and then backported to Cosmos SDK
v0.47.12; see [NOTICE](NOTICE) and the git history for details.

Code derived from [go-ethereum](https://github.com/ethereum/go-ethereum) keeps its
original GNU LGPL v3.0 license notices.

Earlier versions of this repository were based on [Evmos](https://github.com/evmos/evmos)
and are subject to the [Evmos Non-Commercial License](https://github.com/evmos/evmos/blob/main/LICENSE).
That code has been removed from this version.

## Security

Please report security vulnerabilities to our security team at security@qubetics.io.

