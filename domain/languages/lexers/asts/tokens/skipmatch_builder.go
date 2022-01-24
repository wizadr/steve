package tokens

import "errors"

type skipMatchBuilder struct {
	token TokenMatch
	rule  RuleMatch
}

func createSkipMatchBuilder() SkipMatchBuilder {
	out := skipMatchBuilder{
		token: nil,
		rule:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *skipMatchBuilder) Create() SkipMatchBuilder {
	return createSkipMatchBuilder()
}

// WithToken adds a token to the builder
func (app *skipMatchBuilder) WithToken(token TokenMatch) SkipMatchBuilder {
	app.token = token
	return app
}

// WithRule adds a rule to the builder
func (app *skipMatchBuilder) WithRule(rule RuleMatch) SkipMatchBuilder {
	app.rule = rule
	return app
}

// Now builds a new SkipMatch instance
func (app *skipMatchBuilder) Now() (SkipMatch, error) {
	if app.rule != nil {
		content := createSkipMatchContentWithRule(app.rule)
		return createSkipMatch(content), nil
	}

	if app.token != nil {
		content := createSkipMatchContentWithToken(app.token)
		return createSkipMatch(content), nil
	}

	return nil, errors.New("the SkipMatch is invalid")
}
