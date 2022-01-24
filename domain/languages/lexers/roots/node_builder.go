package roots

import "errors"

type nodeBuilder struct {
	name       string
	containers Containers
}

func createNodeBuilder() NodeBuilder {
	out := nodeBuilder{
		name:       "",
		containers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *nodeBuilder) Create() NodeBuilder {
	return createNodeBuilder()
}

// WithName adds a name to the builder
func (app *nodeBuilder) WithName(name string) NodeBuilder {
	app.name = name
	return app
}

// WithContainers add containers to the builder
func (app *nodeBuilder) WithContainers(containers Containers) NodeBuilder {
	app.containers = containers
	return app
}

// Now builds a new Node instance
func (app *nodeBuilder) Now() (Node, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Node instance")
	}

	if app.containers == nil {
		return nil, errors.New("the containers is mandatory in order to build a Node instance")
	}

	return createNode(app.name, app.containers), nil
}
