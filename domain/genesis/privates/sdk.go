package privates

import (
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/genesis"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// Builder represents the privates builder
type Builder interface {
	Create() Builder
	WithList(list []Private) Builder
	Now() (Privates, error)
}

// Privates represents private units
type Privates interface {
	Hash() hash.Hash
	All() []Private
}

// PrivateKeyBuilder represents a private key builder
type PrivateKeyBuilder interface {
	Create() PrivateKeyBuilder
	WithUnit(unit genesis.Unit) PrivateKeyBuilder
	WithPrivateKey(pk signature.PrivateKey) PrivateKeyBuilder
	Now() (Private, error)
}

// Private represents a private unit
type Private interface {
	Hash() hash.Hash
	Unit() genesis.Unit
	PrivateKey() signature.PrivateKey
}
