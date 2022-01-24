package genesis

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type unitsBuilder struct {
	hashAdapter hash.Adapter
	list        []Unit
}

func createUnitsBuilder(
	hashAdapter hash.Adapter,
) UnitsBuilder {
	out := unitsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *unitsBuilder) Create() UnitsBuilder {
	return createUnitsBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *unitsBuilder) WithList(list []Unit) UnitsBuilder {
	app.list = list
	return app
}

// Now builds a new Units instance
func (app *unitsBuilder) Now() (Units, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Unit in order to build a Units instance")
	}

	data := [][]byte{}
	for _, oneElement := range app.list {
		data = append(data, oneElement.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createUnits(*hash, app.list), nil
}
