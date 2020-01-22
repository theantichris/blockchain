package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	blockchain := New()

	if len(blockchain) == 0 {
		t.Fatal("genesis block was not generated")
	}

	got := blockchain[0]

	if got.Index != 0 {
		t.Errorf("got Index %d want %d", got.Index, 0)
	}

	if got.Timestamp == "" {
		t.Error("got empty Timestamp")
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
