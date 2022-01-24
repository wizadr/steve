package paths

import "errors"

type builder struct {
	rules    string
	tokens   string
	channels string
}

func createBuilder() Builder {
	out := builder{
		rules:    "",
		tokens:   "",
		channels: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
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

// WithChannels add channels to the builder
func (app *builder) WithChannels(channels string) Builder {
	app.channels = channels
	return app
}

// Now builds a new Paths instance
func (app *builder) Now() (Paths, error) {
	if app.rules == "" {
		return nil, errors.New("the rules path is mandatory in order to build a Paths instance")
	}

	if app.tokens == "" {
		return nil, errors.New("the tokens path is mandatory in order to build a Paths instance")
	}

	if app.channels != "" {
		return createPathsWithChanels(app.rules, app.tokens, app.channels), nil
	}

	return createPaths(app.rules, app.tokens), nil
}
