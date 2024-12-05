import json

from .utils import (
    ADDRS,
    CONTRACTS,
    KEYS,
    build_deploy_contract_tx,
    deploy_contract,
    send_transaction,
    w3_wait_for_new_blocks,
)


def test_gas_eth_tx(geth, qubetics_cluster):
    tx_value = 10

    # send a transaction with geth
    geth_gas_price = geth.w3.eth.gas_price
    tx = {"to": ADDRS["community"], "value": tx_value, "gasPrice": geth_gas_price}
    geth_receipt = send_transaction(geth.w3, tx, KEYS["validator"])

    # send an equivalent transaction with qubetics
    qubetics_gas_price = qubetics_cluster.w3.eth.gas_price
    tx = {"to": ADDRS["community"], "value": tx_value, "gasPrice": qubetics_gas_price}
    qubetics_receipt = send_transaction(qubetics_cluster.w3, tx, KEYS["validator"])

    assert geth_receipt.gasUsed == qubetics_receipt.gasUsed


def test_gas_deployment(geth, qubetics_cluster):
    # deploy an identical contract on geth and qubetics
    # ensure that the gasUsed is equivalent
    info = json.loads(CONTRACTS["TestERC20A"].read_text())
    geth_tx = build_deploy_contract_tx(geth.w3, info)
    qubetics_tx = build_deploy_contract_tx(qubetics_cluster.w3, info)

    # estimate tx gas
    geth_gas_estimation = geth.w3.eth.estimate_gas(geth_tx)
    qubetics_gas_estimation = qubetics_cluster.w3.eth.estimate_gas(qubetics_tx)

    assert geth_gas_estimation == qubetics_gas_estimation

    # sign and send tx
    geth_contract_receipt = send_transaction(geth.w3, geth_tx)
    qubetics_contract_receipt = send_transaction(qubetics_cluster.w3, qubetics_tx)
    assert geth_contract_receipt.status == 1
    assert qubetics_contract_receipt.status == 1

    assert geth_contract_receipt.gasUsed == qubetics_contract_receipt.gasUsed

    # gasUsed should be same as estimation
    assert geth_contract_receipt.gasUsed == geth_gas_estimation
    assert qubetics_contract_receipt.gasUsed == qubetics_gas_estimation


def test_gas_call(geth, qubetics_cluster):
    function_input = 10

    # deploy an identical contract on geth and qubetics
    # ensure that the contract has a function which consumes non-trivial gas
    geth_contract, _ = deploy_contract(geth.w3, CONTRACTS["BurnGas"])
    qubetics_contract, _ = deploy_contract(qubetics_cluster.w3, CONTRACTS["BurnGas"])

    # call the contract and get tx receipt for geth
    geth_gas_price = geth.w3.eth.gas_price
    geth_tx = geth_contract.functions.burnGas(function_input).build_transaction(
        {"from": ADDRS["validator"], "gasPrice": geth_gas_price}
    )
    geth_gas_estimation = geth.w3.eth.estimate_gas(geth_tx)
    geth_call_receipt = send_transaction(geth.w3, geth_tx)

    # repeat the above for qubetics
    qubetics_gas_price = qubetics_cluster.w3.eth.gas_price
    qubetics_tx = qubetics_contract.functions.burnGas(function_input).build_transaction(
        {"from": ADDRS["validator"], "gasPrice": qubetics_gas_price}
    )
    qubetics_gas_estimation = qubetics_cluster.w3.eth.estimate_gas(qubetics_tx)
    qubetics_call_receipt = send_transaction(qubetics_cluster.w3, qubetics_tx)

    # ensure gas estimation is the same
    assert geth_gas_estimation == qubetics_gas_estimation

    # ensure that the gasUsed is equivalent
    assert geth_call_receipt.gasUsed == qubetics_call_receipt.gasUsed

    # ensure gasUsed == gas estimation
    assert geth_call_receipt.gasUsed == geth_gas_estimation
    assert qubetics_call_receipt.gasUsed == qubetics_gas_estimation


def test_block_gas_limit(qubetics_cluster):
    tx_value = 10

    # get the block gas limit from the latest block
    w3_wait_for_new_blocks(qubetics_cluster.w3, 5)
    block = qubetics_cluster.w3.eth.get_block("latest")
    exceeded_gas_limit = block.gasLimit + 100

    # send a transaction exceeding the block gas limit
    qubetics_gas_price = qubetics_cluster.w3.eth.gas_price
    tx = {
        "to": ADDRS["community"],
        "value": tx_value,
        "gas": exceeded_gas_limit,
        "gasPrice": qubetics_gas_price,
    }

    # expect an error due to the block gas limit
    try:
        send_transaction(qubetics_cluster.w3, tx, KEYS["validator"])
    except Exception as error:
        assert "exceeds block gas limit" in error.args[0]["message"]

    # deploy a contract on qubetics
    qubetics_contract, _ = deploy_contract(qubetics_cluster.w3, CONTRACTS["BurnGas"])

    # expect an error on contract call due to block gas limit
    try:
        qubetics_txhash = qubetics_contract.functions.burnGas(exceeded_gas_limit).transact(
            {
                "from": ADDRS["validator"],
                "gas": exceeded_gas_limit,
                "gasPrice": qubetics_gas_price,
            }
        )
        (qubetics_cluster.w3.eth.wait_for_transaction_receipt(qubetics_txhash))
    except Exception as error:
        assert "exceeds block gas limit" in error.args[0]["message"]

    return
