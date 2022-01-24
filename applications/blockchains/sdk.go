package blockchains

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/steve/domain/chains"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/peers"
	"github.com/steve-care-software/steve/domain/transactions"
)

// Init returns the init public func
func Init(name string, password string, feesAmount uint, unitsAmount uint, chunksAmount uint) error {
	return nil
}

// Application represents a blockchain application
type Application interface {
	Chains(id uuid.UUID) chains.Chain
	Block(blockHash hash.Hash) (chains.Block, error)
	Queue(index uint, amount uint) (transactions.Transactions, error)
	Peers() (peers.Peers, error)
	Transact(trx transactions.Transaction) error
}
