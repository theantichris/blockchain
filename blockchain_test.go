package blockchain

import "testing"

func TestNew(t *testing.T) {
	blockchain := New()

	if len(blockchain.Blocks) == 0 {
		t.Fatal("genesis block was not generated")
	}

	got := blockchain.Blocks[0]

	if got.Index != 0 {
		t.Errorf("got Index %d want %d", got.Index, 0)
	}

	if got.Timestamp == "" {
		t.Error("got empty Timestamp")
	}
}
