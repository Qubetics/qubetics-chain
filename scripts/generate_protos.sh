#!/usr/bin/env bash

# Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
# Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
# SPDX-License-Identifier: Apache-2.0
#
# Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
# the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

# Generates the Go code (gogoproto + grpc-gateway) for the protobuf files in ./proto.
# Based on the Cosmos EVM v0.1.0 script (Apache-2.0).
#
# Run through the Makefile (make proto-gen), or locally with:
# docker run --network host --rm -v $(pwd):/workspace --workdir /workspace ghcr.io/cosmos/proto-builder:0.14.0 sh ./scripts/generate_protos.sh
#
set -eo pipefail

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
	proto_files=$(find "${dir}" -maxdepth 1 -name '*.proto')
	for file in $proto_files; do
		# Only generate files that belong to this Go module
		if grep -q "option go_package.*qubetics/qubetics-blockchain" "$file"; then
			buf generate --template proto/buf.gen.gogo.yaml "$file"
		fi
	done
done

# move the generated files to the right places
cp -r github.com/qubetics/qubetics-blockchain/v2/* ./
rm -rf github.com
