package suites

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

type application struct {
	builder            Builder
	lineBuilder        LineBuilder
	astAdapterBuilder  asts.AdapterBuilder
	tokens             tokens.Tokens
	channels           tokens.Tokens
	tokenNotFoundFn    paths.FetchTokenNotFoundFn
	tokenReplacementFn paths.FetchElementReplacementFn
}

func createApplication(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
) Application {
	return createApplicationInternally(
		builder,
		lineBuilder,
		astAdapterBuilder,
		tokens,
		nil,
		nil,
		nil,
	)
}

func createApplicationWithChannels(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
	channels tokens.Tokens,
) Application {
	return createApplicationInternally(
		builder,
		lineBuilder,
		astAdapterBuilder,
		tokens,
		channels,
		nil,
		nil,
	)
}

func createApplicationWithTokenNotFoundFunc(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
	tokenNotFoundFn paths.FetchTokenNotFoundFn,
) Application {
	return createApplicationInternally(
		builder,
		lineBuilder,
		astAdapterBuilder,
		tokens,
		nil,
		tokenNotFoundFn,
		nil,
	)
}

func createApplicationWithTokenReplacementFunc(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
	tokenReplacementFn paths.FetchElementReplacementFn,
) Application {
	return createApplicationInternally(
		builder,
		lineBuilder,
		astAdapterBuilder,
		tokens,
		nil,
		nil,
		tokenReplacementFn,
	)
}

func createApplicationWithTokenNotFoundAndReplacementFunc(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
	tokenNotFoundFn paths.FetchTokenNotFoundFn,
	tokenReplacementFn paths.FetchElementReplacementFn,
) Application {
	return createApplicationInternally(
		builder,
		lineBuilder,
		astAdapterBuilder,
		tokens,
		nil,
		tokenNotFoundFn,
		tokenReplacementFn,
	)
}

func createApplicationWithChannelsAndTokenNotFoundFunc(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
	channels tokens.Tokens,
	tokenNotFoundFn paths.FetchTokenNotFoundFn,
) Application {
	return createApplicationInternally(
		builder,
		lineBuilder,
		astAdapterBuilder,
		tokens,
		channels,
		tokenNotFoundFn,
		nil,
	)
}

func createApplicationWithChannelsAndTokenReplacementFunc(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
	channels tokens.Tokens,
	tokenReplacementFn paths.FetchElementReplacementFn,
) Application {
	return createApplicationInternally(
		builder,
		lineBuilder,
		astAdapterBuilder,
		tokens,
		channels,
		nil,
		tokenReplacementFn,
	)
}

func createApplicationWithChannelsAndTokenNotFoundAndReplacementFunc(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
	channels tokens.Tokens,
	tokenNotFoundFn paths.FetchTokenNotFoundFn,
	tokenReplacementFn paths.FetchElementReplacementFn,
) Application {
	return createApplicationInternally(
		builder,
		lineBuilder,
		astAdapterBuilder,
		tokens,
		channels,
		tokenNotFoundFn,
		tokenReplacementFn,
	)
}

func createApplicationInternally(
	builder Builder,
	lineBuilder LineBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	tokens tokens.Tokens,
	channels tokens.Tokens,
	tokenNotFoundFn paths.FetchTokenNotFoundFn,
	tokenReplacementFn paths.FetchElementReplacementFn,
) Application {
	out := application{
		builder:            builder,
		lineBuilder:        lineBuilder,
		astAdapterBuilder:  astAdapterBuilder,
		tokens:             tokens,
		channels:           channels,
		tokenNotFoundFn:    tokenNotFoundFn,
		tokenReplacementFn: tokenReplacementFn,
	}

	return &out
}

// Execute executes the tests
func (app *application) Execute() ([]Suite, error) {
	suites := []Suite{}
	tokensList := app.tokens.All()
	for _, oneToken := range tokensList {
		if !oneToken.HasTestSuite() {
			continue
		}

		name := oneToken.Name()
		testLines := oneToken.TestSuite().Lines()
		suite, err := app.toSuite(name, testLines, app.tokens, app.channels)
		if err != nil {
			return nil, err
		}

		suites = append(suites, suite)
	}

	return suites, nil
}

func (app *application) toSuite(name string, testLines []string, tokens tokens.Tokens, channels tokens.Tokens) (Suite, error) {
	lines := []Line{}
	for idx, testLine := range testLines {
		line, err := app.toLine(name, uint(idx), testLine, tokens, channels)
		if err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	return app.builder.Create().WithName(name).WithLines(lines).Now()
}

func (app *application) toLine(rootToken string, index uint, testLine string, tokens tokens.Tokens, channels tokens.Tokens) (Line, error) {
	root, err := app.tokens.Find(rootToken)
	if err != nil {
		return nil, err
	}

	astAdapterBuilder := app.astAdapterBuilder.Create().WithRoot(root).WithTokens(tokens)
	if channels != nil {
		astAdapterBuilder.WithChannels(channels)
	}

	if app.tokenNotFoundFn != nil {
		astAdapterBuilder.WithTokenNotFoundFunc(app.tokenNotFoundFn)
	}

	if app.tokenReplacementFn != nil {
		astAdapterBuilder.WithTokenReplacementFunc(app.tokenReplacementFn)
	}

	astAdapter, err := astAdapterBuilder.Now()
	if err != nil {
		return nil, err
	}

	ast, err := astAdapter.ToAST(testLine)
	if err != nil {
		return nil, err
	}

	lineBuilder := app.lineBuilder.Create().WithIndex(index)
	if ast.IsSuccess() {
		lineBuilder.IsSuccessful()
	}

	return lineBuilder.Now()
}
