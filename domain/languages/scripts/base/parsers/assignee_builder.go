package parsers

import "errors"

type assigneeBuilder struct {
	first AssigneeNameFirsts
	name  AssigneeNames
}

func createAssigneeBuilder() AssigneeBuilder {
	out := assigneeBuilder{
		first: nil,
		name:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *assigneeBuilder) Create() AssigneeBuilder {
	return createAssigneeBuilder()
}

// WithFirst adds an assigneeNameFirst to the builder
func (app *assigneeBuilder) WithFirst(first AssigneeNameFirsts) AssigneeBuilder {
	app.first = first
	return app
}

// WithName adds an assigneeName to the builder
func (app *assigneeBuilder) WithName(name AssigneeNames) AssigneeBuilder {
	app.name = name
	return app
}

// Now builds a new Assignee instance
func (app *assigneeBuilder) Now() (Assignee, error) {
	if app.first != nil {
		return createAssigneeWithFirst(app.first), nil
	}

	if app.name != nil {
		return createAssignee(app.name), nil
	}

	return nil, errors.New("the Assignee is invalid")
}
