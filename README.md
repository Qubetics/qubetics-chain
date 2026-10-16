
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

## Security

Please report security vulnerabilities to our security team at security@qubetics.io.

```bash
./local_node.sh
```

