package wallets

import (
	"time"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/languages/lexers/resources/singles"
	"github.com/steve-care-software/steve/domain/transactions"
)

// Wallets represents wallets
type Wallets interface {
	List() []Wallet
}

// Wallet represents a wallet
type Wallet interface {
	Hash() hash.Hash
	Content() Content
	Transaction() transactions.Transaction
}

// Content represents a wallet content
type Content interface {
	Hash() hash.Hash
	Name() string
	Version() string
	Description() string
	Resource() singles.Single
	Script() string
	CreatedOn() time.Time
	HasDependencies() bool
	Dependencies() Wallets
	HasParent() bool
	Parent() Wallet
}
