package privates

import (
	"github.com/steve-care-software/steve/domain/hash"
	transaction "github.com/steve-care-software/steve/domain/transactions/privates"
	"github.com/steve-care-software/steve/domain/wallets"
)

// Privates represents private wallets
type Privates interface {
	Hash() hash.Hash
	List() []Private
}

// Private represents a private wallet
type Private interface {
	Hash() hash.Hash
	Transaction() transaction.Private
	Wallet() wallets.Wallet
}
