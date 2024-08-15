# Development

<!--TOC-->

- [Development](#development)
  - [Intro](#intro)
  - [Workflow](#workflow)
  - [Troubleshooting](#troubleshooting)
    - [Dev Container](#dev-container)
      - [Linux](#linux)
        - [Podman](#podman)
      - [WSL2](#wsl2)

<!--TOC-->

## Intro

This repo is highly reliant on using devcontainers to keep the development environment the same over time and for all contributors.

All development "actions" are done via the provided Makefile

## Workflow

Due to the somewhat confusing nature of the workflow this list can be used as a reference:

1.  

## Troubleshooting

### Dev Container

General References:

* [Sharing git credentials](https://code.visualstudio.com/remote/advancedcontainers/sharing-git-credentials)

General tips:

* Unless your user has `UID=1000`, and you are using rootless podman you will need to change the docker socket mount in `.devcontainer/devcontainer.json`

#### Linux

##### Podman

> Tip: podman-desktop is a nice and handy tool for a gui overview, for a tui client check out lazydocker

If you are using podman, make sure the docker socket compatibility is enabled: `sudo systemctl enable podman.socket`,
alternatively you can configure the devcontainer to mount the user socket, usually found at `/run/user/$(id -u)/podman/podman.sock`.

You can verify the socket access with curl, if it work it will show *"OK"*:
`curl -H "Content-Type: application/json" --unix-socket /var/run/docker.sock http://localhost/_ping`, if that does not work but the user socket works:
`curl -H "Content-Type: application/json" --unix-socket /run/user/$(id -u)/podman/podman.sock http://localhost/_ping`
you can try to communicate with the system socket using sudo.
If the system socket works as root, but not your user, you might need to add your user to the `podman` group (and log out and in again).

#### WSL2

Some workaround / issues must be handled to get dev containers to work well in WSL2.

1. Make sure you have a working SSH-Agent by updating the one in Windows
    * `Remove-WindowsCapability -Online -Name OpenSSH.Client~~~~0.0.1.0`
    * `winget install openssh-beta`
    * [Reference](https://superuser.com/a/1722263)
2. Make sure GnuPG works:
    * In windows:
        * `winget install GnuPG.Gpg4win`
    * In WSL2:
        * `sudo apt-get install gpg gnupg gpg-agent`
        * `vim ~/.gnupg/gpg-agent.conf`
            ```
            default-cache-ttl 34560000
            max-cache-ttl 34560000
            pinentry-program "/mnt/c/Program Files (x86)/GnuPG/bin/pinentry-basic.exe"
            ```
        * `gpgconf --kill gpg-agent`
    * [Reference](https://www.39digits.com/signed-git-commits-on-wsl2-using-visual-studio-code)
