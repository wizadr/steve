package rings

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Ring
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Ring) Builder {
	app.list = list
	return app
}

// Now builds a new Rings instance
func (app *builder) Now() (Rings, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Ring in order to build an Rings instance")
	}

	data := [][]byte{}
	for _, oneRing := range app.list {
		data = append(data, oneRing.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createRings(*hash, app.list), nil
}
