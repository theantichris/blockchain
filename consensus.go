package blockchain

// Consensus is an enum used to set what type of consensus mechanism the blockchain will use.
type Consensus int

const (
	// NoConsensus sets the blockchain to no use a consensus mechanism.
	NoConsensus Consensus = 0
	// ProofOfWork sets the blockchain to use a proof-of-work mechanism.
	ProofOfWork Consensus = 1
)

func (c Consensus) String() string {
	names := [...]string{
		"No Consensus",
		"Proof-Of-Work", // TODO: implement proof-of-work
	}

	return names[c]
}
