from pathlib import Path

import pytest

from .network import create_snapshots_dir, setup_custom_qubetics
from .utils import memiavl_config, wait_for_block


@pytest.fixture(scope="module")
def custom_qubetics(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp")
    yield from setup_custom_qubetics(
        path,
        26260,
        Path(__file__).parent / "configs/discard-abci-resp.jsonnet",
    )


@pytest.fixture(scope="module")
def custom_qubetics_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp-rocksdb")
    yield from setup_custom_qubetics(
        path,
        26810,
        memiavl_config(path, "discard-abci-resp"),
        post_init=create_snapshots_dir,
        chain_binary="qubeticsd-rocksdb",
    )


@pytest.fixture(scope="module", params=["qubetics", "qubetics-rocksdb"])
def qubetics_cluster(request, custom_qubetics, custom_qubetics_rocksdb):
    """
    run on qubetics and
    qubetics built with rocksdb (memIAVL + versionDB)
    """
    provider = request.param
    if provider == "qubetics":
        yield custom_qubetics
    elif provider == "qubetics-rocksdb":
        yield custom_qubetics_rocksdb
    else:
        raise NotImplementedError


def test_gas_eth_tx(qubetics_cluster):
    """
    When node does not persist ABCI responses
    eth_gasPrice should return an error instead of crashing
    """
    wait_for_block(qubetics_cluster.cosmos_cli(), 3)
    try:
        qubetics_cluster.w3.eth.gas_price
        raise Exception("This query should have failed")
    except Exception as error:
        assert "node is not persisting abci responses" in error.args[0]["message"]
