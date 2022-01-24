package tokens

import "errors"

type lineMatchBuilder struct {
	index   int
	matches Matches
}

func createLineMatchBuilder() LineMatchBuilder {
	out := lineMatchBuilder{
		index:   -1,
		matches: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineMatchBuilder) Create() LineMatchBuilder {
	return createLineMatchBuilder()
}

// WithIndex adds an index to the builder
func (app *lineMatchBuilder) WithIndex(index uint) LineMatchBuilder {
	app.index = int(index)
	return app
}

// WithMatches add matches to the builder
func (app *lineMatchBuilder) WithMatches(matches Matches) LineMatchBuilder {
	app.matches = matches
	return app
}

// Now builds a new LineMatch instance
func (app *lineMatchBuilder) Now() (LineMatch, error) {
	if app.index < 0 {
		return nil, errors.New("the index is mandatory in order to build a LineMatch instance")
	}

	if app.matches == nil {
		return nil, errors.New("the matches is mandatory in order to build a LineMatch instance")
	}

	return createLineMatch(uint(app.index), app.matches), nil
}
