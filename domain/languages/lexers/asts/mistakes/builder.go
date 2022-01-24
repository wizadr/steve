package mistakes

import "errors"

type builder struct {
	idx     Index
	path    Path
	content Content
}

func createBuilder() Builder {
	out := builder{
		idx:     nil,
		path:    nil,
		content: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index Index) Builder {
	app.idx = index
	return app
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path Path) Builder {
	app.path = path
	return app
}

// WithContent adds content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// Now builds a new Mistake instance
func (app *builder) Now() (Mistake, error) {
	if app.idx == nil {
		return nil, errors.New("the Index is mandatory in order to build a Mistake instance")
	}

	if app.path == nil {
		return nil, errors.New("the Path is mandatory in order to build a Mistake instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Mistake instance")
	}

	return createMistake(app.idx, app.path, app.content), nil
}
