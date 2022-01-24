package receipts

import (
	"errors"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

type builder struct {
	hashAdapter hash.Adapter
	secret      secrets.Secret
	sig         signature.Signature
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		secret:      nil,
		sig:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithSecret adds a secret to the builder
func (app *builder) WithSecret(secret secrets.Secret) Builder {
	app.secret = secret
	return app
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(signature signature.Signature) Builder {
	app.sig = signature
	return app
}

// Now builds a new Receipt instance
func (app *builder) Now() (Receipt, error) {
	if app.secret == nil {
		return nil, errors.New("the secret is mandatory in order to build a Receipt instance")
	}

	if app.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a Receipt instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{})
	if err != nil {
		return nil, err
	}

	return createReceipt(*hash, app.secret, app.sig), nil
}
