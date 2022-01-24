package identities

import (
	"github.com/steve-care-software/steve/domain/transactions/privates"
	"github.com/steve-care-software/steve/domain/transactions/secrets"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewTransactionsBuilder creates a new transactions builder
func NewTransactionsBuilder() TransactionsBuilder {
	return createTransactionsBuilder()
}

// Builder represents the identity builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithTransactions(trx Transactions) Builder
	WithGenesis(genesis privates.Privates) Builder
	Now() (Identity, error)
}

// Identity represents the identity
type Identity interface {
	Name() string
	Transactions() Transactions
	HasGenesis() bool
	Genesis() privates.Privates
}

// TransactionsBuilder represents the transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithOutgoing(outgoing privates.Privates) TransactionsBuilder
	WithIncoming(incoming secrets.Secrets) TransactionsBuilder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	HasIncoming() bool
	Incoming() secrets.Secrets
	HasOutgoing() bool
	Outgoing() privates.Privates
}
