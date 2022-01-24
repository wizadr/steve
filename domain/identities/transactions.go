package identities

import (
	"github.com/steve-care-software/digital-diamonds/domain/transactions/privates"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

type transactions struct {
	outgoing privates.Privates
	incoming secrets.Secrets
}

func createTransactions() Transactions {
	return createTransactionsInternally(nil, nil)
}

func createTransactionsWithOutgoing(
	outgoing privates.Privates,
) Transactions {
	return createTransactionsInternally(outgoing, nil)
}

func createTransactionsWithIncoming(
	incoming secrets.Secrets,
) Transactions {
	return createTransactionsInternally(nil, incoming)
}

func createTransactionsWithOutgoingAndIncoming(
	outgoing privates.Privates,
	incoming secrets.Secrets,
) Transactions {
	return createTransactionsInternally(outgoing, incoming)
}

func createTransactionsInternally(
	outgoing privates.Privates,
	incoming secrets.Secrets,
) Transactions {
	out := transactions{
		outgoing: outgoing,
		incoming: incoming,
	}

	return &out
}

// HasOutgoing returns true if there is outgoing transactions, false otherwise
func (obj *transactions) HasOutgoing() bool {
	return obj.outgoing != nil
}

// Outgoing returns the outgoing transactions
func (obj *transactions) Outgoing() privates.Privates {
	return obj.outgoing
}

// HasIncoming returns true if there is incoming transactions, false otherwise
func (obj *transactions) HasIncoming() bool {
	return obj.incoming != nil
}

// Incoming returns the incoming transactions
func (obj *transactions) Incoming() secrets.Secrets {
	return obj.incoming
}
