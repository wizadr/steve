package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

type specifierBuilder struct {
	containerName string
	cardinality   cardinality.Specific
	result        TokenResult
}

func createSpecifierBuilder() SpecifierBuilder {
	out := specifierBuilder{
		containerName: "",
		cardinality:   nil,
		result:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *specifierBuilder) Create() SpecifierBuilder {
	return createSpecifierBuilder()
}

// WithContainerName adds a containerName to the builder
func (app *specifierBuilder) WithContainerName(containerName string) SpecifierBuilder {
	app.containerName = containerName
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *specifierBuilder) WithCardinality(cardinality cardinality.Specific) SpecifierBuilder {
	app.cardinality = cardinality
	return app
}

// WithResult adds a token result to the builder
func (app *specifierBuilder) WithResult(result TokenResult) SpecifierBuilder {
	app.result = result
	return app
}

// Now builds a new Specifier instance
func (app *specifierBuilder) Now() (Specifier, error) {
	if app.containerName == "" {
		return nil, errors.New("the containerName is mandatory in order to build a Specifier instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a Specifier instance")
	}

	if app.result == nil {
		return nil, errors.New("the TokenResult is mandatory in order to build a Specifier instance")
	}

	return createSpecifier(app.containerName, app.cardinality, app.result), nil
}
