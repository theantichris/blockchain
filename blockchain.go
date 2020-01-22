// Package blockchain is a library for creating and running blockchains.
package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

// Blockchain contains the blocks in the chain.
type Blockchain struct {
	Blocks []Block
}

// Block contains the data stored in the blockchain.
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

// New creates and returns a new Blockchain.
func New() Blockchain {
	blockchain := Blockchain{}

	t := time.Now()

	block := Block{
		Index:     0,
		Timestamp: t.String(),
	}

	mutex.Lock()
	blockchain.Blocks = append(blockchain.Blocks, block)
	mutex.Unlock()

	return blockchain
}

// CalculateHash calculates and returns the hash of a Block.
func (b Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}
