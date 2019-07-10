import os
import pytest

import testinfra.utils.ansible_runner

testinfra_hosts = testinfra.utils.ansible_runner.AnsibleRunner(
    os.environ['MOLECULE_INVENTORY_FILE']).get_hosts('all')


# GOOD: Each package is tested
# $ py.test -v test.py
# [...]
# test.py::test_package[local-nginx-1.6] PASSED
# test.py::test_package[local-python-2.7] PASSED
# [...]


@pytest.mark.parametrize("name,version", [
    ("tree", "1.8.0"),
    ("htop", "2.2.0"),
    ("net-tools", "1.60"),
    ("nano", "2.0.6"),
    ("vim", "8.1"),
    ("wget", "1.20.3"),
])
def test_packages(host, name, version):
    pkg = host.package(name)
    ver = host.package(version)
    assert pkg.is_installed
    assert ver.name
