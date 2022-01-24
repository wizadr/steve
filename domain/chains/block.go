package chains

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/owners"
	"github.com/steve-care-software/digital-diamonds/domain/transactions"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type block struct {
	hash         hash.Hash
	height       uint
	fees         owners.Owners
	transactions transactions.Transactions
	createdOn    time.Time
	previous     *hash.Hash
}

func createBlock(
	hash hash.Hash,
	height uint,
	fees owners.Owners,
	transactions transactions.Transactions,
	createdOn time.Time,
) Block {
	return createBlockInternally(hash, height, fees, transactions, createdOn, nil)
}

func createBlockWithPrevious(
	hash hash.Hash,
	height uint,
	fees owners.Owners,
	transactions transactions.Transactions,
	createdOn time.Time,
	previous *hash.Hash,
) Block {
	return createBlockInternally(hash, height, fees, transactions, createdOn, previous)
}

func createBlockInternally(
	hash hash.Hash,
	height uint,
	fees owners.Owners,
	transactions transactions.Transactions,
	createdOn time.Time,
	previous *hash.Hash,
) Block {
	out := block{
		hash:         hash,
		height:       height,
		fees:         fees,
		transactions: transactions,
		createdOn:    createdOn,
		previous:     previous,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Fees returns the fees owners
func (obj *block) Fees() owners.Owners {
	return obj.fees
}

// Height returns the height
func (obj *block) Height() uint {
	return obj.height
}

// Transactions returns the transactions
func (obj *block) Transactions() transactions.Transactions {
	return obj.transactions
}

// CreatedOn returns the creation time
func (obj *block) CreatedOn() time.Time {
	return obj.createdOn
}

// HasPrevious returns true if there is a previous block, false otherwise
func (obj *block) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous block hash,if any
func (obj *block) Previous() *hash.Hash {
	return obj.previous
}
