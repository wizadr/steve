package containers

import "errors"

type builder struct {
	patterns string
	rules    string
	tokens   string
}

func createBuilder() Builder {
	out := builder{
		patterns: "",
		rules:    "",
		tokens:   "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithPatterns add patterns to the builder
func (app *builder) WithPatterns(patterns string) Builder {
	app.patterns = patterns
	return app
}

// WithRules add rules to the builder
func (app *builder) WithRules(rules string) Builder {
	app.rules = rules
	return app
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens string) Builder {
	app.tokens = tokens
	return app
}

// Now builds a new Container instance
func (app *builder) Now() (Container, error) {
	if app.patterns == "" {
		return nil, errors.New("the patterns is mandatory in order to build a Container instance")
	}

	if app.rules == "" {
		return nil, errors.New("the rules is mandatory in order to build a Container instance")
	}

	if app.tokens == "" {
		return nil, errors.New("the tokens is mandatory in order to build a Container instance")
	}

	return createContainer(app.patterns, app.rules, app.tokens), nil
}
