package parsers

import "errors"

type assigneeNameFirstsBuilder struct {
	list []AssigneeNameFirst
}

func createAssigneeNameFirstsBuilder() AssigneeNameFirstsBuilder {
	out := assigneeNameFirstsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assigneeNameFirstsBuilder) Create() AssigneeNameFirstsBuilder {
	return createAssigneeNameFirstsBuilder()
}

// WithNames add names to the builder
func (app *assigneeNameFirstsBuilder) WithNames(names []AssigneeNameFirst) AssigneeNameFirstsBuilder {
	app.list = names
	return app
}

// Now builds a new AssigneeNameFirsts instance
func (app *assigneeNameFirstsBuilder) Now() (AssigneeNameFirsts, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 AssigneeNameFirst in order to build an AssigneeNameFirsts instance")
	}

	return createAssigneeNameFirsts(app.list), nil
}
