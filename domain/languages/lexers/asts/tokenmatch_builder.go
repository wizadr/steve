package asts

import "errors"

type tokenMatchBuilder struct {
	token   string
	matches []Token
}

func createTokenMatchBuilder() TokenMatchBuilder {
	out := tokenMatchBuilder{
		token:   "",
		matches: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenMatchBuilder) Create() TokenMatchBuilder {
	return createTokenMatchBuilder()
}

// WithToken adds a token to the builder
func (app *tokenMatchBuilder) WithToken(token string) TokenMatchBuilder {
	app.token = token
	return app
}

// WithMatches add matches to the builder
func (app *tokenMatchBuilder) WithMatches(matches []Token) TokenMatchBuilder {
	app.matches = matches
	return app
}

// Now builds a new TokenMatch instance
func (app *tokenMatchBuilder) Now() (TokenMatch, error) {
	if app.token == "" {
		return nil, errors.New("the token is mandatory in order to build a TokenMatch instance")
	}

	if app.matches != nil && len(app.matches) <= 0 {
		app.matches = nil
	}

	if app.matches == nil {
		return nil, errors.New("there must be at least 1 Token match in order to build a TokenMatch instance")
	}

	return createTokenMatch(app.token, app.matches), nil
}
