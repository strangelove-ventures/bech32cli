# bech32cli

A simple CLI to transform bech32 addresses

### Usage

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

### Example

``` bash
$ bech32 transform cosmos1ge60jkvf2wygslexprqgshxgmzd6zqludz8wyt osmo
osmo1ge60jkvf2wygslexprqgshxgmzd6zqlu9e57je
```

### Build static bins

```
$ make build-static
building bech32 amd64 static binary...
building bech32 arm64 static binary...
$ ls build
bech32-amd64  bech32-arm64
```
