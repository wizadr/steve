package patterns

import "errors"

type groupBuilder struct {
	name string
	list []string
}

func createGroupBuilder() GroupBuilder {
	out := groupBuilder{
		name: "",
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *groupBuilder) Create() GroupBuilder {
	return createGroupBuilder()
}

// WithName adds a name to the builder
func (app *groupBuilder) WithName(name string) GroupBuilder {
	app.name = name
	return app
}

// WithList adds a list to the builder
func (app *groupBuilder) WithList(list []string) GroupBuilder {
	app.list = list
	return app
}

// Now builds a new Group instance
func (app *groupBuilder) Now() (Group, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Group instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 element in the list in order to build a Group instance")
	}

	content := createGroupContent(app.list)
	return createGroup(app.name, content), nil
}
