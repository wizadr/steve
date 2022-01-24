package parsers

import "errors"

type totalListAssignableBuilder struct {
	listValue ListValue
	slice     Slice
	appnd     Append
}

func createTotalListAssignableBuilder() TotalListAssignableBuilder {
	out := totalListAssignableBuilder{
		listValue: nil,
		slice:     nil,
		appnd:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *totalListAssignableBuilder) Create() TotalListAssignableBuilder {
	return createTotalListAssignableBuilder()
}

// WithListValue adds a listValue to the builder
func (app *totalListAssignableBuilder) WithListValue(listValue ListValue) TotalListAssignableBuilder {
	app.listValue = listValue
	return app
}

// WithSlice adds a slice to the builder
func (app *totalListAssignableBuilder) WithSlice(slice Slice) TotalListAssignableBuilder {
	app.slice = slice
	return app
}

// WithAppend adds an append to the builder
func (app *totalListAssignableBuilder) WithAppend(appnd Append) TotalListAssignableBuilder {
	app.appnd = appnd
	return app
}

// Now builds a new TotalListAssignable instance
func (app *totalListAssignableBuilder) Now() (TotalListAssignable, error) {
	if app.listValue != nil {
		return createTotalListAssignableWithListValue(app.listValue), nil
	}

	if app.slice != nil {
		return createTotalListAssignableWithSlice(app.slice), nil
	}

	if app.appnd != nil {
		return createTotalListAssignableWithAppend(app.appnd), nil
	}

	return nil, errors.New("the TotalListAssignable is invalid")
}
