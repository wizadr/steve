package parsers

import "errors"

type sliceDelimiterBuilder struct {
	index  IndexValue
	amount IndexValue
}

func createSliceDelimiterBuilder() SliceDelimiterBuilder {
	out := sliceDelimiterBuilder{
		index:  nil,
		amount: nil,
	}

	return &out
}

// Create initializes the builder
func (app *sliceDelimiterBuilder) Create() SliceDelimiterBuilder {
	return createSliceDelimiterBuilder()
}

// WithIndex adds an index to the builder
func (app *sliceDelimiterBuilder) WithIndex(index IndexValue) SliceDelimiterBuilder {
	app.index = index
	return app
}

// WithAmount adds an amount to the builder
func (app *sliceDelimiterBuilder) WithAmount(amount IndexValue) SliceDelimiterBuilder {
	app.amount = amount
	return app
}

// Now builds a new SliceDelimiter instance
func (app *sliceDelimiterBuilder) Now() (SliceDelimiter, error) {
	if app.index == nil {
		return nil, errors.New("the index is mandatory in order to build a SliceDelimiter instance")
	}

	if app.amount != nil {
		return createSliceDelimiterWithAmount(app.index, app.amount), nil
	}

	return createSliceDelimiter(app.index), nil
}
