package transactions

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/genesis"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/rings"
)

// Builder represents a transactions builder
type Builder interface {
	Create() Builder
	WithList(list []Transaction) Builder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	List() []Transaction
}

// TransactionBuilder represents a transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithContent(content Content) TransactionBuilder
	WithAuthorization(auth signature.RingSignature) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Content() Content
	Authorization() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithOwner(owner rings.Ring) ContentBuilder
	WithAmount(amount hash.Hash) ContentBuilder
	WithFees(fees uint) ContentBuilder
	WithOrigin(origin Origin) ContentBuilder
	WithExternal(external hash.Hash) ContentBuilder
	WithSides(sides Transactions) ContentBuilder
	Now() (Content, error)
}

// Content represents the content
type Content interface {
	Hash() hash.Hash
	Owner() rings.Ring
	Amount() hash.Hash
	Fees() uint
	Origin() Origin
	CreatedOn() time.Time
	HasExternal() bool
	External() *hash.Hash
	HasSides() bool
	Sides() Transactions
}

// OriginBuilder represents the origin builder
type OriginBuilder interface {
	Create() OriginBuilder
	WithGenesis(genesis genesis.Genesis) OriginBuilder
	WithTransaction(trx Transaction) OriginBuilder
	Now() (Origin, error)
}

// Origin represents a transaction origin
type Origin interface {
	Hash() hash.Hash
	IsGenesis() bool
	Genesis() genesis.Genesis
	IsTransaction() bool
	Transaction() Transaction
}
