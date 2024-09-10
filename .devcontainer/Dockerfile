FROM mcr.microsoft.com/devcontainers/base:ubuntu22.04

###
# Custom shell environment
COPY dockerfile-profile.d.sh /etc/profile.d/dockerfile-profile.d.sh

###
# APT Packages
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update
RUN apt-get dist-upgrade -y

RUN apt-get install -y software-properties-common

RUN add-apt-repository ppa:longsleep/golang-backports

RUN apt-get install -y python3
RUN apt-get install -y jq
RUN apt-get install -y make
RUN apt-get install -y openjdk-17-jre
RUN apt-get install -y golang-go
RUN apt-get install -y git
RUN apt-get install -y docker.io
RUN apt-get install -y bash-completion
RUN apt-get install -y libxml2-utils
RUN apt-get install -y gh
RUN apt-get install -y pipx

ENV DEBIAN_FRONTEND=""

###
# TENV
RUN curl -O -L "https://github.com/tofuutils/tenv/releases/latest/download/tenv_v3.2.2_amd64.deb"
RUN sudo dpkg -i "tenv_v3.2.2_amd64.deb"

###
# Dist directory is hard coded due to .terraformrc file
RUN mkdir /dist
RUN chown -R vscode /dist

###
# Custom user settings
RUN sudo -u vscode bash -c "pipx install argcomplete==3.3.0"
RUN sudo -u vscode bash -c "pipx install Commitizen==3.25.0"

RUN sudo -u vscode bash -c "pipx install pre-commit==3.7.0"
# This is very hacky, but we just want pre-commit available for all during github action runs
RUN ln -s /home/vscode/.local/bin/pre-commit /usr/bin/pre-commit

# Install go packages / tools
RUN sudo -u vscode bash -c "go install -v github.com/go-delve/delve/cmd/dlv@latest"
RUN sudo -u vscode bash -c "go install mvdan.cc/gofumpt@v0.6.0"
RUN sudo -u vscode bash -c "go install golang.org/x/tools/cmd/goimports@v0.19.0"
RUN sudo -u vscode bash -c "go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.19.4"
RUN sudo -u vscode bash -c "go install github.com/jesseduffield/lazygit@latest"

# Add our custom terraform rc config
COPY --chown=vscode:vscode terraformrc /home/vscode/.terraformrc
