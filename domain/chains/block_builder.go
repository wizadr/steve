package chains

import (
	"errors"
	"strconv"
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/owners"
	"github.com/steve-care-software/digital-diamonds/domain/transactions"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type blockBuilder struct {
	hashAdapter  hash.Adapter
	height       uint
	fees         owners.Owners
	transactions transactions.Transactions
	createdOn    *time.Time
	previous     *hash.Hash
}

func createBlockBuilder(
	hashAdapter hash.Adapter,
) BlockBuilder {
	out := blockBuilder{
		hashAdapter:  hashAdapter,
		height:       0,
		fees:         nil,
		transactions: nil,
		createdOn:    nil,
		previous:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder(app.hashAdapter)
}

// WithHeight adds an height to the builder
func (app *blockBuilder) WithHeight(height uint) BlockBuilder {
	app.height = height
	return app
}

// WithFees returns the fees owners
func (app *blockBuilder) WithFees(fees owners.Owners) BlockBuilder {
	app.fees = fees
	return app
}

// WithTransactions adds a transaction hashtree to the builder
func (app *blockBuilder) WithTransactions(transactions transactions.Transactions) BlockBuilder {
	app.transactions = transactions
	return app
}

// WithPrevious adds a previous block hash to the builder
func (app *blockBuilder) WithPrevious(previous hash.Hash) BlockBuilder {
	app.previous = &previous
	return app
}

// CreatedOn adds a creation time to the builder
func (app *blockBuilder) CreatedOn(createdOn time.Time) BlockBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Block instance
func (app *blockBuilder) Now() (Block, error) {
	if app.height == 0 {
		return nil, errors.New("the height is mandatory in order to build a Block instance")
	}

	if app.fees == nil {
		return nil, errors.New("the fees owners is mandatory in order to build a Block instance")
	}

	if app.transactions == nil {
		return nil, errors.New("the transactions are mandatory in order to build a Block instance")
	}

	if app.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Block instance")
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(app.height))),
		app.fees.Hash().Bytes(),
		app.transactions.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.createdOn.Unix()))),
	}

	if app.previous != nil {
		data = append(data, app.previous.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.previous != nil {
		return createBlockWithPrevious(*hash, app.height, app.fees, app.transactions, *app.createdOn, app.previous), nil
	}

	return createBlock(*hash, app.height, app.fees, app.transactions, *app.createdOn), nil
}
