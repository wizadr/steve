package paths

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

type adapterBuilder struct {
	builder                Builder
	elementBuilder         ElementBuilder
	dependenciesBuilder    DependenciesBuilder
	lineBuilder            LineBuilder
	instructionBuilder     InstructionBuilder
	containerBuilder       ContainerBuilder
	rescursiveTokenBuilder RecursiveTokenBuilder
	tokenBuilder           TokenBuilder
	ruleBuilder            RuleBuilder
	tokenPathBuilder       TokenPathBuilder
	specifierBuilder       SpecifierBuilder
	cardinalityBuilder     cardinality.Builder
	tokens                 tokens.Tokens
	tokenNotFoundFn        FetchTokenNotFoundFn
	tokenReplacementFn     FetchElementReplacementFn
}

func createAdapterBuilder(
	builder Builder,
	elementBuilder ElementBuilder,
	dependenciesBuilder DependenciesBuilder,
	lineBuilder LineBuilder,
	instructionBuilder InstructionBuilder,
	containerBuilder ContainerBuilder,
	rescursiveTokenBuilder RecursiveTokenBuilder,
	tokenBuilder TokenBuilder,
	ruleBuilder RuleBuilder,
	tokenPathBuilder TokenPathBuilder,
	specifierBuilder SpecifierBuilder,
	cardinalityBuilder cardinality.Builder,
) AdapterBuilder {
	out := adapterBuilder{
		builder:                builder,
		elementBuilder:         elementBuilder,
		dependenciesBuilder:    dependenciesBuilder,
		lineBuilder:            lineBuilder,
		instructionBuilder:     instructionBuilder,
		containerBuilder:       containerBuilder,
		rescursiveTokenBuilder: rescursiveTokenBuilder,
		tokenBuilder:           tokenBuilder,
		ruleBuilder:            ruleBuilder,
		tokenPathBuilder:       tokenPathBuilder,
		specifierBuilder:       specifierBuilder,
		cardinalityBuilder:     cardinalityBuilder,
		tokens:                 nil,
		tokenNotFoundFn:        nil,
		tokenReplacementFn:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(
		app.builder,
		app.elementBuilder,
		app.dependenciesBuilder,
		app.lineBuilder,
		app.instructionBuilder,
		app.containerBuilder,
		app.rescursiveTokenBuilder,
		app.tokenBuilder,
		app.ruleBuilder,
		app.tokenPathBuilder,
		app.specifierBuilder,
		app.cardinalityBuilder,
	)
}

// WithTokens add tokens to the builder
func (app *adapterBuilder) WithTokens(tokens tokens.Tokens) AdapterBuilder {
	app.tokens = tokens
	return app
}

// WithTokenNotFoundFunc add a tokenNotFoundFunc to the builder
func (app *adapterBuilder) WithTokenNotFoundFunc(tokenNotFoundFn FetchTokenNotFoundFn) AdapterBuilder {
	app.tokenNotFoundFn = tokenNotFoundFn
	return app
}

// WithTokenReplacementFunc adds an element replacement func to thebuilder
func (app *adapterBuilder) WithTokenReplacementFunc(tokenReplacementFn FetchElementReplacementFn) AdapterBuilder {
	app.tokenReplacementFn = tokenReplacementFn
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.tokens == nil {
		return nil, errors.New("the tokens are mandatory in order to build an Adapter instance")
	}

	if app.tokenNotFoundFn != nil && app.tokenReplacementFn != nil {
		return createAdapterWithTokenNotFoundAndTokenReplaceFunc(
			app.builder,
			app.elementBuilder,
			app.dependenciesBuilder,
			app.lineBuilder,
			app.instructionBuilder,
			app.containerBuilder,
			app.rescursiveTokenBuilder,
			app.tokenBuilder,
			app.ruleBuilder,
			app.tokenPathBuilder,
			app.specifierBuilder,
			app.cardinalityBuilder,
			app.tokens,
			app.tokenNotFoundFn,
			app.tokenReplacementFn,
		), nil
	}

	if app.tokenReplacementFn != nil {
		return createAdapterWithTokenReplaceFunc(
			app.builder,
			app.elementBuilder,
			app.dependenciesBuilder,
			app.lineBuilder,
			app.instructionBuilder,
			app.containerBuilder,
			app.rescursiveTokenBuilder,
			app.tokenBuilder,
			app.ruleBuilder,
			app.tokenPathBuilder,
			app.specifierBuilder,
			app.cardinalityBuilder,
			app.tokens,
			app.tokenReplacementFn,
		), nil
	}

	if app.tokenNotFoundFn != nil {
		return createAdapterWithTokenNotFoundFunc(
			app.builder,
			app.elementBuilder,
			app.dependenciesBuilder,
			app.lineBuilder,
			app.instructionBuilder,
			app.containerBuilder,
			app.rescursiveTokenBuilder,
			app.tokenBuilder,
			app.ruleBuilder,
			app.tokenPathBuilder,
			app.specifierBuilder,
			app.cardinalityBuilder,
			app.tokens,
			app.tokenNotFoundFn,
		), nil
	}

	return createAdapter(
		app.builder,
		app.elementBuilder,
		app.dependenciesBuilder,
		app.lineBuilder,
		app.instructionBuilder,
		app.containerBuilder,
		app.rescursiveTokenBuilder,
		app.tokenBuilder,
		app.ruleBuilder,
		app.tokenPathBuilder,
		app.specifierBuilder,
		app.cardinalityBuilder,
		app.tokens,
	), nil
}
