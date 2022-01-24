package parsers

import "errors"

type singleAssignableBuilder struct {
	declarableValue  Value
	operation        Operation
	extract          TotalListAssignable
	listFetchElement ListFetchElement
}

func createSingleAssignableBuilder() SingleAssignableBuilder {
	out := singleAssignableBuilder{
		declarableValue:  nil,
		operation:        nil,
		extract:          nil,
		listFetchElement: nil,
	}

	return &out
}

// Create initializes the builder
func (app *singleAssignableBuilder) Create() SingleAssignableBuilder {
	return createSingleAssignableBuilder()
}

// WithDeclarableValue adds a declarableValue to the builder
func (app *singleAssignableBuilder) WithDeclarableValue(declarableValue Value) SingleAssignableBuilder {
	app.declarableValue = declarableValue
	return app
}

// WithOperation adds an operation to the builder
func (app *singleAssignableBuilder) WithOperation(operation Operation) SingleAssignableBuilder {
	app.operation = operation
	return app
}

// WithExtract adds an extract to the builder
func (app *singleAssignableBuilder) WithExtract(extract TotalListAssignable) SingleAssignableBuilder {
	app.extract = extract
	return app
}

// WithListFetchElement adds a listFetchElement instance
func (app *singleAssignableBuilder) WithListFetchElement(listFetchElement ListFetchElement) SingleAssignableBuilder {
	app.listFetchElement = listFetchElement
	return app
}

// Now builds a new SingleAssignable instance
func (app *singleAssignableBuilder) Now() (SingleAssignable, error) {
	if app.declarableValue != nil {
		return createSingleAssignableWithDeclarableValue(app.declarableValue), nil
	}

	if app.operation != nil {
		return createSingleAssignableWithOperation(app.operation), nil
	}

	if app.extract != nil {
		return createSingleAssignableWithExtract(app.extract), nil
	}

	if app.listFetchElement != nil {
		return createSingleAssignableWithListFetchElement(app.listFetchElement), nil
	}

	return nil, errors.New("the SingleAssignable instance is invalid")
}
