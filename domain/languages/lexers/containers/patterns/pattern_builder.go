package patterns

import "errors"

type patternBuilder struct {
	name string
	list []Choice
}

func createPatternBuilder() PatternBuilder {
	out := patternBuilder{
		name: "",
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *patternBuilder) Create() PatternBuilder {
	return createPatternBuilder()
}

// WithName adds a name to the builder
func (app *patternBuilder) WithName(name string) PatternBuilder {
	app.name = name
	return app
}

// WithList adds a list to the builder
func (app *patternBuilder) WithList(list []Choice) PatternBuilder {
	app.list = list
	return app
}

// Now builds a new Pattern instance
func (app *patternBuilder) Now() (Pattern, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Pattern instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Choice in order to build a Pattern instance")
	}

	content := createPatternContent(app.list)
	return createPattern(app.name, content), nil
}
