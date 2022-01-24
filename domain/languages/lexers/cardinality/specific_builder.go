package cardinality

import (
	"errors"
)

type specificBuilder struct {
	amount *uint
	rnge   Range
}

func createSpecificBuilder() SpecificBuilder {
	out := specificBuilder{
		amount: nil,
		rnge:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *specificBuilder) Create() SpecificBuilder {
	return createSpecificBuilder()
}

// WithAmount adds an amount to the builder
func (app *specificBuilder) WithAmount(amount uint) SpecificBuilder {
	app.amount = &amount
	return app
}

// WithRange adds a range to the builder
func (app *specificBuilder) WithRange(rnge Range) SpecificBuilder {
	app.rnge = rnge
	return app
}

// Now builds a new Specific instance
func (app *specificBuilder) Now() (Specific, error) {
	if app.amount != nil {
		amount := *app.amount
		if amount <= 0 {
			return nil, errors.New("the amount must be greater than zero in order to build a Specific instance")
		}

		return createSpecificWithAmount(app.amount), nil
	}

	if app.rnge != nil {
		return createSpecificWithRange(app.rnge), nil
	}

	return nil, errors.New("the Specific instance is invalid")
}
