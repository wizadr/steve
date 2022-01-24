package mistakes

import (
	"strings"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/results"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type adapter struct {
	builder                      Builder
	indexBuilder                 IndexBuilder
	pathBuilder                  PathBuilder
	contentBuilder               ContentBuilder
	containsValidNotMatchBuilder ContainsValidNotMatchBuilder
	specifierDoNotMatchBuilder   SpecifierDoNotMatchBuilder
	containsNextElementBuilder   ContainsNextElementBuilder
	cardinalityIsInvalidBuilder  CardinalityIsInvalidBuilder
	containsPrefixBuilder        ContainsPrefixBuilder
	newLineCharacter             string
	currentScript                string
}

func createAdapter(
	builder Builder,
	indexBuilder IndexBuilder,
	pathBuilder PathBuilder,
	contentBuilder ContentBuilder,
	containsValidNotMatchBuilder ContainsValidNotMatchBuilder,
	specifierDoNotMatchBuilder SpecifierDoNotMatchBuilder,
	containsNextElementBuilder ContainsNextElementBuilder,
	cardinalityIsInvalidBuilder CardinalityIsInvalidBuilder,
	containsPrefixBuilder ContainsPrefixBuilder,
	newLineCharacter string,
) Adapter {
	out := adapter{
		builder:                      builder,
		indexBuilder:                 indexBuilder,
		pathBuilder:                  pathBuilder,
		contentBuilder:               contentBuilder,
		containsValidNotMatchBuilder: containsValidNotMatchBuilder,
		specifierDoNotMatchBuilder:   specifierDoNotMatchBuilder,
		containsNextElementBuilder:   containsNextElementBuilder,
		cardinalityIsInvalidBuilder:  cardinalityIsInvalidBuilder,
		containsPrefixBuilder:        containsPrefixBuilder,
		newLineCharacter:             newLineCharacter,
		currentScript:                "",
	}

	return &out
}

// ToMistake converts an InnerToken instance to a Mistake instance
func (app *adapter) ToMistake(result tokens.Token, canContainPrefix bool) (Mistake, error) {
	app.currentScript = ""
	return app.token([]string{}, result, canContainPrefix)
}

func (app *adapter) token(parentTokenNames []string, result tokens.Token, canContainPrefix bool) (Mistake, error) {
	if !result.HasContent() {
		return app.rootTokenNotFound(parentTokenNames, result, canContainPrefix)
	}

	// verify if there is not matches:
	name := result.Name()
	content := result.Content()
	dependencies := result.Path().Dependencies()
	if content.HasNot() {
		not := content.Not()
		notMistake, err := app.lineMatchAndResetCurrentScript(parentTokenNames, name, not, dependencies, canContainPrefix)
		if err != nil {
			return nil, err
		}

		if notMistake == nil {
			index := not.Index()
			containsValidNotMatch, err := app.containsValidNotMatchBuilder.Create().WithLine(index).Now()
			if err != nil {
				return nil, err
			}

			content, err := app.contentBuilder.Create().WithContainsValidNotMatch(containsValidNotMatch).Now()
			if err != nil {
				return nil, err
			}

			return app.buildMistakeFromToken(
				name,
				parentTokenNames,
				app.currentScript,
				content,
			)
		}
	}

	must := content.Must()
	out, err := app.lineMatchAndResetCurrentScript(parentTokenNames, name, must, dependencies, canContainPrefix)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (app *adapter) rootTokenNotFound(parentTokenNames []string, result tokens.Token, canContainPrefix bool) (Mistake, error) {
	name := result.Name()
	pathContainer, err := app.pathBuilder.Create().WithParents(parentTokenNames).WithToken(name).Now()
	if err != nil {
		return nil, err
	}

	content, err := app.contentBuilder.Create().IsTokenHasNoMatch().Now()
	if err != nil {
		return nil, err
	}

	return app.buildMistakeFromPath(pathContainer, "", content)
}

func (app *adapter) lineMatchAndResetCurrentScript(parentTokenNames []string, tokenName string, lineMatch tokens.LineMatch, dependencies paths.Dependencies, canContainPrefix bool) (Mistake, error) {
	app.currentScript = ""
	return app.lineMatch(parentTokenNames, tokenName, lineMatch, dependencies, canContainPrefix)
}

func (app *adapter) lineMatch(parentTokenNames []string, tokenName string, lineMatch tokens.LineMatch, dependencies paths.Dependencies, canContainPrefix bool) (Mistake, error) {
	matches := lineMatch.Matches()
	return app.matches(parentTokenNames, tokenName, matches, dependencies, canContainPrefix)
}

func (app *adapter) matches(parentTokenNames []string, tokenName string, matches tokens.Matches, dependencies paths.Dependencies, canContainPrefix bool) (Mistake, error) {
	blockMatches := matches.BlockMatches()
	for _, oneBlockMatch := range blockMatches {
		mistake, err := app.blockMatch(parentTokenNames, tokenName, oneBlockMatch, dependencies, canContainPrefix)
		if err != nil {
			return nil, err
		}

		if mistake != nil {
			return mistake, nil
		}
	}

	return nil, nil
}

func (app *adapter) blockMatch(parentTokenNames []string, tokenName string, blockMatch tokens.BlockMatch, dependencies paths.Dependencies, canContainPrefix bool) (Mistake, error) {
	if blockMatch.HasChannelPrefix() {
		channelPrefix := blockMatch.ChannelPrefix().Matches()
		app.currentScript = strings.Join([]string{
			app.currentScript,
			channelPrefix.Discoveries(),
		}, "")
	}

	content := blockMatch.Content()
	if content.IsContainer() {
		container := content.Container()
		return app.match(parentTokenNames, tokenName, container, dependencies, canContainPrefix)
	}

	if content.IsBlock() {
		block := content.Block()
		return app.blockMatch(parentTokenNames, tokenName, block, dependencies, canContainPrefix)
	}

	nextElement := content.NextElement()
	containsNextElement, err := app.containsNextElementBuilder.Create().WithNextElement(nextElement).Now()
	if err != nil {
		return nil, err
	}

	contentIns, err := app.contentBuilder.Create().WithContainsNextElement(containsNextElement).Now()
	if err != nil {
		return nil, err
	}

	return app.buildMistakeFromToken(
		tokenName,
		parentTokenNames,
		app.currentScript,
		contentIns,
	)
}

func (app *adapter) match(parentTokenNames []string, tokenName string, match tokens.Match, dependencies paths.Dependencies, canContainPrefix bool) (Mistake, error) {
	content := match.Content()
	if content.IsToken() {
		token := content.Token()
		return app.tokenMatch(parentTokenNames, tokenName, token, dependencies, canContainPrefix)
	}

	if content.IsSkip() {
		return nil, nil
	}

	if content.IsLine() {
		lineMatch := content.Line()
		return app.lineMatch(parentTokenNames, tokenName, lineMatch, dependencies, canContainPrefix)
	}

	rule := content.Rule()
	return app.ruleMatch(parentTokenNames, rule, canContainPrefix)
}

func (app *adapter) tokenMatch(parentTokenNames []string, tokenName string, tokenMatch tokens.TokenMatch, dependencies paths.Dependencies, canContainPrefix bool) (Mistake, error) {
	result := tokenMatch.Result()
	element := tokenMatch.Path().Element()
	mistake, err := app.tokenResult(parentTokenNames, tokenName, result, element, canContainPrefix)
	if err != nil {
		return nil, err
	}

	if mistake != nil {
		return mistake, nil
	}

	if tokenMatch.HasSpecifiers() {
		tokenPath := tokenMatch.Path()
		specifiers := tokenMatch.Specifiers()
		return app.tokenSpecifiers(parentTokenNames, tokenName, tokenPath, specifiers)
	}

	return nil, nil
}

func (app *adapter) tokenSpecifiers(parentTokenNames []string, tokenName string, tokenPath paths.TokenPath, specifiers []tokens.Specifier) (Mistake, error) {
	for _, oneSpecifier := range specifiers {
		if oneSpecifier.IsValid() {
			continue
		}

		amount := oneSpecifier.Amount()
		cardinality := oneSpecifier.Cardinality()
		specifierContainerName := oneSpecifier.ContainerName()
		specifierDoNotMatch, err := app.specifierDoNotMatchBuilder.Create().WithAmount(amount).WithCardinality(cardinality).WithContainerName(specifierContainerName).Now()
		if err != nil {
			return nil, err
		}

		content, err := app.contentBuilder.Create().WithSpecifierDoNotMatch(specifierDoNotMatch).Now()
		if err != nil {
			return nil, err
		}

		return app.buildMistakeFromToken(
			tokenName,
			parentTokenNames,
			app.currentScript,
			content,
		)
	}

	return nil, nil
}

func (app *adapter) tokenResult(parentTokenNames []string, tokenName string, tokenResult tokens.TokenResult, path paths.Element, canContainPrefix bool) (Mistake, error) {
	if !tokenResult.HasMatches() {
		if tokenResult.IsCardinalityValid() {
			return nil, nil
		}

		content, err := app.contentBuilder.Create().IsTokenHasNoMatch().Now()
		if err != nil {
			return nil, err
		}

		return app.buildMistakeFromToken(
			tokenName,
			parentTokenNames,
			app.currentScript,
			content,
		)
	}

	if !tokenResult.IsCardinalityValid() {
		amount := tokenResult.Amount()
		cardinality := tokenResult.Cardinality()
		cardinalityIsInvalid, err := app.cardinalityIsInvalidBuilder.Create().WithAmount(amount).WithCardinality(cardinality).Now()
		if err != nil {
			return nil, err
		}

		content, err := app.contentBuilder.Create().WithCardinalityIsInvalid(cardinalityIsInvalid).Now()
		if err != nil {
			return nil, err
		}

		return app.buildMistakeFromToken(
			tokenName,
			parentTokenNames,
			app.currentScript,
			content,
		)
	}

	matches := tokenResult.Matches()
	parentTokenNames = append(parentTokenNames, tokenName)
	return app.tokenResultMatches(parentTokenNames, matches, canContainPrefix)
}

func (app *adapter) tokenResultMatches(parentTokenNames []string, matches tokens.TokenResultMatches, canContainPrefix bool) (Mistake, error) {
	var output Mistake
	tokens := matches.All()
	for _, oneToken := range tokens {
		mistake, err := app.token(parentTokenNames, oneToken, canContainPrefix)
		if err != nil {
			return nil, err
		}

		if mistake != nil {
			output = mistake
			break
		}
	}

	return output, nil
}

func (app *adapter) ruleMatch(parentTokenNames []string, ruleMatch tokens.RuleMatch, canContainPrefix bool) (Mistake, error) {
	ruleName := ruleMatch.Rule().Base().Name()
	ruleMatchResult := ruleMatch.Result()
	if !ruleMatchResult.IsCardinalityValid() {
		amount := ruleMatchResult.Result().Amount()
		cardinality := ruleMatch.Rule().Cardinality()
		cardinalityIsInvalid, err := app.cardinalityIsInvalidBuilder.Create().WithAmount(amount).WithCardinality(cardinality).Now()
		if err != nil {
			return nil, err
		}

		content, err := app.contentBuilder.Create().WithCardinalityIsInvalid(cardinalityIsInvalid).Now()
		if err != nil {
			return nil, err
		}

		return app.buildMistakeFromRule(
			ruleName,
			parentTokenNames,
			app.currentScript,
			content,
		)
	}

	result := ruleMatchResult.Result()
	if result.HasResults() {
		list := result.Results()
		for _, oneResult := range list {
			mistake, err := app.ruleResult(parentTokenNames, ruleName, oneResult, canContainPrefix)
			if err != nil {
				return nil, err
			}

			return mistake, nil
		}
	}

	return nil, nil
}

func (app *adapter) ruleResult(parentTokenNames []string, ruleName string, result results.Result, canContainPrefix bool) (Mistake, error) {
	// the cardinality was valid, but there is no match, therefore return no mistake:
	if !result.HasMatches() {
		return nil, nil
	}

	matches := result.Matches().List()
	return app.patternResults(parentTokenNames, ruleName, matches, canContainPrefix)
}

func (app *adapter) patternResults(parentTokenNames []string, ruleName string, results []patterns.Result, canContainPrefix bool) (Mistake, error) {
	for _, oneResult := range results {
		mistake, err := app.patternResult(parentTokenNames, ruleName, oneResult, canContainPrefix)
		if err != nil {
			return nil, err
		}

		if mistake != nil {
			return mistake, nil
		}
	}

	return nil, nil
}

func (app *adapter) patternResult(parentTokenNames []string, ruleName string, match patterns.Result, canContainPrefix bool) (Mistake, error) {
	app.currentScript = strings.Join([]string{
		app.currentScript,
		match.Discoveries().Content(),
	}, "")

	index := match.Discoveries().Index()
	if !canContainPrefix && index > 0 {
		input := match.Input()
		containsPrefix, err := app.containsPrefixBuilder.Create().WithPrefix(input[:int(index)]).Now()
		if err != nil {
			return nil, err
		}

		content, err := app.contentBuilder.Create().WithContainsPrefix(containsPrefix).Now()
		if err != nil {
			return nil, err
		}

		return app.buildMistakeFromRule(
			ruleName,
			parentTokenNames,
			app.currentScript,
			content,
		)
	}

	return nil, nil
}

func (app *adapter) buildMistakeFromToken(name string, tokenParents []string, script string, content Content) (Mistake, error) {
	path, err := app.pathBuilder.Create().WithToken(name).WithParents(tokenParents).Now()
	if err != nil {
		return nil, err
	}

	return app.buildMistakeFromPath(path, script, content)
}

func (app *adapter) buildMistakeFromRule(name string, tokenParents []string, script string, content Content) (Mistake, error) {
	path, err := app.pathBuilder.Create().WithRule(name).WithParents(tokenParents).Now()
	if err != nil {
		return nil, err
	}

	return app.buildMistakeFromPath(path, script, content)
}

func (app *adapter) buildMistakeFromPath(path Path, script string, content Content) (Mistake, error) {
	idx, line, column := app.getCurrentPosition(script)
	index, err := app.indexBuilder.Create().WithIndex(idx).WithLine(line).WithColumn(column).Now()
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithPath(path).
		WithIndex(index).
		WithContent(content).
		Now()
}

func (app *adapter) getCurrentPosition(content string) (uint, uint, uint) {
	lines := strings.Split(content, app.newLineCharacter)
	amountLines := len(lines)
	if amountLines <= 0 {
		return 0, 0, 0
	}

	lastLine := lines[amountLines-1]
	amountColumns := len(lastLine)
	if amountColumns <= 0 {
		return uint(len(content)), uint(amountLines) - 1, 0
	}

	return uint(len(content)), uint(amountLines) - 1, uint(amountColumns) - 1
}
