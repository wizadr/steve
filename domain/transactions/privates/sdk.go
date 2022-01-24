package privates

import (
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// Privates represents private transactions
type Privates interface {
	Hash() hash.Hash
	All() []Private
}

// Private represents the private part of a transfer
type Private interface {
	Hash() hash.Hash
	Owner() signature.PrivateKey
	Secret() secrets.Secret
	HasOrigin() bool
	Origin() Private
	HasSides() bool
	Sides() Privates
}
