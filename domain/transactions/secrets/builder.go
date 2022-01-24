package secrets

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Secret
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
func (app *builder) WithList(list []Secret) Builder {
	app.list = list
	return app
}

// Now builds a new Secrets instance
func (app *builder) Now() (Secrets, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Secret in order to build an Secrets instance")
	}

	data := [][]byte{}
	for _, oneSecret := range app.list {
		data = append(data, oneSecret.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createSecrets(*hash, app.list), nil
}
