package suites

import "errors"

type lineBuilder struct {
	index     int
	isSuccess bool
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		index:     -1,
		isSuccess: false,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithIndex adds an index to the builder
func (app *lineBuilder) WithIndex(index uint) LineBuilder {
	app.index = int(index)
	return app
}

// IsSuccessful flags the builder as successful
func (app *lineBuilder) IsSuccessful() LineBuilder {
	app.isSuccess = true
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.index < 0 {
		return nil, errors.New("the index is mandatory in order to build a Line instance")
	}

	return createLine(uint(app.index), app.isSuccess), nil
}
