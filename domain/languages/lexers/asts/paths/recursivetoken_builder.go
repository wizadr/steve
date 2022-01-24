package paths

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

type recursiveTokenBuilder struct {
	name        string
	cardinality cardinality.Cardinality
	specifiers  []Specifier
}

func createRecursiveTokenBuilder() RecursiveTokenBuilder {
	out := recursiveTokenBuilder{
		name:        "",
		cardinality: nil,
		specifiers:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *recursiveTokenBuilder) Create() RecursiveTokenBuilder {
	return createRecursiveTokenBuilder()
}

// WithName adds a name to the builder
func (app *recursiveTokenBuilder) WithName(name string) RecursiveTokenBuilder {
	app.name = name
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *recursiveTokenBuilder) WithCardinality(cardinality cardinality.Cardinality) RecursiveTokenBuilder {
	app.cardinality = cardinality
	return app
}

// WithSpecifiers add specifiers to the builder
func (app *recursiveTokenBuilder) WithSpecifiers(specifiers []Specifier) RecursiveTokenBuilder {
	app.specifiers = specifiers
	return app
}

// Now builds a new RecursiveToken instance
func (app *recursiveTokenBuilder) Now() (RecursiveToken, error) {
	if app.specifiers != nil && len(app.specifiers) <= 0 {
		app.specifiers = nil
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a RecursiveToken instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a RecursiveToken instance")
	}

	if len(app.specifiers) > 0 {
		return createRecursiveTokenWithSpecifiers(app.name, app.cardinality, app.specifiers), nil
	}

	return createRecursiveToken(app.name, app.cardinality), nil
}
