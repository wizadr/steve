package parsers

import "errors"

type resultBuilder struct {
	element interface{}
	list    []interface{}
}

func createResultBuilder() ResultBuilder {
	out := resultBuilder{
		element: nil,
		list:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *resultBuilder) Create() ResultBuilder {
	return createResultBuilder()
}

// WithElement adds an element to the builder
func (app *resultBuilder) WithElement(element interface{}) ResultBuilder {
	app.element = element
	return app
}

// WithList adds a list to the builder
func (app *resultBuilder) WithList(list []interface{}) ResultBuilder {
	app.list = list
	return app
}

// WithList adds a list to the builder
func (app *resultBuilder) Now() (Result, error) {
	if app.element != nil {
		app.list = []interface{}{
			app.element,
		}
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 result (first element or list) in order to build a Result instance")
	}

	return createResult(app.list[0], app.list), nil
}
