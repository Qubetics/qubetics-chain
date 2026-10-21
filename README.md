
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

## License

> [!IMPORTANT]
> **Parts of this codebase are derived from [Evmos](https://github.com/evmos/evmos) v19 and are not licensed for commercial use.**

Files that begin with the following header are derived from Evmos and are licensed under the
[Evmos Non-Commercial License (ENCL-1.0)](https://github.com/evmos/evmos/blob/main/LICENSE), owned by Tharsis Labs Ltd.:

```go
// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
```

Under that license, these files:

- may only be used for **non-commercial purposes** (academic, scientific, research and development, or evaluation such as auditing);
- **may not be used to launch a publicly available blockchain**, precompiled smart contracts, or other distributed-ledger system that facilitates transactions of economic value;
- may not be used commercially without permission from Tharsis Labs Ltd.

The Apache-2.0 [LICENSE](LICENSE) file in this repository does **not** apply to these Evmos-derived files.
Files carrying other headers (for example go-ethereum's LGPL-3.0 notice) remain under their original licenses.

If you plan to use this code, please review the
[Evmos license](https://github.com/evmos/evmos/blob/main/LICENSE) and its
[FAQ](https://github.com/evmos/evmos/blob/main/LICENSE_FAQ.md) first.

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

## Security

Please report security vulnerabilities to our security team at security@qubetics.io.

```bash
./local_node.sh
```

