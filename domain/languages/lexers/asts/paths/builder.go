package paths

import "errors"

type builder struct {
	element Element
	dep     Dependencies
}

func createBuilder() Builder {
	out := builder{
		element: nil,
		dep:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element Element) Builder {
	app.element = element
	return app
}

// WithDependencies add dependencies to the builder
func (app *builder) WithDependencies(dep Dependencies) Builder {
	app.dep = dep
	return app
}

// Now builds a new Path instance
func (app *builder) Now() (Path, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Path instance")
	}

	if app.dep == nil {
		return nil, errors.New("the dependencies are mandatory in order to build a Path instance")
	}

	return createPath(app.element, app.dep), nil
}
