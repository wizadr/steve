package singles

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/resources/multiples"
)

type builder struct {
	rootToken string
	multiple  multiples.Multiple
}

func createBuilder() Builder {
	out := builder{
		rootToken: "",
		multiple:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRootToken adds a rootToken to the builder
func (app *builder) WithRootToken(rootToken string) Builder {
	app.rootToken = rootToken
	return app
}

// WithMultiple adds a multiple to the builder
func (app *builder) WithMultiple(multiple multiples.Multiple) Builder {
	app.multiple = multiple
	return app
}

// Now builds a new Single instance
func (app *builder) Now() (Single, error) {
	if app.rootToken == "" {
		return nil, errors.New("the rootToken is mandatory in order to build a Single instance")
	}

	if app.multiple == nil {
		return nil, errors.New("the multiple is mandatory in order to build a Single instance")
	}

	return createSingle(app.rootToken, app.multiple), nil
}
