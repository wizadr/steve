package owners

import (
	"errors"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/rings"
)

type ownerBuilder struct {
	hashAdapter hash.Adapter
	rings       []rings.Ring
}

func createOwnerBuilder(
	hashAdapter hash.Adapter,
) OwnerBuilder {
	out := ownerBuilder{
		hashAdapter: hashAdapter,
		rings:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *ownerBuilder) Create() OwnerBuilder {
	return createOwnerBuilder(
		app.hashAdapter,
	)
}

// WithRings add rings to the builder
func (app *ownerBuilder) WithRings(rings []rings.Ring) OwnerBuilder {
	app.rings = rings
	return app
}

// Now builds a new Owner instance
func (app *ownerBuilder) Now() (Owner, error) {
	if app.rings != nil && len(app.rings) <= 0 {
		app.rings = nil
	}

	if app.rings == nil {
		return nil, errors.New("there must be at least 1 Ring in order to build an Owner instance")
	}

	data := [][]byte{}
	for _, oneRing := range app.rings {
		data = append(data, oneRing.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createOwner(*hash, app.rings), nil
}
