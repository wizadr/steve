package privates

import (
	"github.com/steve-care-software/steve/domain/cryptography/keys/signature"
	"github.com/steve-care-software/steve/domain/genesis"
	"github.com/steve-care-software/steve/domain/hash"
)

type private struct {
	hash hash.Hash
	unit genesis.Unit
	pk   signature.PrivateKey
}

func createPrivate(
	hash hash.Hash,
	unit genesis.Unit,
	pk signature.PrivateKey,
) Private {
	out := private{
		hash: hash,
		unit: unit,
		pk:   pk,
	}

	return &out
}

// Hash returns the hash
func (obj *private) Hash() hash.Hash {
	return obj.hash
}

// Unit returns the genesis unit
func (obj *private) Unit() genesis.Unit {
	return obj.unit
}

// PrivateKey returns the private key
func (obj *private) PrivateKey() signature.PrivateKey {
	return obj.pk
}
