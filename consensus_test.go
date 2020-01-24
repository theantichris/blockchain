package blockchain

import "testing"

func TestString(t *testing.T) {
	cases := []struct {
		name string
		c    Consensus
	}{
		{name: "No Consensus", c: NoConsensus},
		{name: "Proof-Of-Work", c: ProofOfWork},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.String()

			if got != tt.name {
				t.Errorf("got %q want %q", got, tt.name)
			}
		})
	}
}
