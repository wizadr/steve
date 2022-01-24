package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

type tokenSpecifierContentBuilder struct {
	token string
	rule  rules.Rule
}

func createTokenSpecifierContentBuilder() TokenSpecifierContentBuilder {
	out := tokenSpecifierContentBuilder{
		token: "",
		rule:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenSpecifierContentBuilder) Create() TokenSpecifierContentBuilder {
	return createTokenSpecifierContentBuilder()
}

// WithRule adds a rule to the builder
func (app *tokenSpecifierContentBuilder) WithRule(rule rules.Rule) TokenSpecifierContentBuilder {
	app.rule = rule
	return app
}

// WithToken adds a token to the builder
func (app *tokenSpecifierContentBuilder) WithToken(token string) TokenSpecifierContentBuilder {
	app.token = token
	return app
}

// Now builds a new TokenSpecifierContent instance
func (app *tokenSpecifierContentBuilder) Now() (TokenSpecifierContent, error) {
	if app.rule != nil {
		return createTokenSpecifierContentWithRule(app.rule), nil
	}

	if app.token != "" {
		return createTokenSpecifierContentWithToken(app.token), nil
	}

	return nil, errors.New("the TokenSpecifierContent instance is invalid")
}
