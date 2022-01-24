package mistakes

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

type specifierDoNotMatchBuilder struct {
	containerName string
	cardinality   cardinality.Specific
	amount        int
}

func createSpecifierDoNotMatchBuilder() SpecifierDoNotMatchBuilder {
	out := specifierDoNotMatchBuilder{
		containerName: "",
		cardinality:   nil,
		amount:        -1,
	}

	return &out
}

// Create initializes the builder
func (app *specifierDoNotMatchBuilder) Create() SpecifierDoNotMatchBuilder {
	return createSpecifierDoNotMatchBuilder()
}

// WithContainerName adds a containerName to the builder
func (app *specifierDoNotMatchBuilder) WithContainerName(containerName string) SpecifierDoNotMatchBuilder {
	app.containerName = containerName
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *specifierDoNotMatchBuilder) WithCardinality(cardinality cardinality.Specific) SpecifierDoNotMatchBuilder {
	app.cardinality = cardinality
	return app
}

// WithAmount adds an amount to the builder
func (app *specifierDoNotMatchBuilder) WithAmount(amount uint) SpecifierDoNotMatchBuilder {
	app.amount = int(amount)
	return app
}

// Now builds a new SpecifierDoNotMatch instance
func (app *specifierDoNotMatchBuilder) Now() (SpecifierDoNotMatch, error) {
	if app.containerName == "" {
		return nil, errors.New("the containerName is mandatory in order to build a SpecifierDoNotMatch instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a SpecifierDoNotMatch instance")
	}

	if app.amount < 0 {
		return nil, errors.New("the amount is mandatory in order to build a SpecifierDoNotMatch instance")
	}

	return createSpecifierDoNotMatch(app.containerName, app.cardinality, uint(app.amount)), nil
}
