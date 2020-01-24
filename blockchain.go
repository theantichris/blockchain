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

var blockchain []Block
var consensus Consensus

// Block respresents each block in the blockchain.
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

// New initializes the blockchain with a genesis block.
func New(c Consensus) {
	if blockchain != nil {
		return
	}

	consensus = c

	t := time.Now()

	genesisBlock := Block{
		Index:     0,
		Timestamp: t.String(),
	}
	genesisBlock.Hash = genesisBlock.calculateHash()

	mutex.Lock()
	blockchain = append(blockchain, genesisBlock)
	mutex.Unlock()
}

// AddBlock creates a new block, adds it to the blockchain, and returns the new block.
func AddBlock(data string) Block {
	t := time.Now()
	index := len(blockchain)

	newBlock := Block{
		Index:        index,
		Timestamp:    t.String(),
		Data:         data,
		PreviousHash: blockchain[index-1].Hash,
	}
	newBlock.Hash = newBlock.calculateHash()

	mutex.Lock()
	blockchain = append(blockchain, newBlock)
	mutex.Unlock()

	return newBlock
}

func (b Block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}
