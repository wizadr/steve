package paths

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

type tokenBuilder struct {
	path        TokenPath
	cardinality cardinality.Cardinality
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		path:        nil,
		cardinality: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithPath adds a path to the builder
func (app *tokenBuilder) WithPath(path TokenPath) TokenBuilder {
	app.path = path
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *tokenBuilder) WithCardinality(cardinality cardinality.Cardinality) TokenBuilder {
	app.cardinality = cardinality
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.path == nil {
		return nil, errors.New("the RulePath is mandatory in order to build a Token instsnace")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a Token instance")
	}

	return createToken(app.path, app.cardinality), nil
}
