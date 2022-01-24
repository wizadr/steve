package transactions

import (
	"errors"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type transactionBuilder struct {
	hashAdapter hash.Adapter
	content     Content
	auth        signature.RingSignature
}

func createTranactionBuilder(
	hashAdapter hash.Adapter,
) TransactionBuilder {
	out := transactionBuilder{
		hashAdapter: hashAdapter,
		content:     nil,
		auth:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionBuilder) Create() TransactionBuilder {
	return createTranactionBuilder(app.hashAdapter)
}

// WithContent adds a content to the builder
func (app *transactionBuilder) WithContent(content Content) TransactionBuilder {
	app.content = content
	return app
}

// WithAuthorization adds an authorization ring signature to the builder
func (app *transactionBuilder) WithAuthorization(auth signature.RingSignature) TransactionBuilder {
	app.auth = auth
	return app
}

// Now builds a new Transaction instance
func (app *transactionBuilder) Now() (Transaction, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Transaction instance")
	}

	if app.auth == nil {
		return nil, errors.New("the authorization ring signature is mandatory in order to build a Transaction instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		[]byte(app.auth.String()),
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*hash, app.content, app.auth), nil
}
