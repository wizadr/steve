package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

type tokenResultBuilder struct {
	input       string
	cardinality cardinality.Cardinality
	matches     TokenResultMatches
}

func createTokenResultBuilder() TokenResultBuilder {
	out := tokenResultBuilder{
		input:       "",
		cardinality: nil,
		matches:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenResultBuilder) Create() TokenResultBuilder {
	return createTokenResultBuilder()
}

// WithInput adds an input to the builder
func (app *tokenResultBuilder) WithInput(input string) TokenResultBuilder {
	app.input = input
	return app
}

// WithCardinality adds cardinality to the builder
func (app *tokenResultBuilder) WithCardinality(cardinality cardinality.Cardinality) TokenResultBuilder {
	app.cardinality = cardinality
	return app
}

// WithMatches add matches to the builder
func (app *tokenResultBuilder) WithMatches(matches TokenResultMatches) TokenResultBuilder {
	app.matches = matches
	return app
}

// Now builds a new TokenResult instance
func (app *tokenResultBuilder) Now() (TokenResult, error) {
	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a TokenResult instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a TokenResult instance")
	}

	if app.matches != nil {
		return createTokenResultWithMatches(app.input, app.cardinality, app.matches), nil
	}

	return createTokenResult(app.input, app.cardinality), nil
}
