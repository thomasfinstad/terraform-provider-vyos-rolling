# Development

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
