package identities

import (
	"errors"

	"github.com/steve-care-software/steve/domain/transactions/privates"
)

type builder struct {
	name    string
	trx     Transactions
	genesis privates.Privates
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		trx:     nil,
		genesis: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(trx Transactions) Builder {
	app.trx = trx
	return app
}

// WithGenesis add genesis transactions to the builder
func (app *builder) WithGenesis(genesis privates.Privates) Builder {
	app.genesis = genesis
	return app
}

// Nwo builds a new Idenity instance
func (app *builder) Now() (Identity, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transactions is mandatory in order to build an Identity instance")
	}

	if app.genesis != nil {
		return createIdentityWithGenesis(app.name, app.trx, app.genesis), nil
	}

	return createIdentity(app.name, app.trx), nil
}
