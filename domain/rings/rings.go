package rings

import "github.com/steve-care-software/digital-diamonds/domain/hash"

type rings struct {
	hash hash.Hash
	list []Ring
}

func createRings(
	hash hash.Hash,
	list []Ring,
) Rings {
	out := rings{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *rings) Hash() hash.Hash {
	return obj.hash
}

// List returns the list of rings
func (obj *rings) List() []Ring {
	return obj.list
}
