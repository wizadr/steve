package genesis

import (
	"errors"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type unitBuilder struct {
	hashAdapter hash.Adapter
	content     UnitContent
	sig         signature.Signature
}

func createUnitBuilder(
	hashAdapter hash.Adapter,
) UnitBuilder {
	out := unitBuilder{
		hashAdapter: hashAdapter,
		content:     nil,
		sig:         nil,
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

// WithSignature adds a signature to the builder
func (app *unitBuilder) WithSignature(sig signature.Signature) UnitBuilder {
	app.sig = sig
	return app
}

// Now builds a new Unit instance
func (app *unitBuilder) Now() (Unit, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Unit instance")
	}

	if app.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a Unit instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		[]byte(app.sig.String()),
	})

	if err != nil {
		return nil, err
	}

	return createUnit(*hash, app.content, app.sig), nil
}
