package genesis

import (
	"errors"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/owners"
)

type unitBuilder struct {
	hashAdapter hash.Adapter
	content     UnitContent
	owner       owners.Owner
}

func createUnitBuilder(
	hashAdapter hash.Adapter,
) UnitBuilder {
	out := unitBuilder{
		hashAdapter: hashAdapter,
		content:     nil,
		owner:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *unitBuilder) Create() UnitBuilder {
	return createUnitBuilder(
		app.hashAdapter,
	)
}

// WithContent adds a content to the builder
func (app *unitBuilder) WithContent(content UnitContent) UnitBuilder {
	app.content = content
	return app
}

// WithOwner adds an owner to the builder
func (app *unitBuilder) WithOwner(owner owners.Owner) UnitBuilder {
	app.owner = owner
	return app
}

// Now builds a new Unit instance
func (app *unitBuilder) Now() (Unit, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Unit instance")
	}

	if app.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build a Unit instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		[]byte(app.owner.Hash().Bytes()),
	})

	if err != nil {
		return nil, err
	}

	return createUnit(*hash, app.content, app.owner), nil
}
