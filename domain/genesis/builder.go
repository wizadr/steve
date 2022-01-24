package genesis

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	units       Units
	fees        *uint
	createdOn   *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		units:       nil,
		fees:        nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithUnits add units to the builder
func (app *builder) WithUnits(units Units) Builder {
	app.units = units
	return app
}

// WithFees add fees to the builder
func (app *builder) WithFees(fees uint) Builder {
	app.fees = &fees
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.units == nil {
		return nil, errors.New("the units is mandatory in order to build a Genesis instance")
	}

	if app.fees == nil {
		return nil, errors.New("the fees is mandatory in order to build a Genesis instance")
	}

	if app.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Genesis instance")
	}

	list := app.units.All()
	for index, oneUnit := range list {
		unitCreatedOn := oneUnit.Content().CreatedOn()
		if unitCreatedOn.Before(*app.createdOn) {
			str := fmt.Sprintf("the unit (index: %d, createdOn: %s) was created before the genesis creation time (%s)", index, unitCreatedOn.String(), app.createdOn.String())
			return nil, errors.New(str)
		}
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(*app.fees))),
		app.units.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.createdOn.UTC().UnixNano()))),
	})

	if err != nil {
		return nil, err
	}

	return createGenesis(*hash, app.units, *app.fees, *app.createdOn), nil
}
