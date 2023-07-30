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
pre-commit install --hook-type pre-commit --hook-type commit-msg

# Tooling
go install mvdan.cc/gofumpt@latest
go install golang.org/x/tools/cmd/goimports@latest

# Local terraform provider
cat > "$HOME/.terraformrc" <<EOL
provider_installation {

  dev_overrides {
      "ex.c/thomasfinstad/vyos" = "/go/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
EOL

# Default Terraform version
tfenv use "1.4.4"
