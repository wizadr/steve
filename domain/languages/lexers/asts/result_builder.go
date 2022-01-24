package asts

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type resultBuilder struct {
	input       string
	discoveries patterns.Discoveries
	prefix      string
	suffix      string
}

func createResultBuilder() ResultBuilder {
	out := resultBuilder{
		input:       "",
		discoveries: nil,
		prefix:      "",
		suffix:      "",
	}

	return &out
}

// Create initializes the builder
func (app *resultBuilder) Create() ResultBuilder {
	return createResultBuilder()
}

// WithInput adds an input to the builder
func (app *resultBuilder) WithInput(input string) ResultBuilder {
	app.input = input
	return app
}

// WithDiscoveries add discoveries to the builder
func (app *resultBuilder) WithDiscoveries(discoveries patterns.Discoveries) ResultBuilder {
	app.discoveries = discoveries
	return app
}

// WithPrefix add prefix to the builder
func (app *resultBuilder) WithPrefix(prefix string) ResultBuilder {
	app.prefix = prefix
	return app
}

// WithSuffix add suffix to the builder
func (app *resultBuilder) WithSuffix(suffix string) ResultBuilder {
	app.suffix = suffix
	return app
}

// Now builds a new Result instance
func (app *resultBuilder) Now() (Result, error) {
	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Result instance")
	}

	if app.discoveries == nil {
		return nil, errors.New("the discoveries is mandatory in order to build a Result instance")
	}

	if app.prefix != "" && app.suffix != "" {
		return createResultWithPrefixAndSuffix(app.input, app.discoveries, app.prefix, app.suffix), nil
	}

	if app.prefix != "" {
		return createResultWithPrefix(app.input, app.discoveries, app.prefix), nil
	}

	if app.suffix != "" {
		return createResultWithSuffix(app.input, app.discoveries, app.suffix), nil
	}

	return createResult(app.input, app.discoveries), nil
}
