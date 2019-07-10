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
    ("git", "2.22"),
    ("python", "3.6"),
    ("tree", "1.8.0"),
])
def test_packages(host, name, version):
    pkg = host.package(name)
    ver = host.package(version)
    assert pkg.is_installed
    assert ver.name
