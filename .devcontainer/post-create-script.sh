#!/bin/bash -ex

echo "Current user: $(id)"

# Trust repo
git config --global --add safe.directory /workspaces/terraform-provider-vyos-rolling

# Used for commit message formatting
register-python-argcomplete cz >> ~/.bashrc
register-python-argcomplete pre-commit >> ~/.bashrc

# Used to improve commits before they are committed
pre-commit install --hook-type pre-commit --hook-type commit-msg
