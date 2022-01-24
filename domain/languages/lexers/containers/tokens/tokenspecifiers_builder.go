package tokens

import "errors"

type tokenSpecifiersBuilder struct {
	list []TokenSpecifier
}

func createTokenSpecifiersBuilder() TokenSpecifiersBuilder {
	out := tokenSpecifiersBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenSpecifiersBuilder) Create() TokenSpecifiersBuilder {
	return createTokenSpecifiersBuilder()
}

// WithTokenSpecifiers add tokenSpecifiers to the builder
func (app *tokenSpecifiersBuilder) WithTokenSpecifiers(tokenSpecifiers []TokenSpecifier) TokenSpecifiersBuilder {
	app.list = tokenSpecifiers
	return app
}

// Now builds a new TokenSpecifier instance
func (app *tokenSpecifiersBuilder) Now() (TokenSpecifiers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 SubElemnt in order to build a TokenSpecifiers instance")
	}

	mp := map[string]TokenSpecifier{}
	for _, oneTokenSpecifier := range app.list {
		name := oneTokenSpecifier.Content().Name()
		mp[name] = oneTokenSpecifier
	}

	return createTokenSpecifiers(app.list, mp), nil
}
