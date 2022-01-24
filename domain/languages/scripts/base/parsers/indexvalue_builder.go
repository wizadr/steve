package parsers

import "errors"

type indexValueBuilder struct {
	value        Value
	variableName string
}

func createIndexValueBuilder() IndexValueBuilder {
	out := indexValueBuilder{
		value:        nil,
		variableName: "",
	}

	return &out
}

// Create initializes the builder
func (app *indexValueBuilder) Create() IndexValueBuilder {
	return createIndexValueBuilder()
}

// WithValue adds a value to the builder
func (app *indexValueBuilder) WithValue(value Value) IndexValueBuilder {
	app.value = value
	return app
}

// WithVariableName adds a variableName to the builder
func (app *indexValueBuilder) WithVariableName(variableName string) IndexValueBuilder {
	app.variableName = variableName
	return app
}

// Now builds a new IndexValue instance
func (app *indexValueBuilder) Now() (IndexValue, error) {
	if app.value != nil {
		return createIndexValueWithValue(app.value), nil
	}

	if app.variableName != "" {
		return createIndexValueWithVariableName(app.variableName), nil
	}

	return nil, errors.New("the IndexValue is invalid")
}
