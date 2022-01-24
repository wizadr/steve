package wallets

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/wallets"
)

// Application represents a wallet application
type Application interface {
	Save(wallet wallets.Wallet) error
	Retrieve(index uint, amount uint) (wallets.Wallets, error)
	RetrieveByHash(hash hash.Hash) (wallets.Wallet, error)
}
