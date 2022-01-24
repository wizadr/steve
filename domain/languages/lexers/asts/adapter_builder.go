package asts

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	ast_tokens "github.com/steve-care-software/steve/domain/languages/lexers/asts/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

type adapterBuilder struct {
	discoveriesBuilder      patterns.DiscoveriesBuilder
	rulesAdapter            rules.Adapter
	tokensAdapterBuilder    tokens.AdapterBuilder
	astTokensAdapterBuilder ast_tokens.AdapterBuilder
	mistakeAdapter          mistakes.Adapter
	builder                 Builder
	tokenBuilder            TokenBuilder
	lineMatchBuilder        LineMatchBuilder
	matchBuilder            MatchBuilder
	tokenMatchBuilder       TokenMatchBuilder
	ruleMatchBuilder        RuleMatchBuilder
	resultBuilder           ResultBuilder
	root                    tokens.Token
	tokens                  tokens.Tokens
	channels                tokens.Tokens
	tokenNotFoundFn         paths.FetchTokenNotFoundFn
	tokenReplacementFn      paths.FetchElementReplacementFn
	canContainPrefix        bool
	canContainSuffix        bool
}

func createAdapterBuilder(
	discoveriesBuilder patterns.DiscoveriesBuilder,
	rulesAdapter rules.Adapter,
	tokensAdapterBuilder tokens.AdapterBuilder,
	astTokensAdapterBuilder ast_tokens.AdapterBuilder,
	mistakeAdapter mistakes.Adapter,
	builder Builder,
	tokenBuilder TokenBuilder,
	lineMatchBuilder LineMatchBuilder,
	matchBuilder MatchBuilder,
	tokenMatchBuilder TokenMatchBuilder,
	ruleMatchBuilder RuleMatchBuilder,
	resultBuilder ResultBuilder,
) AdapterBuilder {
	out := adapterBuilder{
		discoveriesBuilder:      discoveriesBuilder,
		rulesAdapter:            rulesAdapter,
		tokensAdapterBuilder:    tokensAdapterBuilder,
		astTokensAdapterBuilder: astTokensAdapterBuilder,
		mistakeAdapter:          mistakeAdapter,
		builder:                 builder,
		tokenBuilder:            tokenBuilder,
		lineMatchBuilder:        lineMatchBuilder,
		matchBuilder:            matchBuilder,
		tokenMatchBuilder:       tokenMatchBuilder,
		ruleMatchBuilder:        ruleMatchBuilder,
		resultBuilder:           resultBuilder,
		root:                    nil,
		tokens:                  nil,
		channels:                nil,
		tokenNotFoundFn:         nil,
		tokenReplacementFn:      nil,
		canContainPrefix:        false,
		canContainSuffix:        false,
	}

	return &out

}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(
		app.discoveriesBuilder,
		app.rulesAdapter,
		app.tokensAdapterBuilder,
		app.astTokensAdapterBuilder,
		app.mistakeAdapter,
		app.builder,
		app.tokenBuilder,
		app.lineMatchBuilder,
		app.matchBuilder,
		app.tokenMatchBuilder,
		app.ruleMatchBuilder,
		app.resultBuilder,
	)
}

// WithRoot adds a root to the builder
func (app *adapterBuilder) WithRoot(root tokens.Token) AdapterBuilder {
	app.root = root
	return app
}

// WithTokens adds a tokens to the builder
func (app *adapterBuilder) WithTokens(tokens tokens.Tokens) AdapterBuilder {
	app.tokens = tokens
	return app
}

// WithChannels adds a channels to the builder
func (app *adapterBuilder) WithChannels(channels tokens.Tokens) AdapterBuilder {
	app.channels = channels
	return app
}

// WithTokenNotFoundFunc adds a tokenNotFoundFunc to the builder
func (app *adapterBuilder) WithTokenNotFoundFunc(tokenNotFoundFn paths.FetchTokenNotFoundFn) AdapterBuilder {
	app.tokenNotFoundFn = tokenNotFoundFn
	return app
}

// WithTokenReplacementFunc adds a tokenReplacementFunc to the builder
func (app *adapterBuilder) WithTokenReplacementFunc(tokenReplacementFn paths.FetchElementReplacementFn) AdapterBuilder {
	app.tokenReplacementFn = tokenReplacementFn
	return app
}

// CanContainPrefix adds the builder the possiblity to contain a prefix
func (app *adapterBuilder) CanContainPrefix() AdapterBuilder {
	app.canContainPrefix = true
	return app
}

// CanContainSuffix adds the builder the possiblity to contain a suffix
func (app *adapterBuilder) CanContainSuffix() AdapterBuilder {
	app.canContainSuffix = true
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.root == nil {
		return nil, errors.New("the root Token is mandatory in order to build an Adapter instance")
	}

	if app.tokens == nil {
		return nil, errors.New("the Tokens are mandatory in order to build an Adapter instance")
	}

	astTokenAdapterBuilder := app.astTokensAdapterBuilder.Create().WithToken(app.root).WithTokens(app.tokens)
	if app.channels != nil {
		astTokenAdapterBuilder.WithChannels(app.channels)
	}

	if app.tokenNotFoundFn != nil {
		astTokenAdapterBuilder.WithTokenNotFoundFunc(app.tokenNotFoundFn)
	}

	if app.tokenReplacementFn != nil {
		astTokenAdapterBuilder.WithTokenReplacementFunc(app.tokenReplacementFn)
	}

	if app.canContainPrefix {
		astTokenAdapterBuilder.CanContainPrefix()
	}

	if app.canContainSuffix {
		astTokenAdapterBuilder.CanContainSuffix()
	}

	astTokenAdapter, err := astTokenAdapterBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createAdapter(
		app.mistakeAdapter,
		astTokenAdapter,
		app.builder,
		app.tokenBuilder,
		app.lineMatchBuilder,
		app.matchBuilder,
		app.tokenMatchBuilder,
		app.ruleMatchBuilder,
		app.resultBuilder,
		app.discoveriesBuilder,
		app.canContainPrefix,
	), nil
}
