package parsers

import "errors"

type operationFuncCallBuilder struct {
	null Operation
}

func createOperationFuncCallBuilder() OperationFuncCallBuilder {
	out := operationFuncCallBuilder{
		null: nil,
	}

	return &out
}

// Create initializes the builder
func (app *operationFuncCallBuilder) Create() OperationFuncCallBuilder {
	return createOperationFuncCallBuilder()
}

// WithNull adds a null operation to the builder
func (app *operationFuncCallBuilder) WithNull(null Operation) OperationFuncCallBuilder {
	app.null = null
	return app
}

// Now builds a new OperationFuncCall instance
func (app *operationFuncCallBuilder) Now() (OperationFuncCall, error) {
	if app.null != nil {
		return createOperationFuncCallWithIsNull(app.null), nil
	}

	return nil, errors.New("the OperationFuncCall is invalid")
}
