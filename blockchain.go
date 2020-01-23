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

// Blockchain contains the Blocks in the chain.
var Blockchain []Block

// Block contains the data stored in the Blockchain.
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

// New creates and returns a new Blockchain.
func New() {
	Blockchain = nil

	t := time.Now()

	block := Block{
		Index:     0,
		Timestamp: t.String(),
	}
	block.Hash = block.calculateHash()

	mutex.Lock()
	Blockchain = append(Blockchain, block)
	mutex.Unlock()
}

// AddBlock creates a new Block, adds it to the Blockchain, then returns the Block.
func AddBlock(data string) Block {
	var block Block

	t := time.Now()

	block.Index = len(Blockchain)
	block.Timestamp = t.String()
	block.Data = data
	block.PreviousHash = Blockchain[block.Index-1].Hash
	block.Hash = block.calculateHash()

	mutex.Lock()
	Blockchain = append(Blockchain, block)
	mutex.Unlock()

	return block
}

// calculateHash calculates and returns the hash of a Block.
func (b Block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}
