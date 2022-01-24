package parsers

import "errors"

type sliceBuilder struct {
	list      ListValue
	delimiter SliceDelimiter
}

func createSliceBuilder() SliceBuilder {
	out := sliceBuilder{
		list:      nil,
		delimiter: nil,
	}

	return &out
}

// Create initializes the builder
func (app *sliceBuilder) Create() SliceBuilder {
	return createSliceBuilder()
}

// WithList adds a list to the builder
func (app *sliceBuilder) WithList(list ListValue) SliceBuilder {
	app.list = list
	return app
}

// WithDelimiter adds a delimiter to the builder
func (app *sliceBuilder) WithDelimiter(delimiter SliceDelimiter) SliceBuilder {
	app.delimiter = delimiter
	return app
}

// Now builds a new Slice instance
func (app *sliceBuilder) Now() (Slice, error) {
	if app.list == nil {
		return nil, errors.New("the list is mandatory in order to build a Slice instance")
	}

	if app.delimiter == nil {
		return nil, errors.New("the delimiter is mandatory in order to build a Slice instance")
	}

	return createSlice(app.list, app.delimiter), nil
}
