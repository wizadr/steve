package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

type tokenSpecifierBuilder struct {
	content     TokenSpecifierContent
	cardinality cardinality.Specific
}

func createTokenSpecifierBuilder() TokenSpecifierBuilder {
	out := tokenSpecifierBuilder{
		content:     nil,
		cardinality: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenSpecifierBuilder) Create() TokenSpecifierBuilder {
	return createTokenSpecifierBuilder()
}

// WithContent adds content to the builder
func (app *tokenSpecifierBuilder) WithContent(content TokenSpecifierContent) TokenSpecifierBuilder {
	app.content = content
	return app
}

// WithCardinality adds cardinality to the builder
func (app *tokenSpecifierBuilder) WithCardinality(cardinality cardinality.Specific) TokenSpecifierBuilder {
	app.cardinality = cardinality
	return app
}

// Now builds a new TokenSpecifier instance
func (app *tokenSpecifierBuilder) Now() (TokenSpecifier, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a TokenSpecifier instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the specificCardinality is mandatory in order to build a TokenSpecifier instance")
	}

	return createTokenSpecifier(app.content, app.cardinality), nil
}
