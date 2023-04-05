#!/bin/bash -ex

###
# Apt packages
export DEBIAN_FRONTEND=noninteractive
sudo apt-get update

# Get some completion up in here
sudo apt-get install -y bash-completion

###
# Pre-commit / git tools

# Used for commit message formatting
pipx install argcomplete
pipx install Commitizen
register-python-argcomplete cz >> ~/.bashrc

# Used to improve commits before they are commited
pipx install pre-commit
pre-commit install --hook-type pre-commit --hook-type commit-msg --hook-type pre-push

# Tooling
go install mvdan.cc/gofumpt@latest
