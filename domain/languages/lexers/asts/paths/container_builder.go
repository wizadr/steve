package paths

import "errors"

type containerBuilder struct {
	token     Token
	rule      Rule
	recursive RecursiveToken
	lines     []Line
}

func createContainerBuilder() ContainerBuilder {
	out := containerBuilder{
		token:     nil,
		rule:      nil,
		recursive: nil,
		lines:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *containerBuilder) Create() ContainerBuilder {
	return createContainerBuilder()
}

// WithToken adds a token to the builder
func (app *containerBuilder) WithToken(token Token) ContainerBuilder {
	app.token = token
	return app
}

// WithRule adds a rule to the builder
func (app *containerBuilder) WithRule(rule Rule) ContainerBuilder {
	app.rule = rule
	return app
}

// WithRecursive adds a recursiveToken to the builder
func (app *containerBuilder) WithRecursive(recursive RecursiveToken) ContainerBuilder {
	app.recursive = recursive
	return app
}

// WithLines add lines to the builder
func (app *containerBuilder) WithLines(lines []Line) ContainerBuilder {
	app.lines = lines
	return app
}

// Now builds a new Container instance
func (app *containerBuilder) Now() (Container, error) {
	if app.token != nil {
		return createContainerWithToken(app.token), nil
	}

	if app.rule != nil {
		return createContainerWithRule(app.rule), nil
	}

	if app.recursive != nil {
		return createContainerWithRecursiveToken(app.recursive), nil
	}

	if app.lines != nil && len(app.lines) <= 0 {
		app.lines = nil
	}

	if app.lines != nil {
		return createContainerWithLines(app.lines), nil
	}

	return nil, errors.New("the Container is invalid")
}
