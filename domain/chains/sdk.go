package chains

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/rings"
	"github.com/steve-care-software/steve/domain/transactions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	hashAdapter := hash.NewAdapter()
	return createBlockBuilder(hashAdapter)
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithID(id uuid.UUID) Builder
	WithRoot(root hash.Hash) Builder
	WithHead(head Block) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Chain, error)
}

// Chain represents a blockchain
type Chain interface {
	ID() uuid.UUID
	Hash() hash.Hash
	Root() hash.Hash
	CreatedOn() time.Time
	HasHead() bool
	Head() Block
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithHeight(height uint) BlockBuilder
	WithFees(fees rings.Ring) BlockBuilder
	WithTransactions(transactions transactions.Transactions) BlockBuilder
	WithPrevious(previous hash.Hash) BlockBuilder
	CreatedOn(createdOn time.Time) BlockBuilder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Height() uint
	Fees() rings.Ring
	Transactions() transactions.Transactions
	CreatedOn() time.Time
	HasPrevious() bool
	Previous() *hash.Hash
}
