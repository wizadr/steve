package tokens

import "errors"

type tokenReferenceBuilder struct {
	name       string
	specifiers TokenSpecifiers
}

func createTokenReferenceBuilder() TokenReferenceBuilder {
	out := tokenReferenceBuilder{
		name:       "",
		specifiers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenReferenceBuilder) Create() TokenReferenceBuilder {
	return createTokenReferenceBuilder()
}

// WithName adds a name to the builder
func (app *tokenReferenceBuilder) WithName(name string) TokenReferenceBuilder {
	app.name = name
	return app
}

// WithSpecifiers adds specifiers to the builder
func (app *tokenReferenceBuilder) WithSpecifiers(specifiers TokenSpecifiers) TokenReferenceBuilder {
	app.specifiers = specifiers
	return app
}

// Now builds a new TokenReference instance
func (app *tokenReferenceBuilder) Now() (TokenReference, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a TokenReference instance")
	}

	if app.specifiers != nil {
		return createTokenReferenceWithSpecifiers(app.name, app.specifiers), nil
	}

	return createTokenReference(app.name), nil
}
