package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/results"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

type adapterBuilder struct {
	builder                   Builder
	tokenContentBuilder       TokenContentBuilder
	lineMatchBuilder          LineMatchBuilder
	blockMatchBuilder         BlockMatchBuilder
	nextElementBuilder        NextElementBuilder
	matchesBuilder            MatchesBuilder
	matchBuilder              MatchBuilder
	skipMatchBuilder          SkipMatchBuilder
	tokenMatchBuilder         TokenMatchBuilder
	specifierBuilder          SpecifierBuilder
	tokenResultBuilder        TokenResultBuilder
	tokenResultMatchesBuilder TokenResultMatchesBuilder
	ruleMatchBuilder          RuleMatchBuilder
	ruleMatchResultBuilder    RuleMatchResultBuilder
	resultBuilder             ResultBuilder
	resultAdapter             results.Adapter
	pathTokenPathBuilder      paths.TokenPathBuilder
	pathTokenBuilder          paths.TokenBuilder
	pathBuilder               paths.Builder
	pathAdapterBuilder        paths.AdapterBuilder
	rootToken                 tokens.Token
	tokens                    tokens.Tokens
	channels                  tokens.Tokens
	tokenNotFoundFn           paths.FetchTokenNotFoundFn
	tokenReplacementFn        paths.FetchElementReplacementFn
	canContainPrefix          bool
	canContainSuffix          bool
}

func createAdapterBuilder(
	builder Builder,
	tokenContentBuilder TokenContentBuilder,
	lineMatchBuilder LineMatchBuilder,
	blockMatchBuilder BlockMatchBuilder,
	nextElementBuilder NextElementBuilder,
	matchesBuilder MatchesBuilder,
	matchBuilder MatchBuilder,
	skipMatchBuilder SkipMatchBuilder,
	tokenMatchBuilder TokenMatchBuilder,
	specifierBuilder SpecifierBuilder,
	tokenResultBuilder TokenResultBuilder,
	tokenResultMatchesBuilder TokenResultMatchesBuilder,
	ruleMatchBuilder RuleMatchBuilder,
	ruleMatchResultBuilder RuleMatchResultBuilder,
	resultBuilder ResultBuilder,
	resultAdapter results.Adapter,
	pathTokenPathBuilder paths.TokenPathBuilder,
	pathTokenBuilder paths.TokenBuilder,
	pathBuilder paths.Builder,
	pathAdapterBuilder paths.AdapterBuilder,
) AdapterBuilder {
	out := adapterBuilder{
		builder:                   builder,
		tokenContentBuilder:       tokenContentBuilder,
		lineMatchBuilder:          lineMatchBuilder,
		blockMatchBuilder:         blockMatchBuilder,
		nextElementBuilder:        nextElementBuilder,
		matchesBuilder:            matchesBuilder,
		matchBuilder:              matchBuilder,
		skipMatchBuilder:          skipMatchBuilder,
		tokenMatchBuilder:         tokenMatchBuilder,
		specifierBuilder:          specifierBuilder,
		tokenResultBuilder:        tokenResultBuilder,
		tokenResultMatchesBuilder: tokenResultMatchesBuilder,
		ruleMatchBuilder:          ruleMatchBuilder,
		ruleMatchResultBuilder:    ruleMatchResultBuilder,
		resultBuilder:             resultBuilder,
		resultAdapter:             resultAdapter,
		pathTokenPathBuilder:      pathTokenPathBuilder,
		pathTokenBuilder:          pathTokenBuilder,
		pathBuilder:               pathBuilder,
		pathAdapterBuilder:        pathAdapterBuilder,
		rootToken:                 nil,
		tokens:                    nil,
		channels:                  nil,
		tokenNotFoundFn:           nil,
		tokenReplacementFn:        nil,
		canContainPrefix:          false,
		canContainSuffix:          false,
	}

	return &out
}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(
		app.builder,
		app.tokenContentBuilder,
		app.lineMatchBuilder,
		app.blockMatchBuilder,
		app.nextElementBuilder,
		app.matchesBuilder,
		app.matchBuilder,
		app.skipMatchBuilder,
		app.tokenMatchBuilder,
		app.specifierBuilder,
		app.tokenResultBuilder,
		app.tokenResultMatchesBuilder,
		app.ruleMatchBuilder,
		app.ruleMatchResultBuilder,
		app.resultBuilder,
		app.resultAdapter,
		app.pathTokenPathBuilder,
		app.pathTokenBuilder,
		app.pathBuilder,
		app.pathAdapterBuilder,
	)
}

// WithToken adds a root token to the builder
func (app *adapterBuilder) WithToken(rootToken tokens.Token) AdapterBuilder {
	app.rootToken = rootToken
	return app
}

// WithTokens add tokens to the builder
func (app *adapterBuilder) WithTokens(tokens tokens.Tokens) AdapterBuilder {
	app.tokens = tokens
	return app
}

// WithChannels add channels to the builder
func (app *adapterBuilder) WithChannels(channels tokens.Tokens) AdapterBuilder {
	app.channels = channels
	return app
}

// WithTokenNotFoundFunc add tokenNotFoundFunc to the builder
func (app *adapterBuilder) WithTokenNotFoundFunc(tokenNotFoundFn paths.FetchTokenNotFoundFn) AdapterBuilder {
	app.tokenNotFoundFn = tokenNotFoundFn
	return app
}

// WithTokenReplacementFunc adds a tokenReplacementFunc to the builder
func (app *adapterBuilder) WithTokenReplacementFunc(tokenReplacementFn paths.FetchElementReplacementFn) AdapterBuilder {
	app.tokenReplacementFn = tokenReplacementFn
	return app
}

// CanContainPrefix flags the builder with the ability to contain a prefix
func (app *adapterBuilder) CanContainPrefix() AdapterBuilder {
	app.canContainPrefix = true
	return app
}

// CanContainSuffix flags the builder with the ability to contain a suffix
func (app *adapterBuilder) CanContainSuffix() AdapterBuilder {
	app.canContainSuffix = true
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.rootToken == nil {
		return nil, errors.New("the root token is mandatory in order to build an Adapter instance")
	}

	if app.tokens == nil {
		return nil, errors.New("the tokens are mandatory in order to build an Adapter instance")
	}

	pathAdapterBuilder := app.pathAdapterBuilder.Create().WithTokens(app.tokens)
	if app.tokenNotFoundFn != nil {
		pathAdapterBuilder.WithTokenNotFoundFunc(app.tokenNotFoundFn)
	}

	if app.tokenReplacementFn != nil {
		pathAdapterBuilder.WithTokenReplacementFunc(app.tokenReplacementFn)
	}

	pathAdapter, err := pathAdapterBuilder.Now()
	if err != nil {
		return nil, err
	}

	rootPath, err := pathAdapter.ToPath(app.rootToken)
	if err != nil {
		return nil, err
	}

	if app.channels != nil {
		channelPaths := map[string]paths.Path{}
		channels := app.channels.All()
		for _, oneChannelToken := range channels {
			channelPath, err := pathAdapter.ToPath(oneChannelToken)
			if err != nil {
				return nil, err
			}

			name := oneChannelToken.Name()
			channelPaths[name] = channelPath
		}

		return createAdapterWithChannels(
			app.builder,
			app.tokenContentBuilder,
			app.lineMatchBuilder,
			app.blockMatchBuilder,
			app.nextElementBuilder,
			app.matchesBuilder,
			app.matchBuilder,
			app.skipMatchBuilder,
			app.tokenMatchBuilder,
			app.specifierBuilder,
			app.tokenResultBuilder,
			app.tokenResultMatchesBuilder,
			app.ruleMatchBuilder,
			app.ruleMatchResultBuilder,
			app.resultBuilder,
			app.resultAdapter,
			app.pathTokenPathBuilder,
			app.pathTokenBuilder,
			app.pathBuilder,
			rootPath,
			channelPaths,
			app.canContainPrefix,
			app.canContainSuffix,
		), nil
	}

	return createAdapter(
		app.builder,
		app.tokenContentBuilder,
		app.lineMatchBuilder,
		app.blockMatchBuilder,
		app.nextElementBuilder,
		app.matchesBuilder,
		app.matchBuilder,
		app.skipMatchBuilder,
		app.tokenMatchBuilder,
		app.specifierBuilder,
		app.tokenResultBuilder,
		app.tokenResultMatchesBuilder,
		app.ruleMatchBuilder,
		app.ruleMatchResultBuilder,
		app.resultBuilder,
		app.resultAdapter,
		app.pathTokenPathBuilder,
		app.pathTokenBuilder,
		app.pathBuilder,
		rootPath,
		app.canContainPrefix,
		app.canContainSuffix,
	), nil
}
