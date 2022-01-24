package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
)

type tokenMatchBuilder struct {
	path       paths.TokenPath
	result     TokenResult
	specifiers []Specifier
}

func createTokenMatchBuilder() TokenMatchBuilder {
	out := tokenMatchBuilder{
		path:       nil,
		result:     nil,
		specifiers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenMatchBuilder) Create() TokenMatchBuilder {
	return createTokenMatchBuilder()
}

// WithPath add path to the builder
func (app *tokenMatchBuilder) WithPath(path paths.TokenPath) TokenMatchBuilder {
	app.path = path
	return app
}

// WithResult adds a result to the builder
func (app *tokenMatchBuilder) WithResult(result TokenResult) TokenMatchBuilder {
	app.result = result
	return app
}

// WithSpecifiers add specifiers to the builder
func (app *tokenMatchBuilder) WithSpecifiers(specifiers []Specifier) TokenMatchBuilder {
	app.specifiers = specifiers
	return app
}

// Now builds a new TokenMatch instance
func (app *tokenMatchBuilder) Now() (TokenMatch, error) {
	if app.path == nil {
		return nil, errors.New("the TokenPath is mandatory in order to build a TokenMatch instance")
	}

	if app.result == nil {
		return nil, errors.New("the TokenResult is mandatory in order to build a TokenMatch instance")
	}

	if app.specifiers != nil && len(app.specifiers) <= 0 {
		app.specifiers = nil
	}

	if app.specifiers != nil {
		return createTokenMatchWithSpecifiers(app.path, app.result, app.specifiers), nil
	}

	return createTokenMatch(app.path, app.result), nil
}
