package patterns

import "errors"

type serieBuilder struct {
	name string
	list []Group
}

func createSerieBuilder() SerieBuilder {
	out := serieBuilder{
		name: "",
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *serieBuilder) Create() SerieBuilder {
	return createSerieBuilder()
}

// WithName adds a name to the builder
func (app *serieBuilder) WithName(name string) SerieBuilder {
	app.name = name
	return app
}

// WithList adds a list to the builder
func (app *serieBuilder) WithList(list []Group) SerieBuilder {
	app.list = list
	return app
}

// Now builds a new Serie instance
func (app *serieBuilder) Now() (Serie, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Serie instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Group in the list in order to build a Serie instance")
	}

	content := createSerieContent(app.list)
	return createSerie(app.name, content), nil
}
