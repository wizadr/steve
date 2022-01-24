package transactions

import (
	"errors"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Transaction
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Transaction) Builder {
	app.list = list
	return app
}

// Now builds a new Transactions instance
func (app *builder) Now() (Transactions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Transaction in order to build an Transactions instance")
	}

	data := [][]byte{}
	for _, oneTransaction := range app.list {
		data = append(data, oneTransaction.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createTransactions(*hash, app.list), nil
}
