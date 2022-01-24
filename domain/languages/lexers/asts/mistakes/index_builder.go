package mistakes

import "errors"

type indexBuilder struct {
	idx    int
	line   int
	column int
}

func createIndexBuilder() IndexBuilder {
	out := indexBuilder{
		idx:    -1,
		line:   -1,
		column: -1,
	}

	return &out
}

// Create initializes the builder
func (app *indexBuilder) Create() IndexBuilder {
	return createIndexBuilder()
}

// WithIndex adds an index to the builder
func (app *indexBuilder) WithIndex(index uint) IndexBuilder {
	app.idx = int(index)
	return app
}

// WithLine adds a line to the builder
func (app *indexBuilder) WithLine(line uint) IndexBuilder {
	app.line = int(line)
	return app
}

// WithColumn adds a column to the builder
func (app *indexBuilder) WithColumn(column uint) IndexBuilder {
	app.column = int(column)
	return app
}

// Now builds a new Index instance
func (app *indexBuilder) Now() (Index, error) {
	if app.idx < 0 {
		return nil, errors.New("the index is mandatory in order to build an Index instance")
	}

	if app.line < 0 {
		return nil, errors.New("the line is mandatory in order to build an Index instance")
	}

	if app.column < 0 {
		return nil, errors.New("the column is mandatory in order to build an Index instance")
	}

	return createIndex(uint(app.idx), uint(app.line), uint(app.column)), nil
}
