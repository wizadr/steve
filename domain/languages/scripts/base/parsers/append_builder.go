package parsers

import "errors"

type appendBuilder struct {
	list       ListValue
	assignable SingleAssignable
}

func createAppendBuilder() AppendBuilder {
	out := appendBuilder{
		list:       nil,
		assignable: nil,
	}

	return &out
}

// Create initializes the builder
func (app *appendBuilder) Create() AppendBuilder {
	return createAppendBuilder()
}

// WithList adds a list to the builder
func (app *appendBuilder) WithList(list ListValue) AppendBuilder {
	app.list = list
	return app
}

// WithAssignable adds an assignable to the builder
func (app *appendBuilder) WithAssignable(assignable SingleAssignable) AppendBuilder {
	app.assignable = assignable
	return app
}

// Now builds a new Append instance
func (app *appendBuilder) Now() (Append, error) {
	if app.list == nil {
		return nil, errors.New("the list is mandatory in order to build an Append instance")
	}

	if app.assignable == nil {
		return nil, errors.New("the assignable is mandatory in order to build an Append instance")
	}

	return createAppend(app.list, app.assignable), nil
}
