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
var Blockchain []block

type block struct {
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

	genesisBlock := block{
		Index:     0,
		Timestamp: t.String(),
	}
	genesisBlock.Hash = genesisBlock.calculateHash()

	mutex.Lock()
	Blockchain = append(Blockchain, genesisBlock)
	mutex.Unlock()
}

// AddBlock creates a new Block, adds it to the Blockchain, then returns the Block.
func AddBlock(data string) block {
	var newBlock block

	t := time.Now()

	newBlock.Index = len(Blockchain)
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PreviousHash = Blockchain[newBlock.Index-1].Hash
	newBlock.Hash = newBlock.calculateHash()

	mutex.Lock()
	Blockchain = append(Blockchain, newBlock)
	mutex.Unlock()

	return newBlock
}

// calculateHash calculates and returns the hash of a Block.
func (b block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}
