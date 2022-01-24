package transactions

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/rings"
)

type content struct {
	hash      hash.Hash
	owner     rings.Ring
	amount    hash.Hash
	origin    Origin
	external  hash.Hash
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	owner rings.Ring,
	amount hash.Hash,
	origin Origin,
	external hash.Hash,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		owner:     owner,
		amount:    amount,
		origin:    origin,
		external:  external,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Owner returns the owner
func (obj *content) Owner() rings.Ring {
	return obj.owner
}

// Amount returns the amount
func (obj *content) Amount() hash.Hash {
	return obj.amount
}

// Origin returns the origin
func (obj *content) Origin() Origin {
	return obj.origin
}

// External returns the external hash
func (obj *content) External() hash.Hash {
	return obj.external
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
