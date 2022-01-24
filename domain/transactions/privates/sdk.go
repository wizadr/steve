package privates

import (
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/receipts"
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
