package privates

import (
	"github.com/steve-care-software/steve/domain/cryptography/keys/signature"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/receipts"
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

// Privates represents private transactions
type Privates interface {
	Hash() hash.Hash
	List() []Private
}

// PrivateBuilder represents a private builder
type PrivateBuilder interface {
	Create() PrivateBuilder
	WithPrivateKey(pk signature.PrivateKey) PrivateBuilder
	WithReceipt(receipt receipts.Receipt) PrivateBuilder
	Now() (Private, error)
}

// Private represents the private part of a transfer
type Private interface {
	Hash() hash.Hash
	PrivateKey() signature.PrivateKey
	Receipt() receipts.Receipt
}
