package asts

import (
	"errors"
)

type tokenBuilder struct {
	name  string
	match LineMatch
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		name:  "",
		match: nil,
	}

	return &out
}

// Create initializes the tokenBuilder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithName adds a name to the tokenBuilder
func (app *tokenBuilder) WithName(name string) TokenBuilder {
	app.name = name
	return app
}

// WithMatch adds a match to the tokenBuilder
func (app *tokenBuilder) WithMatch(match LineMatch) TokenBuilder {
	app.match = match
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.match == nil {
		return nil, errors.New("the LineMatch is mandatory in order to build a Token instance")
	}

	return createToken(app.name, app.match), nil
}
