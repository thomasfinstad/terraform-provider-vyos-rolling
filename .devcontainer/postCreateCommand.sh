#!/bin/bash -ex

###
# Apt packages
export DEBIAN_FRONTEND=noninteractive
apt-get update

# Get some completion up in here
apt-get install -y bash-completion

# Xml formatter
apt-get install -y libxml2-utils

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
  filesystem_mirror {
    path    = "${PWD}/dist/"
    include = ["local/providers/*"]
  }

  direct {
    exclude = ["local/providers/*"]
  }
}
EOL

# Default Terraform version
tfenv use "1.5.7"
