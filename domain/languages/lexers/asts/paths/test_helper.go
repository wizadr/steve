package paths

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

func toPathForTests(patternsScript string, rulesScript string, tokensScript string, tokenName string) Path {
	tokens := toTokensForTests(patternsScript, rulesScript, tokensScript)
	adapter, err := NewAdapterBuilder().Create().WithTokens(tokens).Now()
	if err != nil {
		panic(err)
	}

	token, err := tokens.Find(tokenName)
	if err != nil {
		panic(err)
	}

	path, err := adapter.ToPath(token)
	if err != nil {
		panic(err)
	}

	return path
}

func toTokensForTests(patternsScript string, rulesScript string, tokensScript string) tokens.Tokens {
	patterns, err := patterns.NewAdapter().ToPatterns(patternsScript)
	if err != nil {
		panic(err)
	}

	rules, err := rules.NewAdapter().Rules(rulesScript, patterns)
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

	return tokens
}
