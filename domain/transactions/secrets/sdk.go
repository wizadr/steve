package secrets

import (
	"github.com/steve-care-software/digital-diamonds/domain/transactions"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// Secrets represents secret transfers
type Secrets interface {
	Hash() hash.Hash
	All() []Secret
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
