# bech32cli

A simple CLI to transform bech32 addresses

## Usage

### Transform address from one bech32 prefix to another (same coin type)

``` bash
$ bech32 transform -h
Transforms bech32 string to new prefix

Usage:
  bech32 transform [bech32Address] [newBech32Prefix] [flags]

Aliases:
  transform, t

Examples:
$ bech32 transform cosmos1ge60jkvf2wygslexprqgshxgmzd6zqludz8wyt osmo
$ bech32 t cosmos1ge60jkvf2wygslexprqgshxgmzd6zqludz8wyt osmo
```

### Derive a validator address from either a hex address or a base64 pubkey

```bash
$ bech32 valcons -h
validator consensus address transformation

Usage:
  bech32 valcons [flags]

Aliases:
  valcons, v

Examples:
$ bech32 valcons osmo --pubkey wC+QT4cw8WWOwRZhL/XZ8XusXSH7Q3kvhEnFFPagXis=
$ bech32 v osmo --pubkey wC+QT4cw8WWOwRZhL/XZ8XusXSH7Q3kvhEnFFPagXis=
$ bech32 v osmo --address 023DCF3F6AEA4E0098ABBA2AF23F3D65AC324851
$ bech32 v osmo --address 023DCF3F6AEA4E0098ABBA2AF23F3D65AC324851

Flags:
      --address string   validator hex address to transform
      --pubkey string    validator base64 pubkey to transform
```


### Build static bins

```
$ make build-static
building bech32 amd64 static binary...
building bech32 arm64 static binary...
$ ls build
bech32-amd64  bech32-arm64
```
