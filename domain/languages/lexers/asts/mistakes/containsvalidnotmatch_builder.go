package mistakes

import "errors"

type containsValidNotMatchBuilder struct {
	line int
}

func createContainsValidNotMatchBuilder() ContainsValidNotMatchBuilder {
	out := containsValidNotMatchBuilder{
		line: -1,
	}

	return &out
}

// Create initializes the builder
func (app *containsValidNotMatchBuilder) Create() ContainsValidNotMatchBuilder {
	return createContainsValidNotMatchBuilder()
}

// WithLine adds a line to the builder
func (app *containsValidNotMatchBuilder) WithLine(line uint) ContainsValidNotMatchBuilder {
	app.line = int(line)
	return app
}

// Now builds a new ContainsValidNotMatch instance
func (app *containsValidNotMatchBuilder) Now() (ContainsValidNotMatch, error) {
	if app.line < 0 {
		return nil, errors.New("the line is mandatory in order to build a ContainsValidNotMatch instance")
	}

	return createContainsValidNotMatch(uint(app.line)), nil
}
