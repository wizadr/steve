package privates

import (
	"errors"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/receipts"
)

type privateBuilder struct {
	hashAdapter hash.Adapter
	pk          signature.PrivateKey
	receipt     receipts.Receipt
}

func createPrivateBuilder(
	hashAdapter hash.Adapter,
) PrivateBuilder {
	out := privateBuilder{
		hashAdapter: hashAdapter,
		pk:          nil,
		receipt:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *privateBuilder) Create() PrivateBuilder {
	return createPrivateBuilder(app.hashAdapter)
}

// WithPrivateKey adds a private key to the builder
func (app *privateBuilder) WithPrivateKey(pk signature.PrivateKey) PrivateBuilder {
	app.pk = pk
	return app
}

// WithReceipt adds a receipt to the builder
func (app *privateBuilder) WithReceipt(receipt receipts.Receipt) PrivateBuilder {
	app.receipt = receipt
	return app
}

// now builds a new Private instance
func (app *privateBuilder) Now() (Private, error) {
	if app.pk == nil {
		return nil, errors.New("the private key is mandatory in order to build a Private instance")
	}

	if app.receipt == nil {
		return nil, errors.New("the receipt is mandatory in order to build a Private instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{})
	if err != nil {
		return nil, err
	}

	return createPrivateKey(*hash, app.pk, app.receipt), nil
}
