package paths

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"
)

type builder struct {
	rootToken string
	container containers.Container
}

func createBuilder() Builder {
	out := builder{
		rootToken: "",
		container: nil,
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

// WithContainer adds a container to the builder
func (app *builder) WithContainer(container containers.Container) Builder {
	app.container = container
	return app
}

// Now builds a new Path instance
func (app *builder) Now() (Path, error) {
	if app.rootToken == "" {
		return nil, errors.New("the rootToken is mandatory in order to build a Path instance")
	}

	if app.container == nil {
		return nil, errors.New("the container is mandatory in order to build a Path instance")
	}

	return createPath(app.rootToken, app.container), nil
}
