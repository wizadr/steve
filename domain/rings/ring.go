package rings

import "github.com/steve-care-software/digital-diamonds/domain/hash"

type ring struct {
	hash hash.Hash
	list []hash.Hash
}

func createRing(
	hash hash.Hash,
	list []hash.Hash,
) Ring {
	out := ring{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *ring) Hash() hash.Hash {
	return obj.hash
}

// List returns the list of hashes
func (obj *ring) List() []hash.Hash {
	return obj.list
}
