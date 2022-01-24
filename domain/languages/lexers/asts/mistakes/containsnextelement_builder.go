package mistakes

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/tokens"
)

type containsNextElementBuilder struct {
	nextElement tokens.NextElement
}

func createContainsNextElementBuilder() ContainsNextElementBuilder {
	out := containsNextElementBuilder{
		nextElement: nil,
	}

	return &out
}

// Create initializes the builder
func (app *containsNextElementBuilder) Create() ContainsNextElementBuilder {
	return createContainsNextElementBuilder()
}

// WithNextElement adds a nextElement to the builder
func (app *containsNextElementBuilder) WithNextElement(nextElement tokens.NextElement) ContainsNextElementBuilder {
	app.nextElement = nextElement
	return app
}

// Now builds a new ContainsNextElement instance
func (app *containsNextElementBuilder) Now() (ContainsNextElement, error) {
	if app.nextElement == nil {
		return nil, errors.New("the nextElement is mandatory in order to build a ContainsNextElement instance")
	}

	return createContainsNextElement(app.nextElement), nil
}
