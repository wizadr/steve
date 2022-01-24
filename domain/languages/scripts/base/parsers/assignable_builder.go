package parsers

import "errors"

type assignableBuilder struct {
	list   TotalListAssignable
	single SingleAssignable
}

func createAssignableBuilder() AssignableBuilder {
	out := assignableBuilder{
		list:   nil,
		single: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignableBuilder) Create() AssignableBuilder {
	return createAssignableBuilder()
}

// WithList adds a list to the builder
func (app *assignableBuilder) WithList(list TotalListAssignable) AssignableBuilder {
	app.list = list
	return app
}

// WithSingle adds a single to the builder
func (app *assignableBuilder) WithSingle(single SingleAssignable) AssignableBuilder {
	app.single = single
	return app
}

// Now builds a new Assignable instance
func (app *assignableBuilder) Now() (Assignable, error) {
	if app.list != nil {
		return createAssignableWithList(app.list), nil
	}

	if app.single != nil {
		return createAssignableWithSingle(app.single), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
