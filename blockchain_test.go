package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	New()

	if len(Blockchain) == 0 {
		t.Fatal("genesis block was not generated")
	}

	got := Blockchain[0]

	if got.Index != 0 {
		t.Errorf("got Index %d want %d", got.Index, 0)
	}

	if got.Timestamp == "" {
		t.Error("got empty Timestamp")
	}
}

func TestAddBlock(t *testing.T) {
	New()

	data := `{"key": "value"}`
	newBlock := AddBlock(data)

	if len(Blockchain) != 2 {
		t.Fatal("the block was not added to the chain")
	}

	if newBlock.Index != len(Blockchain)-1 {
		t.Errorf("the block was not indexed correctly, got %d want %d", newBlock.Index, len(Blockchain)-1)
	}

	if newBlock != Blockchain[len(Blockchain)-1] {
		t.Errorf("block does not match last block on chain: got %v want %v", newBlock, Blockchain[len(Blockchain)-1])
	}

	if newBlock.Timestamp == "" {
		t.Error("the block did not save the timestamp")
	}

	if newBlock.Data != data {
		t.Errorf("the data was not saved to the block: got %q want %q", newBlock.Data, data)
	}

	if newBlock.PreviousHash != Blockchain[newBlock.Index-1].Hash {
		t.Errorf("the previous block's hash does not match: got %q want %q", newBlock.PreviousHash, Blockchain[newBlock.Index-1].Hash)
	}

	if newBlock.Hash != newBlock.calculateHash() {
		t.Errorf("the block did the correct hash: got %q want %q", newBlock.Hash, newBlock.calculateHash())
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
