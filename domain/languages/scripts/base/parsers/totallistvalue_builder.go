package parsers

import "errors"

type totalListValueBuilder struct {
	assignable   TotalListAssignable
	variableName string
}

func createTotalListValueBuilder() TotalListValueBuilder {
	out := totalListValueBuilder{
		assignable:   nil,
		variableName: "",
	}

	return &out
}

// Create initializes the builder
func (app *totalListValueBuilder) Create() TotalListValueBuilder {
	return createTotalListValueBuilder()
}

// WithAssignable adds an assignable to the builder
func (app *totalListValueBuilder) WithAssignable(assignable TotalListAssignable) TotalListValueBuilder {
	app.assignable = assignable
	return app
}

// WithVariableName adds a variableName to the builder
func (app *totalListValueBuilder) WithVariableName(variableName string) TotalListValueBuilder {
	app.variableName = variableName
	return app
}

// Now builds a new TotalListValue instance
func (app *totalListValueBuilder) Now() (TotalListValue, error) {
	if app.assignable != nil {
		return createTotalListValueWithAssignable(app.assignable), nil
	}

	if app.variableName != "" {
		return createTotalListValueWithVariableName(app.variableName), nil
	}

	return nil, errors.New("the TotalListValue is invalid")
}
