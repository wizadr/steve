package parsers

import "errors"

type operationElementBuilder struct {
	boolValue       *bool
	variableName    *string
	computableValue Value
	operation       Operation
	fnCall          OperationFuncCall
}

func createOperationElementBuilder() OperationElementBuilder {
	out := operationElementBuilder{
		boolValue:       nil,
		variableName:    nil,
		computableValue: nil,
		operation:       nil,
		fnCall:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *operationElementBuilder) Create() OperationElementBuilder {
	return createOperationElementBuilder()
}

// WithBoolValue adds a boolValue to the builder
func (app *operationElementBuilder) WithBoolValue(boolValue bool) OperationElementBuilder {
	app.boolValue = &boolValue
	return app
}

// WithVariableName adds a variableName to the builder
func (app *operationElementBuilder) WithVariableName(variableName string) OperationElementBuilder {
	app.variableName = &variableName
	return app
}

// WithComputableValue adds a computableValue to the builder
func (app *operationElementBuilder) WithComputableValue(computableValue Value) OperationElementBuilder {
	app.computableValue = computableValue
	return app
}

// WithOperation adds an operation to the builder
func (app *operationElementBuilder) WithOperation(operation Operation) OperationElementBuilder {
	app.operation = operation
	return app
}

// WithFuncCall adds a funcCall to the builder
func (app *operationElementBuilder) WithFuncCall(funcCall OperationFuncCall) OperationElementBuilder {
	app.fnCall = funcCall
	return app
}

// Now builds a new OperationElement instance
func (app *operationElementBuilder) Now() (OperationElement, error) {
	if app.boolValue != nil {
		return createOperationElementWithBoolValue(app.boolValue), nil
	}

	if app.variableName != nil {
		return createOperationElementWithVariableName(app.variableName), nil
	}

	if app.computableValue != nil {
		return createOperationElementWithComputableValue(app.computableValue), nil
	}

	if app.fnCall != nil {
		return createOperationElementWithFuncCall(app.fnCall), nil
	}

	if app.operation != nil {
		return createOperationElementWithOperation(app.operation), nil
	}

	return nil, errors.New("the OperationElement is invalid")
}
