package asts

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
)

type builder struct {
	mistake mistakes.Mistake
	token   Token
}

func createBuilder() Builder {
	out := builder{
		mistake: nil,
		token:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMistake adds a mistake to the builder
func (app *builder) WithMistake(mistake mistakes.Mistake) Builder {
	app.mistake = mistake
	return app
}

// WithSuccess adds a successful Token to the builder
func (app *builder) WithSuccess(token Token) Builder {
	app.token = token
	return app
}

// Now builds a new AST instance
func (app *builder) Now() (AST, error) {
	if app.mistake != nil {
		return createASTWithMistake(app.mistake), nil
	}

	if app.token != nil {
		return createASTWithSuccess(app.token), nil
	}

	return nil, errors.New("the AST is invalid")
}
