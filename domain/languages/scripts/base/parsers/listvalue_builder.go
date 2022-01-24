package parsers

import "errors"

type listValueBuilder struct {
	assignables  Assignables
	variableName string
}

func createListValueBuilder() ListValueBuilder {
	out := listValueBuilder{
		assignables:  nil,
		variableName: "",
	}

	return &out
}

// Create initializes the builder
func (app *listValueBuilder) Create() ListValueBuilder {
	return createListValueBuilder()
}

// WithAssignables add assignables to the builder
func (app *listValueBuilder) WithAssignables(assignables Assignables) ListValueBuilder {
	app.assignables = assignables
	return app
}

// WithVariableName adds a variableName to the builder
func (app *listValueBuilder) WithVariableName(variableName string) ListValueBuilder {
	app.variableName = variableName
	return app
}

// Now builds a new ListValue instance
func (app *listValueBuilder) Now() (ListValue, error) {
	if app.assignables != nil {
		return createListValueWithAssignables(app.assignables), nil
	}

	if app.variableName != "" {
		return createListValueWithVariableName(app.variableName), nil
	}

	return nil, errors.New("the ListValue is invalid")
}
