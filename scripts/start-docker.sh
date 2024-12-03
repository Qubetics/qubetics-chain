#!/bin/bash

KEY="dev0"
CHAINID="qubetics_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t qubetics-datadir.XXXXX)

echo "create and add new keys"
./qubeticsd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Qubetics with moniker=$MONIKER and chain-id=$CHAINID"
./qubeticsd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./qubeticsd add-genesis-account \
"$(./qubeticsd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000tics,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./qubeticsd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./qubeticsd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./qubeticsd validate-genesis --home $DATA_DIR

echo "starting qubetics node $i in background ..."
./qubeticsd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started qubetics node"
tail -f /dev/null