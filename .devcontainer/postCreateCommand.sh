#!/bin/bash -ex

# Used for commit message formatting
register-python-argcomplete cz >> ~/.bashrc
register-python-argcomplete pre-commit >> ~/.bashrc

# Used to improve commits before they are committed
pre-commit install --hook-type pre-commit --hook-type commit-msg
