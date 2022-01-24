package identities

import (
	"github.com/steve-care-software/digital-diamonds/domain/transactions/privates"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

type transactionsBuilder struct {
	outgoing privates.Privates
	incoming secrets.Secrets
}

func createTransactionsBuilder() TransactionsBuilder {
	out := transactionsBuilder{
		outgoing: nil,
		incoming: nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionsBuilder) Create() TransactionsBuilder {
	return createTransactionsBuilder()
}

// WithOutgoing adds an outgoing to the builder
func (app *transactionsBuilder) WithOutgoing(outgoing privates.Privates) TransactionsBuilder {
	app.outgoing = outgoing
	return app
}

// WithIncoming adds an incoming to the builder
func (app *transactionsBuilder) WithIncoming(incoming secrets.Secrets) TransactionsBuilder {
	app.incoming = incoming
	return app
}

// Now builds a new Transactions instance
func (app *transactionsBuilder) Now() (Transactions, error) {
	if app.outgoing != nil && app.incoming != nil {
		return createTransactionsWithOutgoingAndIncoming(app.outgoing, app.incoming), nil
	}

	if app.outgoing != nil {
		return createTransactionsWithOutgoing(app.outgoing), nil
	}

	if app.incoming != nil {
		return createTransactionsWithIncoming(app.incoming), nil
	}

	return createTransactions(), nil
}
