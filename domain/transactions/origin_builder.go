package transactions

import (
	"errors"

	"github.com/steve-care-software/steve/domain/genesis"
)

type originBuilder struct {
	genesis genesis.Unit
	trx     Transaction
}

func createOriginBuilder() OriginBuilder {
	out := originBuilder{
		genesis: nil,
		trx:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *originBuilder) Create() OriginBuilder {
	return createOriginBuilder()
}

// WithGenesis adds a genesis to the builder
func (app *originBuilder) WithGenesis(genesis genesis.Unit) OriginBuilder {
	app.genesis = genesis
	return app
}

// WithTransaction adds a transaction to the builder
func (app *originBuilder) WithTransaction(trx Transaction) OriginBuilder {
	app.trx = trx
	return app
}

// Now builds a new Origin instance
func (app *originBuilder) Now() (Origin, error) {
	if app.genesis != nil {
		return createOriginWithGenesis(app.genesis), nil
	}

	if app.trx != nil {
		return createOriginWithTransaction(app.trx), nil
	}

	return nil, errors.New("the Origin is invalid")
}
