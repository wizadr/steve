package suites

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

type applicationBuilder struct {
	builder            Builder
	lineBuilder        LineBuilder
	astAdapterBuilder  asts.AdapterBuilder
	tokens             tokens.Tokens
	channels           tokens.Tokens
	tokenNotFoundFn    paths.FetchTokenNotFoundFn
	tokenReplacementFn paths.FetchElementReplacementFn
}

func createApplicationBuilder(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
) ApplicationBuilder {
	out := applicationBuilder{
		builder:            builder,
		lineBuilder:        lineBuilder,
		astAdapterBuilder:  astAdapterBuilder,
		tokens:             nil,
		channels:           nil,
		tokenNotFoundFn:    nil,
		tokenReplacementFn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() ApplicationBuilder {
	return createApplicationBuilder(
		app.builder,
		app.lineBuilder,
		app.astAdapterBuilder,
	)
}

// WithTokens adds a tokens to the builder
func (app *applicationBuilder) WithTokens(tokens tokens.Tokens) ApplicationBuilder {
	app.tokens = tokens
	return app
}

// WithChannels adds a channels to the builder
func (app *applicationBuilder) WithChannels(channels tokens.Tokens) ApplicationBuilder {
	app.channels = channels
	return app
}

// WithTokenNotFoundFunc adds a tokenNotFoundFunc to the builder
func (app *applicationBuilder) WithTokenNotFoundFunc(tokenNotFoundFn paths.FetchTokenNotFoundFn) ApplicationBuilder {
	app.tokenNotFoundFn = tokenNotFoundFn
	return app
}

// WithTokenReplacementFunc adds a tokenReplacementFunc to the builder
func (app *applicationBuilder) WithTokenReplacementFunc(tokenReplacementFn paths.FetchElementReplacementFn) ApplicationBuilder {
	app.tokenReplacementFn = tokenReplacementFn
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (Application, error) {
	if app.tokens == nil {
		return nil, errors.New("the Tokens are mandatory in order to build an Application instance")
	}

	if app.channels != nil && app.tokenNotFoundFn != nil && app.tokenReplacementFn != nil {
		return createApplicationWithChannelsAndTokenNotFoundAndReplacementFunc(
			app.builder,
			app.lineBuilder,
			app.astAdapterBuilder,
			app.tokens,
			app.channels,
			app.tokenNotFoundFn,
			app.tokenReplacementFn,
		), nil
	}

	if app.channels != nil && app.tokenNotFoundFn != nil {
		return createApplicationWithChannelsAndTokenNotFoundFunc(
			app.builder,
			app.lineBuilder,
			app.astAdapterBuilder,
			app.tokens,
			app.channels,
			app.tokenNotFoundFn,
		), nil
	}

	if app.channels != nil && app.tokenReplacementFn != nil {
		return createApplicationWithChannelsAndTokenReplacementFunc(
			app.builder,
			app.lineBuilder,
			app.astAdapterBuilder,
			app.tokens,
			app.channels,
			app.tokenReplacementFn,
		), nil
	}

	if app.tokenNotFoundFn != nil && app.tokenReplacementFn != nil {
		return createApplicationWithTokenNotFoundAndReplacementFunc(
			app.builder,
			app.lineBuilder,
			app.astAdapterBuilder,
			app.tokens,
			app.tokenNotFoundFn,
			app.tokenReplacementFn,
		), nil
	}

	if app.channels != nil {
		return createApplicationWithChannels(
			app.builder,
			app.lineBuilder,
			app.astAdapterBuilder,
			app.tokens,
			app.channels,
		), nil
	}

	if app.tokenReplacementFn != nil {
		return createApplicationWithTokenReplacementFunc(
			app.builder,
			app.lineBuilder,
			app.astAdapterBuilder,
			app.tokens,
			app.tokenReplacementFn,
		), nil
	}

	if app.tokenNotFoundFn != nil {
		return createApplicationWithTokenNotFoundFunc(
			app.builder,
			app.lineBuilder,
			app.astAdapterBuilder,
			app.tokens,
			app.tokenNotFoundFn,
		), nil
	}

	return createApplication(
		app.builder,
		app.lineBuilder,
		app.astAdapterBuilder,
		app.tokens,
	), nil
}
