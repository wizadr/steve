package results

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

type builder struct {
	input   rules.Rule
	matches []patterns.Result
}

func createBuilder() Builder {
	out := builder{
		input:   nil,
		matches: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input rules.Rule) Builder {
	app.input = input
	return app
}

// WithMatches add matches to the builder
func (app *builder) WithMatches(matches []patterns.Result) Builder {
	app.matches = matches
	return app
}

// Now builds a Result instance
func (app *builder) Now() (Result, error) {
	if app.input == nil {
		return nil, errors.New("the rule input is mandatory in order to build a Result instance")
	}

	if app.matches != nil {
		matches := createMatches(app.matches)
		return createResultWithMatches(app.input, matches), nil
	}

	return createResult(app.input), nil
}
