package blockchains

import (
	"github.com/steve-care-software/digital-diamonds/domain/chains"
	"github.com/steve-care-software/digital-diamonds/domain/peers"
	"github.com/steve-care-software/digital-diamonds/domain/transactions"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// Application represents a blockchain application
type Application interface {
	Chains() chains.Chain
	Block(blockHash hash.Hash) (chains.Block, error)
	Queue(index uint, amount uint) (transactions.Transactions, error)
	Peers() (peers.Peers, error)
	Transact(trx transactions.Transaction) error
}
