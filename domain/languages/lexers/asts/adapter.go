package asts

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
	ast_results "github.com/steve-care-software/steve/domain/languages/lexers/asts/results"
	ast_tokens "github.com/steve-care-software/steve/domain/languages/lexers/asts/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type adapter struct {
	mistakeAdapter     mistakes.Adapter
	astTokenAdapter    ast_tokens.Adapter
	builder            Builder
	tokenBuilder       TokenBuilder
	lineMatchBuilder   LineMatchBuilder
	matchBuilder       MatchBuilder
	tokenMatchBuilder  TokenMatchBuilder
	ruleMatchBuilder   RuleMatchBuilder
	resultBuilder      ResultBuilder
	discoveriesBuilder patterns.DiscoveriesBuilder
	canContainPrefix   bool
}

func createAdapter(
	mistakeAdapter mistakes.Adapter,
	astTokenAdapter ast_tokens.Adapter,
	builder Builder,
	tokenBuilder TokenBuilder,
	lineMatchBuilder LineMatchBuilder,
	matchBuilder MatchBuilder,
	tokenMatchBuilder TokenMatchBuilder,
	ruleMatchBuilder RuleMatchBuilder,
	resultBuilder ResultBuilder,
	discoveriesBuilder patterns.DiscoveriesBuilder,
	canContainPrefix bool,
) Adapter {
	out := adapter{
		mistakeAdapter:     mistakeAdapter,
		astTokenAdapter:    astTokenAdapter,
		builder:            builder,
		tokenBuilder:       tokenBuilder,
		lineMatchBuilder:   lineMatchBuilder,
		matchBuilder:       matchBuilder,
		tokenMatchBuilder:  tokenMatchBuilder,
		ruleMatchBuilder:   ruleMatchBuilder,
		resultBuilder:      resultBuilder,
		discoveriesBuilder: discoveriesBuilder,
		canContainPrefix:   canContainPrefix,
	}

	return &out
}

// ToAST converts a script to an AST instance
func (app *adapter) ToAST(script string) (AST, error) {
	astToken, err := app.astTokenAdapter.ToToken(script)
	if err != nil {
		return nil, err
	}

	mistake, err := app.mistakeAdapter.ToMistake(astToken, app.canContainPrefix)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create()
	if mistake != nil {
		return builder.WithMistake(mistake).Now()
	}

	token, err := app.toToken(astToken)
	if err != nil {
		return nil, err
	}

	return builder.WithSuccess(token).Now()
}

func (app *adapter) toToken(token ast_tokens.Token) (Token, error) {
	name := token.Name()
	must := token.Content().Must()
	match, err := app.lineMatch(name, must)
	if err != nil {
		return nil, err
	}

	return app.tokenBuilder.Create().WithName(name).WithMatch(match).Now()
}

func (app *adapter) lineMatch(tokenName string, lineMatch ast_tokens.LineMatch) (LineMatch, error) {
	matches := lineMatch.Matches()
	matchList, err := app.matches(tokenName, matches)
	if err != nil {
		return nil, err
	}

	index := lineMatch.Index()
	return app.lineMatchBuilder.Create().WithIndex(index).WithMatches(matchList).Now()
}

func (app *adapter) matches(tokenName string, matches ast_tokens.Matches) ([]Match, error) {
	blockMatches := matches.BlockMatches()
	return app.blockMatches(tokenName, blockMatches)
}

func (app *adapter) blockMatches(tokenName string, blockMatches []ast_tokens.BlockMatch) ([]Match, error) {
	out := []Match{}
	for _, oneBlockMatch := range blockMatches {
		matches, err := app.blockMatch(tokenName, oneBlockMatch)
		if err != nil {
			return nil, err
		}

		out = append(out, matches...)
	}

	return out, nil
}

func (app *adapter) blockMatch(tokenName string, blockMatch ast_tokens.BlockMatch) ([]Match, error) {
	matchList := []Match{}
	if blockMatch.HasChannelPrefix() {
		blockMatches := blockMatch.ChannelPrefix().Matches().BlockMatches()
		matches, err := app.blockMatches(tokenName, blockMatches)
		if err != nil {
			return nil, err
		}

		matchList = append(matchList, matches...)
	}

	content := blockMatch.Content()
	if content.IsNextElement() {
		str := fmt.Sprintf("the element after the token (name: %s) was expecting a nextElement", tokenName)
		return nil, errors.New(str)
	}

	if content.IsContainer() {
		container := content.Container()
		match, err := app.match(container)
		if err != nil {
			return nil, err
		}

		if match != nil {
			matchList = append(matchList, match)
		}
	}

	if content.IsBlock() {
		blockMatch := content.Block()
		matches, err := app.blockMatch(tokenName, blockMatch)
		if err != nil {
			return nil, err
		}

		matchList = append(matchList, matches...)
	}

	return matchList, nil
}

func (app *adapter) match(match ast_tokens.Match) (Match, error) {
	builder := app.matchBuilder.Create()
	content := match.Content()
	if content.IsToken() {
		token := content.Token()
		tokenMatch, err := app.tokenMatch(token)
		if err != nil {
			return nil, err
		}

		if tokenMatch == nil {
			return nil, nil
		}

		builder.WithToken(tokenMatch)
	}

	if content.IsRule() {
		rule := content.Rule()
		ruleMatch, err := app.ruleMatch(rule)
		if err != nil {
			return nil, err
		}

		if ruleMatch == nil {
			return nil, nil
		}

		builder.WithRule(ruleMatch)
	}

	if content.IsSkip() {
		return nil, nil
	}

	return builder.Now()
}

func (app *adapter) tokenMatch(tokenMatch ast_tokens.TokenMatch) (TokenMatch, error) {
	result := tokenMatch.Result()
	name := tokenMatch.Path().Element().Name()
	if !result.HasMatches() {
		return nil, nil
	}

	out := []Token{}
	tokens := result.Matches().All()
	for _, oneToken := range tokens {
		token, err := app.toToken(oneToken)
		if err != nil {
			return nil, err
		}

		out = append(out, token)
	}

	return app.tokenMatchBuilder.Create().WithToken(name).WithMatches(out).Now()
}

func (app *adapter) ruleMatch(ruleMatch ast_tokens.RuleMatch) (RuleMatch, error) {
	name := ruleMatch.Rule().Base().Name()
	astResult := ruleMatch.Result().Result()
	input := ruleMatch.Result().Path().Base().Code()
	result, err := app.ruleResult(name, astResult, input)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return app.ruleMatchBuilder.Create().WithRule(name).WithResult(result).Now()
}

func (app *adapter) ruleResult(rule string, ruleResult ast_tokens.Result, input string) (Result, error) {
	if !ruleResult.HasResults() {
		return nil, nil
	}

	firstIndex := uint(0)
	results := ruleResult.Results()
	discoveriesContent := ruleResult.Discoveries()
	discoveriesList := []patterns.Discovery{}
	for _, oneResult := range results {
		if !oneResult.HasMatches() {
			break
		}

		matches := oneResult.Matches().List()
		for idx, oneMatch := range matches {
			oneDiscoveries := oneMatch.Discoveries()
			discoveriesList = append(discoveriesList, oneDiscoveries.All()...)

			if idx <= 0 {
				firstIndex = oneMatch.Discoveries().Index()
			}
		}
	}

	if len(discoveriesList) <= 0 {
		return nil, nil
	}

	discoveries, err := app.discoveriesBuilder.Create().WithList(discoveriesList).Now()
	if err != nil {
		return nil, err
	}

	builder := app.resultBuilder.Create().WithInput(input).WithDiscoveries(discoveries)
	if firstIndex > 0 {
		builder.WithPrefix(input[:firstIndex])
	}

	inputLength := uint(len(input))
	contentWithPrefixLength := firstIndex + uint(len(discoveriesContent))
	if contentWithPrefixLength < inputLength {
		builder.WithSuffix(input[:contentWithPrefixLength])
	}

	return builder.Now()
}

func (app *adapter) results(rule string, results ast_results.Result) (Result, error) {
	input := results.Input().Code()
	if !results.HasMatches() {
		return nil, nil
	}

	matches := results.Matches().List()
	discoveriesContent := results.Content()
	discoveriesList := []patterns.Discovery{}
	for _, oneMatch := range matches {
		oneDiscoveries := oneMatch.Discoveries()
		discoveriesList = append(discoveriesList, oneDiscoveries.All()...)
	}

	discoveries, err := app.discoveriesBuilder.Create().WithList(discoveriesList).Now()
	if err != nil {
		return nil, err
	}

	index := matches[0].Discoveries().Index()
	builder := app.resultBuilder.Create().WithInput(input).WithDiscoveries(discoveries)
	if index > 0 {
		builder.WithPrefix(input[:index])
	}

	inputLength := uint(len(input))
	contentWithPrefixLength := index + uint(len(discoveriesContent))
	if contentWithPrefixLength < inputLength {
		builder.WithSuffix(input[:contentWithPrefixLength])
	}

	return builder.Now()
}
