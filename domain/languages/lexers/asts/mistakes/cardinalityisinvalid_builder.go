package mistakes

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

type cardinalityIsInvalidBuilder struct {
	cardinality cardinality.Cardinality
	amount      int
}

func createCardinalityIsInvalidBuilder() CardinalityIsInvalidBuilder {
	out := cardinalityIsInvalidBuilder{
		cardinality: nil,
		amount:      -1,
	}

	return &out
}

// Create initializes the builder
func (app *cardinalityIsInvalidBuilder) Create() CardinalityIsInvalidBuilder {
	return createCardinalityIsInvalidBuilder()
}

// WithCardinality adds a cardinality to the builder
func (app *cardinalityIsInvalidBuilder) WithCardinality(cardinality cardinality.Cardinality) CardinalityIsInvalidBuilder {
	app.cardinality = cardinality
	return app
}

// WithAmount adds an amount to the builder
func (app *cardinalityIsInvalidBuilder) WithAmount(amount uint) CardinalityIsInvalidBuilder {
	app.amount = int(amount)
	return app
}

// Now builds a new CardinalityIsInvalid instance
func (app *cardinalityIsInvalidBuilder) Now() (CardinalityIsInvalid, error) {
	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a CardinalityIsInvalid instance")
	}

	if app.amount < 0 {
		return nil, errors.New("the amount is mandatory in order to build a CardinalityIsInvalid instance")
	}

	return createCardinalityIsInvalid(app.cardinality, uint(app.amount)), nil
}
