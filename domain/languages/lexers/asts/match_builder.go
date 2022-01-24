package asts

import "errors"

type matchBuilder struct {
	token TokenMatch
	rule  RuleMatch
}

func createMatchBuilder() MatchBuilder {
	out := matchBuilder{
		token: nil,
		rule:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *matchBuilder) Create() MatchBuilder {
	return createMatchBuilder()
}

// WithToken adds a token to the builder
func (app *matchBuilder) WithToken(token TokenMatch) MatchBuilder {
	app.token = token
	return app
}

// WithRule adds a rule to the builder
func (app *matchBuilder) WithRule(rule RuleMatch) MatchBuilder {
	app.rule = rule
	return app
}

// Now builds a new Match instance
func (app *matchBuilder) Now() (Match, error) {
	if app.token != nil {
		return createMatchWithToken(app.token), nil
	}

	if app.rule != nil {
		return createMatchWithRule(app.rule), nil
	}

	return nil, errors.New("the Match is invalid")
}
