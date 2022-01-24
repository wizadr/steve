package patterns

import "errors"

type builder struct {
	list []Pattern
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Pattern) Builder {
	app.list = list
	return app
}

// Now builds a new Patterns instance
func (app *builder) Now() (Patterns, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be atleast 1 Pattern in order to build a Patterns instance")
	}

	mp := map[string]Pattern{}
	for _, onePattern := range app.list {
		name := onePattern.Name()
		mp[name] = onePattern
	}

	return createPatterns(app.list, mp), nil
}
