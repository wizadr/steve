package privates

import (
	"errors"

	"github.com/steve-care-software/steve/domain/cryptography/keys/signature"
	"github.com/steve-care-software/steve/domain/genesis"
	"github.com/steve-care-software/steve/domain/hash"
)

type privateBuilder struct {
	hashAdapter hash.Adapter
	unit        genesis.Unit
	pk          signature.PrivateKey
}

func createPrivateBuilder(
	hashAdapter hash.Adapter,
) PrivateBuilder {
	out := privateBuilder{
		hashAdapter: hashAdapter,
		unit:        nil,
		pk:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *privateBuilder) Create() PrivateBuilder {
	return createPrivateBuilder(app.hashAdapter)
}

// WithUnit adds a unit to the builder
func (app *privateBuilder) WithUnit(unit genesis.Unit) PrivateBuilder {
	app.unit = unit
	return app
}

// WithPrivateKey adds a privateKey to the builder
func (app *privateBuilder) WithPrivateKey(pk signature.PrivateKey) PrivateBuilder {
	app.pk = pk
	return app
}

// Now builds a new Private instance
func (app *privateBuilder) Now() (Private, error) {
	if app.unit == nil {
		return nil, errors.New("the genesis unit is mandatory in order to build a Private instance")
	}

	if app.pk == nil {
		return nil, errors.New("the private key is mandatory in order to build a Private instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.unit.Hash().Bytes(),
		[]byte(app.pk.String()),
	})
	if err != nil {
		return nil, err
	}

	return createPrivate(*hash, app.unit, app.pk), nil
}
