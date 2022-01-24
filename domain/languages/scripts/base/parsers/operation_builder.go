package parsers

import "errors"

type operationBuilder struct {
	element OperationElement
	suffix  OperationSuffix
}

func createOperationBuilder() OperationBuilder {
	out := operationBuilder{
		element: nil,
		suffix:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder()
}

// WithElement adds an element to the builder
func (app *operationBuilder) WithElement(element OperationElement) OperationBuilder {
	app.element = element
	return app
}

// WithSuffix adds a suffix to the builder
func (app *operationBuilder) WithSuffix(suffix OperationSuffix) OperationBuilder {
	app.suffix = suffix
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.element == nil {
		return nil, errors.New("the operationElement is mandatory in order to build an Operation instance")
	}

	if app.suffix != nil {
		return createOperationWithSuffix(app.element, app.suffix), nil
	}

	return createOperation(app.element).Optimize(), nil
}
