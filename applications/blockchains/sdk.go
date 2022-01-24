package blockchains

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/digital-diamonds/domain/chains"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/peers"
	"github.com/steve-care-software/digital-diamonds/domain/transactions"
)

// Application represents a blockchain application
type Application interface {
	Init(name string, password string, amount uint) error
	Chains(id uuid.UUID) chains.Chain
	Block(blockHash hash.Hash) (chains.Block, error)
	Queue(index uint, amount uint) (transactions.Transactions, error)
	Peers() (peers.Peers, error)
	Transact(trx transactions.Transaction) error
}
