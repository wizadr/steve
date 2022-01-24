package chains

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/digital-diamonds/domain/owners"
	"github.com/steve-care-software/digital-diamonds/domain/transactions"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
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
	WithFees(fees owners.Owners) BlockBuilder
	WithTransactions(transactions transactions.Transactions) BlockBuilder
	WithPrevious(previous hash.Hash) BlockBuilder
	CreatedOn(createdOn time.Time) BlockBuilder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Height() uint
	Fees() owners.Owners
	Transactions() transactions.Transactions
	CreatedOn() time.Time
	HasPrevious() bool
	Previous() *hash.Hash
}
