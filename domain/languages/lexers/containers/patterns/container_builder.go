package patterns

import "errors"

type containerBuilder struct {
	reverse Serie
	group   Group
}

func createContainerBuilder() ContainerBuilder {
	out := containerBuilder{
		reverse: nil,
		group:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *containerBuilder) Create() ContainerBuilder {
	return createContainerBuilder()
}

// WithReverse adds a reverse to the builder
func (app *containerBuilder) WithReverse(reverse Serie) ContainerBuilder {
	app.reverse = reverse
	return app
}

// WithGroup adds a group to the builder
func (app *containerBuilder) WithGroup(group Group) ContainerBuilder {
	app.group = group
	return app
}

// Now builds a new Container instance
func (app *containerBuilder) Now() (Container, error) {
	if app.reverse != nil {
		return createContainerWithReverse(app.reverse), nil
	}

	if app.group != nil {
		return createContainerWithGroup(app.group), nil
	}

	return nil, errors.New("the container is invalid")
}
