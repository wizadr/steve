package roots

import "errors"

type containerBuilder struct {
	content string
	name    string
	root    Root
}

func createContainerBuilder() ContainerBuilder {
	out := containerBuilder{
		content: "",
		name:    "",
		root:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *containerBuilder) Create() ContainerBuilder {
	return createContainerBuilder()
}

// WithContent adds content to the builder
func (app *containerBuilder) WithContent(content string) ContainerBuilder {
	app.content = content
	return app
}

// WithName adds a name to the builder
func (app *containerBuilder) WithName(name string) ContainerBuilder {
	app.name = name
	return app
}

// WithRoot adds a root to the builder
func (app *containerBuilder) WithRoot(root Root) ContainerBuilder {
	app.root = root
	return app
}

// Now builds a new Container instance
func (app *containerBuilder) Now() (Container, error) {
	if app.content == "" {
		return nil, errors.New("the content is mandatory in order to build a Container instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Container instance")
	}

	if app.root != nil {
		return createContainerWithRoot(app.content, app.name, app.root), nil
	}

	return createContainer(app.content, app.name), nil
}
