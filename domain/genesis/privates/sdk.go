package privates

import (
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/genesis"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewPrivateBuilder creates a new private builder
func NewPrivateBuilder() PrivateBuilder {
	hashAdapter := hash.NewAdapter()
	return createPrivateBuilder(hashAdapter)
}

// Builder represents the privates builder
type Builder interface {
	Create() Builder
	WithList(list []Private) Builder
	Now() (Privates, error)
}

// Privates represents private units
type Privates interface {
	Hash() hash.Hash
	List() []Private
}

// PrivateBuilder represents a private builder
type PrivateBuilder interface {
	Create() PrivateBuilder
	WithUnit(unit genesis.Unit) PrivateBuilder
	WithPrivateKey(pk signature.PrivateKey) PrivateBuilder
	Now() (Private, error)
}

// Private represents a private unit
type Private interface {
	Hash() hash.Hash
	Unit() genesis.Unit
	PrivateKey() signature.PrivateKey
}
