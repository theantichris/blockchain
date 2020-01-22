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
	block := AddBlock(data)

	if len(Blockchain) != 2 {
		t.Fatal("the block was not added to the chain")
	}

	if block.Index != len(Blockchain)-1 {
		t.Errorf("the block was not indexed correctly, got %d want %d", block.Index, len(Blockchain)-1)
	}

	if block != Blockchain[len(Blockchain)-1] {
		t.Errorf("block does not match last block on chain: got %v want %v", block, Blockchain[len(Blockchain)-1])
	}

	if block.Timestamp == "" {
		t.Error("the block did not save the timestamp")
	}

	if block.Data != data {
		t.Errorf("the data was not saved to the block: got %q want %q", block.Data, data)
	}

	if block.PreviousHash != Blockchain[block.Index-1].Hash {
		t.Errorf("the previous block's hash does not match: got %q want %q", block.PreviousHash, Blockchain[block.Index-1].Hash)
	}

	if block.Hash != block.CalculateHash() {
		t.Errorf("the block did the correct hash: got %q want %q", block.Hash, block.CalculateHash())
	}
}

func TestCalculateHash(t *testing.T) {
	block := Block{
		Index:        0,
		Timestamp:    "timestamp string",
		Data:         "data string",
		PreviousHash: "",
	}

	block.Hash = block.CalculateHash()

	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	want := hex.EncodeToString(hashed)

	if block.Hash != want {
		t.Errorf("got %q want %q", block.Hash, want)
	}
}
