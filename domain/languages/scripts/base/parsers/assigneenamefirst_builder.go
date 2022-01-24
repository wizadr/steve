package parsers

import "errors"

type assigneeNameFirstBuilder struct {
	declaration  VariableDeclaration
	assigneeName AssigneeName
}

func createAssigneeNameFirstBuilder() AssigneeNameFirstBuilder {
	out := assigneeNameFirstBuilder{
		declaration:  nil,
		assigneeName: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assigneeNameFirstBuilder) Create() AssigneeNameFirstBuilder {
	return createAssigneeNameFirstBuilder()
}

// WithDeclaration adds a declaration to the builder
func (app *assigneeNameFirstBuilder) WithDeclaration(declaration VariableDeclaration) AssigneeNameFirstBuilder {
	app.declaration = declaration
	return app
}

// WithAssigneeName adds an assigneeName to the builder
func (app *assigneeNameFirstBuilder) WithAssigneeName(name AssigneeName) AssigneeNameFirstBuilder {
	app.assigneeName = name
	return app
}

// Now builds a new AssigneeName instance
func (app *assigneeNameFirstBuilder) Now() (AssigneeNameFirst, error) {
	if app.declaration != nil {
		return createAssigneeNameFirstWithDeclaration(app.declaration), nil
	}

	if app.assigneeName != nil {
		return createAssigneeNameFirstWithAssigneeName(app.assigneeName), nil
	}

	return nil, errors.New("the AssigneeNameFirst is invalid")
}
