package roots

import (
	"errors"
)

type nodesBuilder struct {
	list map[string]Node
}

func createNodesBuilder() NodesBuilder {
	out := nodesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *nodesBuilder) Create() NodesBuilder {
	return createNodesBuilder()
}

// WithNodes add nodes to the builder
func (app *nodesBuilder) WithNodes(nodes map[string]Node) NodesBuilder {
	app.list = nodes
	return app
}

// Now builds a new Nodes instance
func (app *nodesBuilder) Now() (Nodes, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Container instance in order to build a Nodes instance")
	}

	return createNodes(app.list), nil
}
