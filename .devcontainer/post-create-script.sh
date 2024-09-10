#!/bin/bash -ex

echo "Current user: $(id)"

# Trust repo
git config --global --add safe.directory $PWD

# Used for commit message formatting
register-python-argcomplete cz >> ~/.bashrc
register-python-argcomplete pre-commit >> ~/.bashrc

# Used to improve commits before they are committed
pre-commit install --hook-type pre-commit --hook-type commit-msg

# tenv setup
tenv completion bash >> ~/.bashrc
cd examples/provider
tenv tofu install
cd ../..
tofu -install-autocomplete
