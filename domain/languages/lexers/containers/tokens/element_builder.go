package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

type elementBuilder struct {
	content     Content
	code        string
	cardinality cardinality.Cardinality
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		content:     nil,
		code:        "",
		cardinality: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithContent adds content to the builder
func (app *elementBuilder) WithContent(content Content) ElementBuilder {
	app.content = content
	return app
}

// WithCode adds code to the builder
func (app *elementBuilder) WithCode(code string) ElementBuilder {
	app.code = code
	return app
}

// WithCardinality add cardinality to the builder
func (app *elementBuilder) WithCardinality(cardinality cardinality.Cardinality) ElementBuilder {
	app.cardinality = cardinality
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build an Element instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build an Element instance")
	}

	if app.cardinality != nil {
		return createElementWithCardinality(app.content, app.code, app.cardinality), nil
	}

	return createElement(app.content, app.code), nil
}
