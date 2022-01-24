package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
)

type builder struct {
	name    string
	path    paths.Path
	content TokenContent
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		path:    nil,
		content: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path paths.Path) Builder {
	app.path = path
	return app
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content TokenContent) Builder {
	app.content = content
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Token instance")
	}

	if app.content != nil {
		return createTokenWithContent(app.name, app.path, app.content), nil
	}

	return createToken(app.name, app.path), nil
}
