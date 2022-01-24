package paths

import "errors"

type tokenPathBuilder struct {
	element    Element
	specifiers []Specifier
}

func createTokenPathBuilder() TokenPathBuilder {
	out := tokenPathBuilder{
		element:    nil,
		specifiers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenPathBuilder) Create() TokenPathBuilder {
	return createTokenPathBuilder()
}

// WithName adds a name to the builder
func (app *tokenPathBuilder) WithElement(element Element) TokenPathBuilder {
	app.element = element
	return app
}

// WithSpecifiers add specifiers to the builder
func (app *tokenPathBuilder) WithSpecifiers(specifiers []Specifier) TokenPathBuilder {
	app.specifiers = specifiers
	return app
}

// Now builds a new TokenPath instance
func (app *tokenPathBuilder) Now() (TokenPath, error) {
	if app.element == nil {
		return nil, errors.New("the token element is mandatory in order to build a TokenPath instance")
	}

	if app.specifiers != nil && len(app.specifiers) <= 0 {
		app.specifiers = nil
	}

	if app.specifiers != nil {
		return createTokenPathWithSpecifiers(app.element, app.specifiers), nil
	}

	return createTokenPath(app.element), nil
}
