# blockchain

[![GoDoc](https://godoc.org/github.com/theantichris/blockchain?status.svg)](https://godoc.org/github.com/theantichris/blockchain)
[![Actions Status](https://github.com/theantichris/blockchain/workflows/Go/badge.svg)](https://github.com/theantichris/blockchain/actions)

A library for creating and running blockchains.

## Install

```bash
go get -u github.com/theantichris/blockchain
```

## Example

```go
package main

import "github.com/theantichris/blockchain"

func main() {
  blockchain.New(NoConsensus)

  // your code to accept, validate, and convert data to a string
  data := "your data"

  blockchain.Add(data)
}
```

## Consensus mechanisms

When instantiating the blockchain you will need to specify the consensus mechanism you want to use.

* NoConsensus
* ProofOfWork

## License

MIT licensed. See LICENSE file for details.
