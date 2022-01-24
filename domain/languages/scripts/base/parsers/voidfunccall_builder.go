package parsers

import "errors"

type voidFuncCallBuilder struct {
	log Assignable
}

func createVoidFuncCallBuilder() VoidFuncCallBuilder {
	out := voidFuncCallBuilder{
		log: nil,
	}

	return &out
}

// Create initializes the builder
func (app *voidFuncCallBuilder) Create() VoidFuncCallBuilder {
	return createVoidFuncCallBuilder()
}

// WithLog adds a log to the builder
func (app *voidFuncCallBuilder) WithLog(log Assignable) VoidFuncCallBuilder {
	app.log = log
	return app
}

// Now builds anew VoidFuncCall instance
func (app *voidFuncCallBuilder) Now() (VoidFuncCall, error) {
	if app.log != nil {
		return createVoidFuncCallWithLog(app.log), nil
	}

	return nil, errors.New("the VoidFuncCall is invalid")
}
