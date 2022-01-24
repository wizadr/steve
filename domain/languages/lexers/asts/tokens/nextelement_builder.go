package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type nextElementBuilder struct {
	constant string
	pattern  patterns.Pattern
}

func createNextElementBuilder() NextElementBuilder {
	out := nextElementBuilder{
		constant: "",
		pattern:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *nextElementBuilder) Create() NextElementBuilder {
	return createNextElementBuilder()
}

// WithConstant adds a constant to the builder
func (app *nextElementBuilder) WithConstant(constant string) NextElementBuilder {
	app.constant = constant
	return app
}

// WithPattern adds a pattern to the builder
func (app *nextElementBuilder) WithPattern(pattern patterns.Pattern) NextElementBuilder {
	app.pattern = pattern
	return app
}

// Now builds a new NextElement instance
func (app *nextElementBuilder) Now() (NextElement, error) {
	if app.constant != "" {
		return createNextElementWithConstant(app.constant), nil
	}

	if app.pattern != nil {
		return createNextElementWithPattern(app.pattern), nil
	}

	return nil, errors.New("the NextElement is invalid")
}
