package identities

import (
	"errors"

	"github.com/steve-care-software/digital-diamonds/domain/transactions/privates"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

type builder struct {
	name     string
	outgoing privates.Privates
	incoming secrets.Secrets
}

func createBuilder() Builder {
	out := builder{
		name:     "",
		outgoing: nil,
		incoming: nil,
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

// WithOutgoing adds an outgoing to the builder
func (app *builder) WithOutgoing(outgoing privates.Privates) Builder {
	app.outgoing = outgoing
	return app
}

// WithIncoming adds an incoming to the builder
func (app *builder) WithIncoming(incoming secrets.Secrets) Builder {
	app.incoming = incoming
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.outgoing != nil && app.incoming != nil {
		return createIdentityWithOutgoingAndIncoming(app.name, app.outgoing, app.incoming), nil
	}

	if app.outgoing != nil {
		return createIdentityWithOutgoing(app.name, app.outgoing), nil
	}

	if app.incoming != nil {
		return createIdentityWithIncoming(app.name, app.incoming), nil
	}

	return createIdentity(app.name), nil
}
