package parsers

import "errors"

type intIncrementBuilder struct {
	variableName string
	increment    Value
}

func createIntIncrementBuilder() IntIncrementBuilder {
	out := intIncrementBuilder{
		variableName: "",
		increment:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *intIncrementBuilder) Create() IntIncrementBuilder {
	return createIntIncrementBuilder()
}

// WithVariableName adds a variableName to the builder
func (app *intIncrementBuilder) WithVariableName(variableName string) IntIncrementBuilder {
	app.variableName = variableName
	return app
}

// WithIncrement adds an increment to the builder
func (app *intIncrementBuilder) WithIncrement(increment Value) IntIncrementBuilder {
	app.increment = increment
	return app
}

// Now builds a new IntIncrement instance
func (app *intIncrementBuilder) Now() (IntIncrement, error) {
	if app.variableName == "" {
		return nil, errors.New("the variableName is mandatory in order to build an IntIncrement instance")
	}

	if app.increment == nil {
		return nil, errors.New("the increment is mandatory in order to build an IntIncrement instance")
	}

	return createIntIncrement(app.variableName, app.increment), nil
}
