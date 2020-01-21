// Package blockchain is a library for creating and running blockchains.
package blockchain

import "time"

// Blockchain contains the blocks in the chain.
type Blockchain struct {
	Blocks []Block
}

// Block contains the data stored in the blockchain.
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	Hash         string
	PreviousHash string
}

// New creates and returns a new Blockchain.
func New() Blockchain {
	blockchain := Blockchain{}

	t := time.Now()

	block := Block{
		Index:     0,
		Timestamp: t.String(),
	}

	blockchain.Blocks = append(blockchain.Blocks, block)

	return blockchain
}
