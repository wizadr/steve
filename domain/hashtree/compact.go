package hashtree

import (
	"github.com/steve-care-software/steve/domain/hash"
)

type compact struct {
	hash   hash.Hash
	leaves Leaves
}

func createCompact(head hash.Hash, leaves Leaves) Compact {
	out := compact{
		hash:   head,
		leaves: leaves,
	}

	return &out
}

// Head returns the head hash
func (obj *compact) Head() hash.Hash {
	return obj.hash
}

// Leaves returns the leaves
func (obj *compact) Leaves() Leaves {
	return obj.leaves
}

// Length returns the length of the compact hashtree
func (obj *compact) Length() int {
	return len(obj.leaves.Leaves())
}
