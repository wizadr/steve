package links

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/paths"
)

type includeBuilder struct {
	name  string
	paths paths.Paths
}

func createIncludeBuilder() IncludeBuilder {
	out := includeBuilder{
		name:  "",
		paths: nil,
	}

	return &out
}

// Create initializes the builder
func (app *includeBuilder) Create() IncludeBuilder {
	return createIncludeBuilder()
}

// WithName adds a name to the builder
func (app *includeBuilder) WithName(name string) IncludeBuilder {
	app.name = name
	return app
}

// WithPaths adds a paths to the builder
func (app *includeBuilder) WithPaths(paths paths.Paths) IncludeBuilder {
	app.paths = paths
	return app
}

// Now builds a new Include instance
func (app *includeBuilder) Now() (Include, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Include instance")
	}

	if app.paths == nil {
		return nil, errors.New("the Paths instance is mandatory in order to build an Include instance")
	}

	return createInclude(app.name, app.paths), nil
}
