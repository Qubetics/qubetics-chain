// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

//go:build !rocksdb
// +build !rocksdb

package config

import (
	"path/filepath"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/server/types"
)

// OpenDB opens a database based on the specified backend type.
// It takes the home directory where the database data will be stored, along with the backend type.
// It opens a database named "application" using the specified backend type and the data directory.
// It returns the opened database and an error (if any). If the database opens successfully, the error will be nil.
//
// NOTE: this is included in builds without rocksdb.
// When building the binary with rocksdb, the code in 'rocksdb.go' will be included
// instead of this
func OpenDB(_ types.AppOptions, home string, backendType dbm.BackendType) (dbm.DB, error) {
	dataDir := filepath.Join(home, "data")
	return dbm.NewDB("application", backendType, dataDir)
}

// OpenReadOnlyDB opens rocksdb backend in read-only mode.
func OpenReadOnlyDB(home string, backendType dbm.BackendType) (dbm.DB, error) {
	return OpenDB(nil, home, backendType)
}
