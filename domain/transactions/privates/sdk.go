package privates

import (
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

// Builder represents the privates builder
type Builder interface {
	Create() Builder
	WithList(list []Private) Builder
	Now() (Privates, error)
}

// Privates represents private transactions
type Privates interface {
	Hash() hash.Hash
	All() []Private
}

// PrivateBuilder represents a private builder
type PrivateBuilder interface {
	Create() PrivateBuilder
	WithOwner(owner signature.PrivateKey) PrivateBuilder
	WithSecret(secret secrets.Secret) PrivateBuilder
	WithOrigin(origin Private) PrivateBuilder
	WithSides(sides Privates) PrivateBuilder
	Now() (Private, error)
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
