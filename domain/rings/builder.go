package rings

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	min         uint
	max         uint
	list        []hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
	min uint,
	max uint,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		min:         min,
		max:         max,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.min,
		app.max,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []hash.Hash) Builder {
	app.list = list
	return app
}

// Now builds a new Ring instance
func (app *builder) Now() (Ring, error) {
	if app.list == nil {
		return nil, errors.New("the hashes cannot be nil")
	}

	amount := uint(len(app.list))
	if app.min > amount {
		str := fmt.Sprintf("there must be at least %d hashes in the ring, %d returned", app.min, amount)
		return nil, errors.New(str)
	}

	if app.max < amount {
		str := fmt.Sprintf("there must be a maximum of %d hashes in the ring, %d returned", app.max, amount)
		return nil, errors.New(str)
	}

	data := [][]byte{}
	for _, oneHash := range app.list {
		data = append(data, oneHash.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createRing(*hash, app.list), nil
}
