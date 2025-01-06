import pytest
from web3 import Web3

from .ibc_utils import ATOM_IBC_DENOM, assert_ready, get_balance, prepare_network
from .network import CosmosChain, Qubetics
from .utils import (
    ADDRS,
    KEYS,
    WQUBETICS_ADDRESS,
    erc20_balance,
    erc20_transfer,
    eth_to_bech32,
    wait_for_ack,
    wait_for_fn,
)

# uatom from cosmoshub-2 -> cosmoshub-1 IBC representation on the Qubetics chain.
ATOM_2_IBC_DENOM_MULTI_HOP = (
    "ibc/D219F3A490310B65BDC312B5A644B0D56FFF1789D894B902A49FBF9D2F560B32"
)
# uatom from cosmoshub-2 -> cosmoshub-1 IBC representation
ATOM_1_IBC_DENOM_ATOM_2 = (
    "ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9"
)
# The ERC20 address of ATOM on Qubetics
ATOM_1_ERC20_ADDRESS = Web3.toChecksumAddress(
    "0xf36e4C1F926001CEaDa9cA97ea622B25f41e5eB2"
)


@pytest.fixture(scope="module", params=["qubetics"])
def ibc(request, tmp_path_factory):
    """Prepare the network"""
    name = "str-v2"
    qubetics_build = request.param
    path = tmp_path_factory.mktemp(name)
    # specify the custom_scenario
    network = prepare_network(path, name, [qubetics_build, "cosmoshub-1", "cosmoshub-2"])
    yield from network


def test_str_v2_single_hop(ibc):
    """
    Test Single Token Representation v2 with single hop Coin.
    It should automatically create an ERC20 precompiled contract.
    And register a token pair.
    """
    assert_ready(ibc)

    qubetics: Qubetics = ibc.chains["qubetics"]
    gaia: CosmosChain = ibc.chains["cosmoshub-1"]

    qubetics_cli = qubetics.cosmos_cli()
    qubetics_addr = ADDRS["signer2"]
    bech_dst = eth_to_bech32(qubetics_addr)

    gaia_cli = gaia.cosmos_cli()
    gaia_addr = gaia_cli.address("signer2")

    old_dst_balance = get_balance(qubetics, bech_dst, ATOM_IBC_DENOM)
    rsp = gaia_cli.ibc_transfer(
        gaia_addr, bech_dst, "5000uatom", "channel-0", 1, fees="10000uatom"
    )
    assert rsp["code"] == 0

    wait_for_ack(qubetics_cli, "Qubetics")

    w3 = qubetics.w3
    active_dynamic_precompiles = qubetics_cli.erc20_params()["params"][
        "dynamic_precompiles"
    ]
    new_dest_balance = get_balance(qubetics, bech_dst, ATOM_IBC_DENOM)
    erc_dest_balance = erc20_balance(w3, ATOM_1_ERC20_ADDRESS, qubetics_addr)

    assert len(active_dynamic_precompiles) == 1
    assert old_dst_balance + 5000 == new_dest_balance
    assert old_dst_balance + 5000 == erc_dest_balance


def test_str_v2_multi_hop(ibc):
    """
    Test Single Token Representation v2 with multi hop Coin.
    It should NOT create an ERC20 precompiled contract and token pair.
    """
    assert_ready(ibc)

    qubetics: Qubetics = ibc.chains["qubetics"]
    gaia: CosmosChain = ibc.chains["cosmoshub-1"]
    gaia2: CosmosChain = ibc.chains["cosmoshub-2"]

    qubetics_cli = qubetics.cosmos_cli()
    qubetics_addr = ADDRS["signer2"]
    bech_dst = eth_to_bech32(qubetics_addr)

    # The starting balance of the destination address
    qubetics_old_balance = get_balance(qubetics, bech_dst, ATOM_2_IBC_DENOM_MULTI_HOP)

    # Cosmos hub 1
    gaia_cli = gaia.cosmos_cli()
    gaia_addr = gaia_cli.address("signer2")
    gaia1_old_balance = get_balance(gaia, gaia_addr, ATOM_1_IBC_DENOM_ATOM_2)

    # Cosmos hub 2
    gaia2_cli = gaia2.cosmos_cli()

    rsp = gaia2_cli.ibc_transfer(
        gaia_addr, gaia_addr, "50000uatom", "channel-1", 1, fees="10000uatom"
    )
    assert rsp["code"] == 0

    new_dst_balance = 0

    def check_balance_change():
        nonlocal new_dst_balance
        new_dst_balance = get_balance(gaia, gaia_addr, ATOM_1_IBC_DENOM_ATOM_2)
        return gaia1_old_balance != new_dst_balance

    wait_for_fn("balance change", check_balance_change)

    new_gaia1_balance = get_balance(gaia, gaia_addr, ATOM_1_IBC_DENOM_ATOM_2)
    assert gaia1_old_balance + 50000 == new_gaia1_balance

    rsp = gaia_cli.ibc_transfer(
        gaia_addr,
        bech_dst,
        f"50000{ATOM_1_IBC_DENOM_ATOM_2}",
        "channel-0",
        1,
        fees="10000uatom",
    )
    assert rsp["code"] == 0

    wait_for_ack(qubetics_cli, "Qubetics")

    qubetics_balance = get_balance(qubetics, bech_dst, ATOM_2_IBC_DENOM_MULTI_HOP)
    dynamic_precompiles = qubetics_cli.erc20_params()["params"]["dynamic_precompiles"]
    token_pairs = qubetics_cli.get_token_pairs()

    # Here it's only one from the previous one we've registered in the first test
    assert qubetics_old_balance + 50000 == qubetics_balance
    assert len(dynamic_precompiles) == 1
    assert dynamic_precompiles[0] == ATOM_1_ERC20_ADDRESS
    assert len(token_pairs) == 2


def test_wqubetics_precompile_transfer(ibc):
    """
    Test the ERC20 transfer from one signer to another using the now
    registered ERC20 precompiled contract for WQUBETICS.
    """
    assert_ready(ibc)

    qubetics: Qubetics = ibc.chains["qubetics"]
    signer1 = ADDRS["signer1"]
    signer2 = ADDRS["signer2"]
    bech_dst = eth_to_bech32(signer2)
    src_denom = "tics"

    w3 = qubetics.w3
    qubetics_balance = get_balance(qubetics, bech_dst, src_denom)
    signer2_balance = erc20_balance(w3, WQUBETICS_ADDRESS, signer2)

    assert qubetics_balance == signer2_balance

    receipt = erc20_transfer(
        w3, WQUBETICS_ADDRESS, signer1, signer2, 1000000, KEYS["signer1"]
    )
    assert receipt.status == 1

    signer_2_balance_after = erc20_balance(w3, WQUBETICS_ADDRESS, signer2)
    assert signer_2_balance_after == signer2_balance + 1000000

    qubetics_balance_after = get_balance(qubetics, bech_dst, src_denom)
    assert qubetics_balance_after == qubetics_balance + 1000000
