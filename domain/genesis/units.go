package genesis

import "github.com/steve-care-software/digital-diamonds/domain/hash"

type units struct {
	hash hash.Hash
	list []Unit
}

func createUnits(
	hash hash.Hash,
	list []Unit,
) Units {
	out := units{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *units) Hash() hash.Hash {
	return obj.hash
}

// All returns the units
func (obj *units) All() []Unit {
	return obj.list
}
