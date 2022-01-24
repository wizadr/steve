package tokens

import (
	"errors"
)

type matchBuilder struct {
	token TokenMatch
	rule  RuleMatch
	skip  SkipMatch
	line  LineMatch
}

func createMatchBuilder() MatchBuilder {
	out := matchBuilder{
		token: nil,
		rule:  nil,
		skip:  nil,
		line:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *matchBuilder) Create() MatchBuilder {
	return createMatchBuilder()
}

// WithToken adds a token match to the builder
func (app *matchBuilder) WithToken(token TokenMatch) MatchBuilder {
	app.token = token
	return app
}

// WithRule adds a rule match to the builder
func (app *matchBuilder) WithRule(rule RuleMatch) MatchBuilder {
	app.rule = rule
	return app
}

// WithSkip adds a skip match to the builder
func (app *matchBuilder) WithSkip(skip SkipMatch) MatchBuilder {
	app.skip = skip
	return app
}

// WithLine adds a line match to the builder
func (app *matchBuilder) WithLine(line LineMatch) MatchBuilder {
	app.line = line
	return app
}

// Now builds a new Match instance
func (app *matchBuilder) Now() (Match, error) {
	if app.token != nil {
		content := createMatchContentWithToken(app.token)
		return createMatch(content), nil
	}

	if app.rule != nil {
		content := createMatchContentWithRule(app.rule)
		return createMatch(content), nil
	}

	if app.skip != nil {
		content := createMatchContentWithSkip(app.skip)
		return createMatch(content), nil
	}

	if app.line != nil {
		content := createMatchContentWithLine(app.line)
		return createMatch(content), nil
	}

	return nil, errors.New("the Match is invalid")
}
