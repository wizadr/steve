package suites

import "errors"

type builder struct {
	name  string
	lines []Line
}

func createBuilder() Builder {
	out := builder{
		name:  "",
		lines: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithLines add lines to the builder
func (app *builder) WithLines(lines []Line) Builder {
	app.lines = lines
	return app
}

// Now builds a new Suite instance
func (app *builder) Now() (Suite, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.lines != nil && len(app.lines) <= 0 {
		app.lines = nil
	}

	if app.lines == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Suite instance")
	}

	return createSuite(app.name, app.lines), nil
}
