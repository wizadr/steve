package genesis

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type unitContent struct {
	hash        hash.Hash
	amount      uint
	nonce       string
	activatedOn time.Time
	createdOn   time.Time
}

func createUnitContent(
	hash hash.Hash,
	amount uint,
	nonce string,
	activatedOn time.Time,
	createdOn time.Time,
) UnitContent {
	out := unitContent{
		hash:        hash,
		amount:      amount,
		nonce:       nonce,
		activatedOn: activatedOn,
		createdOn:   createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *unitContent) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *unitContent) Amount() uint {
	return obj.amount
}

// Nonce returns the nonce
func (obj *unitContent) Nonce() string {
	return obj.nonce
}

// ActivatedOn returns the activation time
func (obj *unitContent) ActivatedOn() time.Time {
	return obj.activatedOn
}

// CreatedOn returns the creation time
func (obj *unitContent) CreatedOn() time.Time {
	return obj.createdOn
}
