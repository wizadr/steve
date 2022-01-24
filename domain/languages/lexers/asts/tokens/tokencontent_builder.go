package tokens

import "errors"

type tokenContentBuilder struct {
	must   LineMatch
	not    LineMatch
	suffix string
}

func createTokenContentBuilder() TokenContentBuilder {
	out := tokenContentBuilder{
		must:   nil,
		not:    nil,
		suffix: "",
	}

	return &out
}

// Create initializes the builder
func (app *tokenContentBuilder) Create() TokenContentBuilder {
	return createTokenContentBuilder()
}

// WithMust adds a must line match to the builder
func (app *tokenContentBuilder) WithMust(matches LineMatch) TokenContentBuilder {
	app.must = matches
	return app
}

// WithNot adds a not line match to the builder
func (app *tokenContentBuilder) WithNot(not LineMatch) TokenContentBuilder {
	app.not = not
	return app
}

// Now builds a new Token instance
func (app *tokenContentBuilder) Now() (TokenContent, error) {
	if app.must == nil {
		return nil, errors.New("the must LineMatch is mandatory in order to build a Token instance")
	}

	if app.not != nil {
		return createTokenContentWithNot(app.must, app.not), nil
	}

	return createTokenContent(app.must), nil
}
