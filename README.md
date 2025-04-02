# TOTP CLI

A simple utility for generating TOTP auth codes (aka 2FA token or MFA tokens)

## Building

``` bash
# Grab a recent copy of golang
brew install golang

# Grab the repo
git clone git@github.com:trickyearlobe/totp-cli

# Install into your GOPATH
cd totp-cli
make install

# Cross compile for other OS'es if you need them.
# Binaries tar/gzipped into totp-cli-<YYYYMMDDHHMMSS>.tgz
make zipfile
```

Then run `totp-cli help` for usage

