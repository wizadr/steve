package privates

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Private
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
func (app *builder) WithList(list []Private) Builder {
	app.list = list
	return app
}

// Now builds a new Privates instance
func (app *builder) Now() (Privates, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Private in order to build an Privates instance")
	}

	data := [][]byte{}
	for _, onePrivate := range app.list {
		data = append(data, onePrivate.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createPrivates(*hash, app.list), nil
}
