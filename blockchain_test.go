package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	New(NoConsensus)

	if len(blockchain) == 0 {
		t.Fatal("genesis block was not generated")
	}

	got := blockchain[0]

	if got.Index != 0 {
		t.Errorf("genesis block has an incorrect Index %d want %d", got.Index, 0)
	}

	if got.Timestamp == "" {
		t.Error("genesis block has an empty Timestamp")
	}

	if consensus != NoConsensus {
		t.Errorf("the consensus was not set correctly: got %q want %q", consensus.String(), NoConsensus.String())
	}
}

func TestAddBlock(t *testing.T) {
	New(NoConsensus)

	data := `{"key": "value"}`
	newBlock := AddBlock(data)

	if len(blockchain) != 2 {
		t.Fatal("the block was not added to the chain")
	}

	if newBlock.Index != len(blockchain)-1 {
		t.Errorf("the block was not indexed correctly, got %d want %d", newBlock.Index, len(blockchain)-1)
	}

	if newBlock != blockchain[len(blockchain)-1] {
		t.Errorf("the block does not match last block on chain: got %v want %v", newBlock, blockchain[len(blockchain)-1])
	}

	if newBlock.Timestamp == "" {
		t.Error("the block does not have a timestamp")
	}

	if newBlock.Data != data {
		t.Errorf("the data was not saved to the block: got %q want %q", newBlock.Data, data)
	}

	if newBlock.PreviousHash != blockchain[newBlock.Index-1].Hash {
		t.Errorf("the previous block's hash does not match: got %q want %q", newBlock.PreviousHash, blockchain[newBlock.Index-1].Hash)
	}

	if newBlock.Hash != newBlock.calculateHash() {
		t.Errorf("the block did not get the correct hash: got %q want %q", newBlock.Hash, newBlock.calculateHash())
	}
}

func TestCalculateHash(t *testing.T) {
	newBlock := block{
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
