package transactions

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/genesis"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/rings"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewTransactionBuilder creates a new transaction builder
func NewTransactionBuilder() TransactionBuilder {
	hashAdapter := hash.NewAdapter()
	return createTranactionBuilder(hashAdapter)
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// NewOriginBuilder creates a new origin builder
func NewOriginBuilder() OriginBuilder {
	return createOriginBuilder()
}

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
	WithOrigin(origin Origin) ContentBuilder
	WithExternal(external hash.Hash) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents the content
type Content interface {
	Hash() hash.Hash
	Owner() rings.Ring
	Amount() hash.Hash
	Origin() Origin
	External() hash.Hash
	CreatedOn() time.Time
}

// OriginBuilder represents the origin builder
type OriginBuilder interface {
	Create() OriginBuilder
	WithGenesis(genesis genesis.Unit) OriginBuilder
	WithTransaction(trx Transaction) OriginBuilder
	Now() (Origin, error)
}

// Origin represents a transaction origin
type Origin interface {
	Hash() hash.Hash
	IsGenesis() bool
	Genesis() genesis.Unit
	IsTransaction() bool
	Transaction() Transaction
}
