package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name string
		c    Consensus
	}{
		{name: "creates new no consensus blockchain", c: NoConsensus},
		{name: "creates new proof-of-work blockchain", c: ProofOfWork},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			blockchain = Blockchain{}
			New(tt.c)

			if blockchain.Length() == 0 {
				t.Fatal("genesis block was not generated")
			}

			got := blockchain.Block(0)

			if got.Index != 0 {
				t.Errorf("genesis block has an incorrect Index %d want %d", got.Index, 0)
			}

			if got.Timestamp == "" {
				t.Error("genesis block has an empty Timestamp")
			}

			if blockchain.Consensus() != tt.c.String() {
				t.Errorf("the consensus was not set correctly: got %q want %q", blockchain.Consensus(), tt.c.String())
			}
		})
	}

	t.Run("does not overwrite existing blockchain", func(t *testing.T) {
		blockchain = Blockchain{}
		New(NoConsensus)

		block1 := blockchain.Block(0)

		New(NoConsensus)
		block2 := blockchain.Block(0)

		if block1 != block2 {
			t.Errorf("a new blockchain overwrote the old blockchain, block1: %v block2: %v", block1, block2)
		}
	})
}

func TestAddBlock(t *testing.T) {
	blockchain = Blockchain{}
	New(NoConsensus)

	data := `{"key": "value"}`
	newBlock := AddBlock(data)
	length := blockchain.Length()

	if length != 2 {
		t.Fatal("the block was not added to the chain")
	}

	if newBlock.Index != length-1 {
		t.Errorf("the block was not indexed correctly, got %d want %d", newBlock.Index, length-1)
	}

	if newBlock != blockchain.Block(length-1) {
		t.Errorf("the block does not match last block on chain: got %v want %v", newBlock, blockchain.Block(length-1))
	}

	if newBlock.Timestamp == "" {
		t.Error("the block does not have a timestamp")
	}

	if newBlock.Data != data {
		t.Errorf("the data was not saved to the block: got %q want %q", newBlock.Data, data)
	}

	if newBlock.PreviousHash != blockchain.Block(newBlock.Index-1).Hash {
		t.Errorf("the previous block's hash does not match: got %q want %q", newBlock.PreviousHash, blockchain.Block(newBlock.Index-1).Hash)
	}

	if newBlock.Hash != newBlock.calculateHash() {
		t.Errorf("the block did not get the correct hash: got %q want %q", newBlock.Hash, newBlock.calculateHash())
	}
}

func TestCalculateHash(t *testing.T) {
	newBlock := Block{
		Index:        0,
		Timestamp:    "timestamp string",
		Data:         "data string",
		PreviousHash: "",
	}

	newBlock.Hash = newBlock.calculateHash()

	record := strconv.Itoa(newBlock.Index) + newBlock.Timestamp + newBlock.Data + newBlock.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	want := hex.EncodeToString(hashed)

	if newBlock.Hash != want {
		t.Errorf("got %q want %q", newBlock.Hash, want)
	}
}
