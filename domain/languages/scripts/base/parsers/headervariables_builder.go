package parsers

import "errors"

type headerVariablesBuilder struct {
	list []HeaderVariable
}

func createHeaderVariablesBuilder() HeaderVariablesBuilder {
	out := headerVariablesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *headerVariablesBuilder) Create() HeaderVariablesBuilder {
	return createHeaderVariablesBuilder()
}

// WithHeaderVariables add headerVariables to the builder
func (app *headerVariablesBuilder) WithVariables(headerVariables []HeaderVariable) HeaderVariablesBuilder {
	app.list = headerVariables
	return app
}

// Now builds a new HeaderVariables instance
func (app *headerVariablesBuilder) Now() (HeaderVariables, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 HeaderVariable in order to build an HeaderVariables instance")
	}

	return createHeaderVariables(app.list), nil
}
