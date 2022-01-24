package multiples

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"
)

type builder struct {
	container containers.Container
	channels  string
}

func createBuilder() Builder {
	out := builder{
		container: nil,
		channels:  "",
	}

	return &out
}

// Create returns true if there is a builder, false otherwise
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithContainer adds a container to the builder
func (app *builder) WithContainer(container containers.Container) Builder {
	app.container = container
	return app
}

// WithChannels adds channels to the builder
func (app *builder) WithChannels(channels string) Builder {
	app.channels = channels
	return app
}

// Now builds a new Multiple instance
func (app *builder) Now() (Multiple, error) {
	if app.container == nil {
		return nil, errors.New("the container is mandatory in order to build a Multiple instance")
	}

	if app.channels != "" {
		return createMultipleWithChannels(app.container, app.channels), nil
	}

	return createMultiple(app.container), nil
}
