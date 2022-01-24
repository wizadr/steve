package references

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/references/links"
)

type builder struct {
	references   links.Links
	replacements links.Links
}

func createBuilder() Builder {
	out := builder{
		references:   nil,
		replacements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithReferences add references to the builder
func (app *builder) WithReferences(references links.Links) Builder {
	app.references = references
	return app
}

// WithReplacements add replacements to the builder
func (app *builder) WithReplacements(replacements links.Links) Builder {
	app.replacements = replacements
	return app
}

// Now builds a new References instance
func (app *builder) Now() (References, error) {
	if app.references != nil && app.replacements != nil {
		return createReferencesWithReferencesAndReplacements(app.references, app.replacements), nil
	}

	if app.references != nil {
		return createReferencesWithReferences(app.references), nil
	}

	if app.replacements != nil {
		return createReferencesWithReplacements(app.replacements), nil
	}

	return createReferences(), nil
}
