package parsers

import "errors"

type assigneeNameBuilder struct {
	isSkip   bool
	variable string
}

func createAssigneeNameBuilder() AssigneeNameBuilder {
	out := assigneeNameBuilder{
		isSkip:   false,
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *assigneeNameBuilder) Create() AssigneeNameBuilder {
	return createAssigneeNameBuilder()
}

// WithVariable adds a variable to the builder
func (app *assigneeNameBuilder) WithVariable(variable string) AssigneeNameBuilder {
	app.variable = variable
	return app
}

// IsSkip flags the builder as skip
func (app *assigneeNameBuilder) IsSkip() AssigneeNameBuilder {
	app.isSkip = true
	return app
}

// Now builds a new AssigneeName instance
func (app *assigneeNameBuilder) Now() (AssigneeName, error) {
	if app.variable != "" {
		return createAssigneeNameWithVariable(app.variable), nil
	}

	if app.isSkip {
		return createAssigneeNameWithSkip(), nil
	}

	return nil, errors.New("the AssigneeName is invalid")
}
