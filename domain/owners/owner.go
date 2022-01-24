package owners

import (
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/rings"
)

type owner struct {
	hash  hash.Hash
	rings []rings.Ring
}

func createOwner(
	hash hash.Hash,
	rings []rings.Ring,
) Owner {
	out := owner{
		hash:  hash,
		rings: rings,
	}

	return &out
}

// Hash returns the hash
func (obj *owner) Hash() hash.Hash {
	return obj.hash
}

// Rings returns the rings
func (obj *owner) Rings() []rings.Ring {
	return obj.rings
}
