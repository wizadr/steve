package parsers

import "errors"

type operationSuffixBuilder struct {
	operator *uint16
	element  OperationElement
}

func createOperationSuffixBuilder() OperationSuffixBuilder {
	out := operationSuffixBuilder{
		operator: nil,
		element:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *operationSuffixBuilder) Create() OperationSuffixBuilder {
	return createOperationSuffixBuilder()
}

// WithOperator adds an operator to the builder
func (app *operationSuffixBuilder) WithOperator(operator uint16) OperationSuffixBuilder {
	app.operator = &operator
	return app
}

// WithElement adds an element to the builder
func (app *operationSuffixBuilder) WithElement(element OperationElement) OperationSuffixBuilder {
	app.element = element
	return app
}

// Now builds a new OperationSuffix instance
func (app *operationSuffixBuilder) Now() (OperationSuffix, error) {
	if app.operator == nil {
		return nil, errors.New("the operator is mandatory in order to build an OperationSuffix instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build an OperationSuffix instance")
	}

	return createOperationSuffix(*app.operator, app.element), nil
}
