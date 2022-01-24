package peers

import (
	"errors"
)

type builder struct {
	list []Peer
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Peer) Builder {
	app.list = list
	return app
}

// Now builds a new Peers instance
func (app *builder) Now() (Peers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Peer in order to build an Peers instance")
	}

	return createPeers(app.list), nil
}
