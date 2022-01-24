package transactions

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/genesis"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/owners"
)

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	All() []Transaction
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Content() Content
	Authorization() signature.RingSignature
}

// Content represents the content
type Content interface {
	Hash() hash.Hash
	Owner() owners.Owner
	Amount() hash.Hash
	Fees() uint
	Origin() Origin
	CreatedOn() time.Time
	HasExternal() bool
	External() hash.Hash
	HasSides() bool
	Sides() Transactions
}

// Origin represents a transaction origin
type Origin interface {
	Hash() hash.Hash
	IsGenesis() bool
	Genesis() genesis.Genesis
	IsTransaction() bool
	Transaction() Transaction
}
