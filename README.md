# blockchain

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
