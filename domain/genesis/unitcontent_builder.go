package genesis

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type unitContentBuilder struct {
	hashAdapter hash.Adapter
	amount      uint
	nonce       string
	activatedOn *time.Time
	createdOn   *time.Time
}

func createUnitContentBuilder(
	hashAdapter hash.Adapter,
) UnitContentBuilder {
	out := unitContentBuilder{
		hashAdapter: hashAdapter,
		amount:      0,
		nonce:       "",
		activatedOn: nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *unitContentBuilder) Create() UnitContentBuilder {
	return createUnitContentBuilder(app.hashAdapter)
}

// WithAmount adds an amount to the builder
func (app *unitContentBuilder) WithAmount(amount uint) UnitContentBuilder {
	app.amount = amount
	return app
}

// WithNonce adds a nonce to the builder
func (app *unitContentBuilder) WithNonce(nonce string) UnitContentBuilder {
	app.nonce = nonce
	return app
}

// ActivatedOn adds an activation time to the builder
func (app *unitContentBuilder) ActivatedOn(activatedOn time.Time) UnitContentBuilder {
	app.activatedOn = &activatedOn
	return app
}

// CreatedOn adds a creation time to the builder
func (app *unitContentBuilder) CreatedOn(createdOn time.Time) UnitContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new UnitContent instance
func (app *unitContentBuilder) Now() (UnitContent, error) {
	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a UnitContent instance")
	}

	if app.nonce == "" {
		return nil, errors.New("the nonce is mandatory in order to build a UnitContent instance")
	}

	if app.activatedOn == nil {
		return nil, errors.New("the activation time is mandatory in order to build a UnitContent instance")
	}

	if app.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a UnitContent instance")
	}

	if app.activatedOn.Before(*app.createdOn) {
		str := fmt.Sprintf("the activation time (%s) cannot be before the creation time (%s)", app.activatedOn.String(), app.createdOn.String())
		return nil, errors.New(str)
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.amount))),
		[]byte(app.nonce),
		[]byte(strconv.Itoa(int(app.activatedOn.UTC().UnixNano()))),
		[]byte(strconv.Itoa(int(app.createdOn.UTC().UnixNano()))),
	})

	if err != nil {
		return nil, err
	}

	return createUnitContent(*hash, app.amount, app.nonce, *app.activatedOn, *app.createdOn), nil
}
