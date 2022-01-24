package patterns

import "errors"

type discoveryBuilder struct {
	index     *uint
	content   string
	container Container
}

func createDiscoveryBuilder() DiscoveryBuilder {
	out := discoveryBuilder{
		index:     nil,
		content:   "",
		container: nil,
	}

	return &out
}

// Create initializes the builder
func (app *discoveryBuilder) Create() DiscoveryBuilder {
	return createDiscoveryBuilder()
}

// WithIndex adds an index to the builder
func (app *discoveryBuilder) WithIndex(index uint) DiscoveryBuilder {
	app.index = &index
	return app
}

// WithContent adds content to the builder
func (app *discoveryBuilder) WithContent(content string) DiscoveryBuilder {
	app.content = content
	return app
}

// WithContainer adds a container to the builder
func (app *discoveryBuilder) WithContainer(container Container) DiscoveryBuilder {
	app.container = container
	return app
}

// Now builds a new Discovery instance
func (app *discoveryBuilder) Now() (Discovery, error) {
	if app.index == nil {
		return nil, errors.New("the index is mandatory in order to build a Discovery instance")
	}

	if app.content == "" {
		return nil, errors.New("the content is mandatory in order to build a Discovery instance")
	}

	if app.container == nil {
		return nil, errors.New("the container is mandatory in order to build a Discovery instance")
	}

	return createDiscovery(*app.index, app.content, app.container), nil
}
