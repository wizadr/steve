package parsers

import "errors"

type assigneeNamesBuilder struct {
	list []AssigneeName
}

func createAssigneeNamesBuilder() AssigneeNamesBuilder {
	out := assigneeNamesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assigneeNamesBuilder) Create() AssigneeNamesBuilder {
	return createAssigneeNamesBuilder()
}

// WithNames add names to the builder
func (app *assigneeNamesBuilder) WithNames(names []AssigneeName) AssigneeNamesBuilder {
	app.list = names
	return app
}

// Now builds a new AssigneeNames instance
func (app *assigneeNamesBuilder) Now() (AssigneeNames, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 AssigneeName in order to build an AssigneeNames instance")
	}

	return createAssigneeNames(app.list), nil
}
