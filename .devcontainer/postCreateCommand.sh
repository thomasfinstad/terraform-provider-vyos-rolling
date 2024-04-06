#!/bin/bash -ex

###
# Apt packages
APT="sudo apt-get"
export DEBIAN_FRONTEND=noninteractive
$APT update

# Get some completion up in here
$APT install -y bash-completion

# Xml formatter
$APT install -y libxml2-utils

# Github CLI
$APT install -y gh

###
# Pre-commit / git tools

# Used for commit message formatting
pipx install argcomplete
pipx install Commitizen
register-python-argcomplete cz >> ~/.bashrc

# Used to improve commits before they are committed
pipx install pre-commit
pre-commit install --hook-type pre-commit --hook-type commit-msg

# Tooling
go install mvdan.cc/gofumpt@v0.6.0
go install golang.org/x/tools/cmd/goimports@v0.19.0
go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.18.0

# Local terraform provider
# Ref: https://developer.hashicorp.com/terraform/cli/config/config-file#provider-installation
cat > "$HOME/.terraformrc" <<EOL
# Ref: https://servian.dev/terraform-local-providers-and-registry-mirror-configuration-b963117dfffa
provider_installation {
  filesystem_mirror {
    path = "${PWD}/dist"
    include = ["providers.localhost/dev/*"]
  }
  direct {
    exclude = ["providers.localhost/dev/*"]
  }
}
EOL

# Default Terraform version, before the infamous license change
tfenv use "1.5.7"
