package tokens

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

func createAdapterForTests(patternsScript string, rulesScript string, tokensScript string, channelScript string, rootTokenName string) Adapter {
	var ptrns patterns.Patterns
	if patternsScript != "" {
		patterns, err := patterns.NewAdapter().ToPatterns(patternsScript)
		if err != nil {
			panic(err)
		}

		ptrns = patterns
	}

	rules, err := rules.NewAdapter().Rules(rulesScript, ptrns)
	if err != nil {
		panic(err)
	}

	tokenAdapter, err := tokens.NewAdapterBuilder().Create().WithRules(rules).Now()
	if err != nil {
		panic(err)
	}

	tokens, err := tokenAdapter.ToTokens(tokensScript)
	if err != nil {
		panic(err)
	}

	rootToken, err := tokens.Find(rootTokenName)
	if err != nil {
		panic(err)
	}

	builder := NewAdapterBuilder().Create().WithTokens(tokens).WithToken(rootToken)
	if channelScript != "" {
		channels, err := tokenAdapter.ToTokens(channelScript)
		if err != nil {
			panic(err)
		}

		builder.WithChannels(channels)
	}

	adapter, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return adapter
}
