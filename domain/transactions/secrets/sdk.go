package secrets

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transactions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewSecretBuilder creates a new secret builder
func NewSecretBuilder() SecretBuilder {
	hashAdapter := hash.NewAdapter()
	return createSecretBuilder(hashAdapter)
}

// Builder represents a secrets builder
type Builder interface {
	Create() Builder
	WithList(list []Secret) Builder
	Now() (Secrets, error)
}

// Secrets represents secret transfers
type Secrets interface {
	Hash() hash.Hash
	List() []Secret
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
