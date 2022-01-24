package secrets

import (
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/transactions"
)

// Builder represents a secrets builder
type Builder interface {
	Create() Builder
	WithList(list []Secret) Builder
	Now() (Secrets, error)
}

// Secrets represents secret transfers
type Secrets interface {
	Hash() hash.Hash
	All() []Secret
}

// SecretBuilder represents the secret builder
type SecretBuilder interface {
	Create() SecretBuilder
	WithAmount(amount uint) SecretBuilder
	WithNonce(nonce string) SecretBuilder
	WithPublic(public transactions.Transaction) SecretBuilder
	WithOrigin(origin Secret) SecretBuilder
	WithSides(sides Secrets) SecretBuilder
	Now() (Secret, error)
}

// Secret represents the secret part of a transfer
type Secret interface {
	Hash() hash.Hash
	Amount() uint
	Nonce() string
	Public() transactions.Transaction
	HasOrigin() bool
	Origin() Secret
	HasSides() bool
	Sides() Secrets
}
