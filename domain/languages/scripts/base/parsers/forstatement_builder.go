package parsers

import "errors"

type forStatementBuilder struct {
	index    AssigneeName
	value    AssigneeName
	iterable TotalListValue
}

func createForStatementBuilder() ForStatementBuilder {
	out := forStatementBuilder{
		index:    nil,
		value:    nil,
		iterable: nil,
	}

	return &out
}

// Create intiializes the builder
func (app *forStatementBuilder) Create() ForStatementBuilder {
	return createForStatementBuilder()
}

// WithIndex adds an index to the builder
func (app *forStatementBuilder) WithIndex(index AssigneeName) ForStatementBuilder {
	app.index = index
	return app
}

// WithValue adds a value to the builder
func (app *forStatementBuilder) WithValue(value AssigneeName) ForStatementBuilder {
	app.value = value
	return app
}

// WithIterable adds an iterable to the builder
func (app *forStatementBuilder) WithIterable(iterable TotalListValue) ForStatementBuilder {
	app.iterable = iterable
	return app
}

// Now builds a new ForStatement instance
func (app *forStatementBuilder) Now() (ForStatement, error) {
	if app.index == nil {
		return nil, errors.New("the index is mandatory in order to build a ForStatement instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a ForStatement instance")
	}

	if app.iterable == nil {
		return nil, errors.New("the iterable is mandatory in order to build a ForStatement instance")
	}

	return createForStatement(app.index, app.value, app.iterable), nil
}
