package parsers

import (
	"errors"
)

type assignmentBuilder struct {
	assignee    Assignee
	assignables Assignables
}

func createAssignmentBuilder() AssignmentBuilder {
	out := assignmentBuilder{
		assignee:    nil,
		assignables: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignmentBuilder) Create() AssignmentBuilder {
	return createAssignmentBuilder()
}

// WithAssignee adds an assignee to the builder
func (app *assignmentBuilder) WithAssignee(assignee Assignee) AssignmentBuilder {
	app.assignee = assignee
	return app
}

// WithAssignables adds an assignables to the builder
func (app *assignmentBuilder) WithAssignables(assignables Assignables) AssignmentBuilder {
	app.assignables = assignables
	return app
}

// Now builds a new Assignment instance
func (app *assignmentBuilder) Now() (Assignment, error) {
	if app.assignee == nil {
		return nil, errors.New("the assignee is mandatory in order to build an Assignment instance")
	}

	if app.assignables == nil {
		return nil, errors.New("the assignables is mandatory in order to build an Assignment instance")
	}

	return createAssignment(app.assignee, app.assignables), nil
}
