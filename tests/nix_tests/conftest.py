import pytest

from .network import setup_qubetics, setup_qubetics_rocksdb, setup_geth


@pytest.fixture(scope="session")
def qubetics(tmp_path_factory):
    path = tmp_path_factory.mktemp("qubetics")
    yield from setup_qubetics(path, 26650)


@pytest.fixture(scope="session")
def qubetics_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("qubetics-rocksdb")
    yield from setup_qubetics_rocksdb(path, 20650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["qubetics", "qubetics-ws"])
def qubetics_rpc_ws(request, qubetics):
    """
    run on both qubetics and qubetics websocket
    """
    provider = request.param
    if provider == "qubetics":
        yield qubetics
    elif provider == "qubetics-ws":
        qubetics_ws = qubetics.copy()
        qubetics_ws.use_websocket()
        yield qubetics_ws
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["qubetics", "qubetics-ws", "qubetics-rocksdb", "geth"])
def cluster(request, qubetics, qubetics_rocksdb, geth):
    """
    run on qubetics, qubetics websocket,
    qubetics built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "qubetics":
        yield qubetics
    elif provider == "qubetics-ws":
        qubetics_ws = qubetics.copy()
        qubetics_ws.use_websocket()
        yield qubetics_ws
    elif provider == "geth":
        yield geth
    elif provider == "qubetics-rocksdb":
        yield qubetics_rocksdb
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["qubetics", "qubetics-rocksdb"])
def qubetics_cluster(request, qubetics, qubetics_rocksdb):
    """
    run on qubetics default build &
    qubetics with rocksdb build and memIAVL + versionDB
    """
    provider = request.param
    if provider == "qubetics":
        yield qubetics
    elif provider == "qubetics-rocksdb":
        yield qubetics_rocksdb
    else:
        raise NotImplementedError
