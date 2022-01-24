package mistakes

import "errors"

type contentBuilder struct {
	isTokenHasNoMatch     bool
	containsValidNotMatch ContainsValidNotMatch
	specifierDoNotMatch   SpecifierDoNotMatch
	containsNextElement   ContainsNextElement
	cardinalityIsInvalid  CardinalityIsInvalid
	containsPrefix        ContainsPrefix
}

func createContentBuilder() ContentBuilder {
	out := contentBuilder{
		isTokenHasNoMatch:     false,
		containsValidNotMatch: nil,
		specifierDoNotMatch:   nil,
		containsNextElement:   nil,
		cardinalityIsInvalid:  nil,
		containsPrefix:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder()
}

// IsTokenHasNoMatch flags the builder as tokenHasNoMatch
func (app *contentBuilder) IsTokenHasNoMatch() ContentBuilder {
	app.isTokenHasNoMatch = true
	return app
}

// WithContainsValidNotMatch adds a containsValidNotMatch mistake to the builder
func (app *contentBuilder) WithContainsValidNotMatch(containsValidNotMatch ContainsValidNotMatch) ContentBuilder {
	app.containsValidNotMatch = containsValidNotMatch
	return app
}

// WithSpecifierDoNotMatch adds a specifierDoNotMatch mistake to the builder
func (app *contentBuilder) WithSpecifierDoNotMatch(specifierDoNotMatch SpecifierDoNotMatch) ContentBuilder {
	app.specifierDoNotMatch = specifierDoNotMatch
	return app
}

// WithContainsNextElement adds a specifierDoNotMatch mistake to the builder
func (app *contentBuilder) WithContainsNextElement(containsNextElement ContainsNextElement) ContentBuilder {
	app.containsNextElement = containsNextElement
	return app
}

// WithCardinalityIsInvalid adds a cardinalityIsInvalid mistake to the builder
func (app *contentBuilder) WithCardinalityIsInvalid(cardinalityIsInvalid CardinalityIsInvalid) ContentBuilder {
	app.cardinalityIsInvalid = cardinalityIsInvalid
	return app
}

// WithContainsPrefix adds a containsPrefix mistake to the builder
func (app *contentBuilder) WithContainsPrefix(containsPrefix ContainsPrefix) ContentBuilder {
	app.containsPrefix = containsPrefix
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.isTokenHasNoMatch {
		return createContentWithTokenHasNoMatch(), nil
	}

	if app.containsValidNotMatch != nil {
		return createContentWithContainsValidNotMatch(app.containsValidNotMatch), nil
	}

	if app.specifierDoNotMatch != nil {
		return createContentWithSpecifierDoNotMatch(app.specifierDoNotMatch), nil
	}

	if app.containsNextElement != nil {
		return createContentWithContainsNextElement(app.containsNextElement), nil
	}

	if app.cardinalityIsInvalid != nil {
		return createContentWithCardinalityIsInvalid(app.cardinalityIsInvalid), nil
	}

	if app.containsPrefix != nil {
		return createContentWithContainsPrefix(app.containsPrefix), nil
	}

	return nil, errors.New("the Content is invalid")
}
