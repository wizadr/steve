package roots

import "errors"

type builder struct {
	content string
	name    string
	nodes   Nodes
}

func createBuilder() Builder {
	out := builder{
		content: "",
		name:    "",
		nodes:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content string) Builder {
	app.content = content
	return app
}

// WithName adds a token name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithNodes adds a nodes to the builder
func (app *builder) WithNodes(nodes Nodes) Builder {
	app.nodes = nodes
	return app
}

// Now builds a new Root instance
func (app *builder) Now() (Root, error) {
	if app.content == "" {
		return nil, errors.New("the content is mandatory in order to build a Root instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Root instance")
	}

	if app.nodes == nil {
		return nil, errors.New("the nodes is mandatory in order to build a Root instance")
	}

	return createRoot(app.content, app.name, app.nodes), nil
}
