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
var blockchain Blockchain

// Blockchain stores information on the chain as well as the blocks.
type Blockchain struct {
	blocks    []Block
	consensus Consensus
}

// Consensus returns the name of the consensus mechanism the blockchain is using.
func (bc Blockchain) Consensus() string {
	return bc.consensus.String()
}

// Block respresents each block in the blockchain.
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

// New initializes the blockchain with a genesis block.
func New(consensus Consensus) {
	if blockchain.blocks != nil {
		return
	}

	blockchain.consensus = consensus

	t := time.Now()

	genesisBlock := Block{
		Index:     0,
		Timestamp: t.String(),
	}
	genesisBlock.Hash = genesisBlock.calculateHash()

	mutex.Lock()
	blockchain.blocks = append(blockchain.blocks, genesisBlock)
	mutex.Unlock()
}

// AddBlock creates a new block, adds it to the blockchain, and returns the new block.
func AddBlock(data string) Block {
	t := time.Now()
	index := len(blockchain.blocks)

	newBlock := Block{
		Index:        index,
		Timestamp:    t.String(),
		Data:         data,
		PreviousHash: blockchain.blocks[index-1].Hash,
	}
	newBlock.Hash = newBlock.calculateHash()

	newBlocks := append(blockchain.blocks, newBlock)

	mutex.Lock()
	if len(newBlocks) > len(blockchain.blocks) {
		blockchain.blocks = newBlocks
	}
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
