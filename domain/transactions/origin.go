package transactions

import (
	"github.com/steve-care-software/digital-diamonds/domain/genesis"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type origin struct {
	genesis genesis.Unit
	trx     Transaction
}

func createOriginWithGenesis(
	genesis genesis.Unit,
) Origin {
	return createOriginInternally(genesis, nil)
}

func createOriginWithTransaction(
	trx Transaction,
) Origin {
	return createOriginInternally(nil, trx)
}

func createOriginInternally(
	genesis genesis.Unit,
	trx Transaction,
) Origin {
	out := origin{
		genesis: genesis,
		trx:     trx,
	}

	return &out
}

// Hash returns the hash
func (obj *origin) Hash() hash.Hash {
	if obj.IsGenesis() {
		return obj.genesis.Hash()
	}

	return obj.trx.Hash()
}

// IsGenesis returns true if there is genesis, false otherwise
func (obj *origin) IsGenesis() bool {
	return obj.genesis != nil
}

// Genesis returns the genesis, if any
func (obj *origin) Genesis() genesis.Unit {
	return obj.genesis
}

// IsTransaction returns true if there is a transaction, false otherwise
func (obj *origin) IsTransaction() bool {
	return obj.trx != nil
}

// Transaction returns the transaction, if any
func (obj *origin) Transaction() Transaction {
	return obj.trx
}
