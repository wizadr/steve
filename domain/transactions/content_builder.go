package transactions

import (
	"errors"
	"strconv"
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/rings"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	owner       rings.Ring
	amount      *hash.Hash
	origin      Origin
	external    *hash.Hash
	createdOn   *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		owner:       nil,
		amount:      nil,
		origin:      nil,
		external:    nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter)
}

// WithOwner adds an owner to the builder
func (app *contentBuilder) WithOwner(owner rings.Ring) ContentBuilder {
	app.owner = owner
	return app
}

// WithAmount adds an amount to the builder
func (app *contentBuilder) WithAmount(amount hash.Hash) ContentBuilder {
	app.amount = &amount
	return app
}

// WithOrigin adds an origin to the builder
func (app *contentBuilder) WithOrigin(origin Origin) ContentBuilder {
	app.origin = origin
	return app
}

// WithExternal adds an external hash to the builder
func (app *contentBuilder) WithExternal(external hash.Hash) ContentBuilder {
	app.external = &external
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build a Content instance")
	}

	if app.amount == nil {
		return nil, errors.New("the amount is mandatory in order to build a Content instance")
	}

	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Content instance")
	}

	if app.external == nil {
		return nil, errors.New("the external hash is mandatory in order to build a Content instance")
	}

	if app.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Content instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.owner.Hash().Bytes(),
		app.amount.Bytes(),
		app.origin.Hash().Bytes(),
		app.external.Bytes(),
		[]byte(strconv.Itoa(int(app.createdOn.UTC().Unix()))),
	})

	if err != nil {
		return nil, err
	}

	return createContent(*hash, app.owner, *app.amount, app.origin, *app.external, *app.createdOn), nil
}
