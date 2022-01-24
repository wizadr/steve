package paths

import "errors"

type elementBuilder struct {
	name string
	must []Line
	not  []Line
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		name: "",
		must: nil,
		not:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithName adds a name to the builder
func (app *elementBuilder) WithName(name string) ElementBuilder {
	app.name = name
	return app
}

// WithMust add must lines to the builder
func (app *elementBuilder) WithMust(must []Line) ElementBuilder {
	app.must = must
	return app
}

// WithNot add not lines to the builder
func (app *elementBuilder) WithNot(not []Line) ElementBuilder {
	app.not = not
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Element instance")
	}

	if app.must != nil && len(app.must) <= 0 {
		app.must = nil
	}

	if app.must == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Element instance")
	}

	if app.not != nil && len(app.not) <= 0 {
		app.not = nil
	}

	if app.not != nil {
		return createElementwithNotContainers(app.name, app.must, app.not), nil
	}

	return createElement(app.name, app.must), nil
}
