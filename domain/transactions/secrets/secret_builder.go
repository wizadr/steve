package secrets

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transactions"
)

type secretBuilder struct {
	hashAdapter hash.Adapter
	amount      uint
	nonce       string
	public      transactions.Transaction
	origin      Secret
	sides       Secrets
}

func createSecretBuilder(
	hashAdapter hash.Adapter,
) SecretBuilder {
	out := secretBuilder{
		hashAdapter: hashAdapter,
		amount:      0,
		nonce:       "",
		public:      nil,
		origin:      nil,
		sides:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *secretBuilder) Create() SecretBuilder {
	return createSecretBuilder(app.hashAdapter)
}

// WithAmount adds an amount to the builder
func (app *secretBuilder) WithAmount(amount uint) SecretBuilder {
	app.amount = amount
	return app
}

// WithNonce adds a nonce to the builder
func (app *secretBuilder) WithNonce(nonce string) SecretBuilder {
	app.nonce = nonce
	return app
}

// WithPublic adds a public to the builder
func (app *secretBuilder) WithPublic(public transactions.Transaction) SecretBuilder {
	app.public = public
	return app
}

// WithOrigin adds an origin to the builder
func (app *secretBuilder) WithOrigin(origin Secret) SecretBuilder {
	app.origin = origin
	return app
}

// WithSides add sides to the builder
func (app *secretBuilder) WithSides(sides Secrets) SecretBuilder {
	app.sides = sides
	return app
}

// Now builds a new Secret instance
func (app *secretBuilder) Now() (Secret, error) {
	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a Secret instance")
	}

	if app.nonce == "" {
		return nil, errors.New("the nonce is mandatory in order to build a Secret instance")
	}

	if app.public == nil {
		return nil, errors.New("the public transaction is mandatory in order to build a Secret instance")
	}

	data := [][]byte{
		app.public.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.amount))),
		[]byte(app.nonce),
	}

	if app.origin != nil {
		data = append(data, app.origin.Hash().Bytes())
	}

	if app.sides != nil {
		data = append(data, app.sides.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.origin != nil && app.sides != nil {
		return createSecretWithOriginAndSides(*hash, app.amount, app.nonce, app.public, app.origin, app.sides), nil
	}

	if app.origin != nil {
		return createSecretWithOrigin(*hash, app.amount, app.nonce, app.public, app.origin), nil
	}

	if app.sides != nil {
		return createSecretWithSides(*hash, app.amount, app.nonce, app.public, app.sides), nil
	}

	return createSecret(*hash, app.amount, app.nonce, app.public), nil
}
