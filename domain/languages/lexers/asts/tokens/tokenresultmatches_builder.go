package tokens

import "errors"

type tokenResultMatchesBuilder struct {
	list []Token
}

func createTokenResultMatchesBuilder() TokenResultMatchesBuilder {
	out := tokenResultMatchesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenResultMatchesBuilder) Create() TokenResultMatchesBuilder {
	return createTokenResultMatchesBuilder()
}

// WithResults add results to the builder
func (app *tokenResultMatchesBuilder) WithResults(results []Token) TokenResultMatchesBuilder {
	app.list = results
	return app
}

// Now builds a new TokenResultMatches instance
func (app *tokenResultMatchesBuilder) Now() (TokenResultMatches, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at elast 1 Token instance in order to build a TokenResultMatches instance")
	}

	return createTokenResultMatches(app.list), nil
}
