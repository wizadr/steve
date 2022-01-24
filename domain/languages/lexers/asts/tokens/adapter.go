package tokens

import (
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/results"
	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

type adapter struct {
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
	rootPath                  paths.Path
	channels                  map[string]paths.Path
	canContainPrefix          bool
	canContainSuffix          bool
	recursiveTokens           map[string]string
}

func createAdapterWithChannels(
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
	rootPath paths.Path,
	channels map[string]paths.Path,
	canContainPrefix bool,
	canContainSuffix bool,
) Adapter {
	return createAdapterInternally(
		builder,
		tokenContentBuilder,
		lineMatchBuilder,
		blockMatchBuilder,
		nextElementBuilder,
		matchesBuilder,
		matchBuilder,
		skipMatchBuilder,
		tokenMatchBuilder,
		specifierBuilder,
		tokenResultBuilder,
		tokenResultMatchesBuilder,
		ruleMatchBuilder,
		ruleMatchResultBuilder,
		resultBuilder,
		resultAdapter,
		pathTokenPathBuilder,
		pathTokenBuilder,
		pathBuilder,
		rootPath,
		channels,
		canContainPrefix,
		canContainSuffix,
	)
}

func createAdapter(
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
	rootPath paths.Path,
	canContainPrefix bool,
	canContainSuffix bool,
) Adapter {
	return createAdapterInternally(
		builder,
		tokenContentBuilder,
		lineMatchBuilder,
		blockMatchBuilder,
		nextElementBuilder,
		matchesBuilder,
		matchBuilder,
		skipMatchBuilder,
		tokenMatchBuilder,
		specifierBuilder,
		tokenResultBuilder,
		tokenResultMatchesBuilder,
		ruleMatchBuilder,
		ruleMatchResultBuilder,
		resultBuilder,
		resultAdapter,
		pathTokenPathBuilder,
		pathTokenBuilder,
		pathBuilder,
		rootPath,
		nil,
		canContainPrefix,
		canContainSuffix,
	)
}

func createAdapterInternally(
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
	rootPath paths.Path,
	channels map[string]paths.Path,
	canContainPrefix bool,
	canContainSuffix bool,
) Adapter {
	out := adapter{
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
		rootPath:                  rootPath,
		channels:                  channels,
		canContainPrefix:          canContainPrefix,
		canContainSuffix:          canContainSuffix,
		recursiveTokens:           map[string]string{},
	}

	return &out
}

// ToToken converts a script to a Token instance
func (app *adapter) ToToken(script string) (Token, error) {
	app.recursiveTokens = map[string]string{}
	return app.elementToToken(app.rootPath.Element(), app.rootPath.Dependencies(), script, true)
}

func (app *adapter) elementToToken(element paths.Element, dependencies paths.Dependencies, script string, chansAreActivated bool) (Token, error) {
	name := element.Name()
	path, err := app.pathBuilder.Create().WithElement(element).WithDependencies(dependencies).Now()
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithName(name).WithPath(path)
	tokenContent, err := app.elementToTokenContent(element, dependencies, script, chansAreActivated)
	if err != nil {
		return nil, err
	}

	if tokenContent != nil {
		builder.WithContent(tokenContent)
	}

	return builder.Now()
}

func (app *adapter) elementToTokenContent(element paths.Element, dependencies paths.Dependencies, script string, chansAreActivated bool) (TokenContent, error) {
	mustLines := element.Must()
	must, err := app.bestLineToMatches(mustLines, dependencies, script, chansAreActivated)
	if err != nil {
		return nil, err
	}

	if must == nil {
		return nil, nil
	}

	builder := app.tokenContentBuilder.Create().WithMust(must)
	if element.HasNot() {
		notLines := element.Not()
		not, err := app.bestLineToMatches(notLines, dependencies, script, chansAreActivated)
		if err != nil {
			return nil, err
		}

		if not != nil {
			builder.WithNot(not)
		}
	}

	return builder.Now()
}

func (app *adapter) bestLineToMatches(lines []paths.Line, dependencies paths.Dependencies, script string, chansAreActivated bool) (LineMatch, error) {
	biggestValidDiscoveries := 0
	biggestAmountOfInvalidMatches := 0
	invalidMatchIndex := -1
	var invalidMatches Matches
	validMatchIndex := -1
	var validMatches Matches
	for index, oneLine := range lines {
		matches, err := app.lineToMatches(oneLine, dependencies, script, chansAreActivated)
		if err != nil {
			return nil, err
		}

		if matches == nil {
			continue
		}

		if !matches.IsValid() {
			amount := len(matches.BlockMatches())
			if amount > biggestAmountOfInvalidMatches {
				invalidMatches = matches
				invalidMatchIndex = index
			}

			continue
		}

		discoveriesLength := len(matches.Discoveries())
		if discoveriesLength > biggestValidDiscoveries {
			validMatchIndex = index
			validMatches = matches
			biggestValidDiscoveries = discoveriesLength
		}
	}

	if validMatches != nil {
		return app.lineMatchBuilder.Create().WithIndex(uint(validMatchIndex)).WithMatches(validMatches).Now()
	}

	if invalidMatches == nil {
		return nil, nil
	}

	return app.lineMatchBuilder.Create().WithIndex(uint(invalidMatchIndex)).WithMatches(invalidMatches).Now()
}

func (app *adapter) findMatchesInChannels(script string) (LineMatch, error) {
	for _, oneChannel := range app.channels {
		token, err := app.elementToToken(oneChannel.Element(), oneChannel.Dependencies(), script, false)
		if err != nil {
			return nil, err
		}

		// if the token is invalid:
		if !token.IsValid() {
			continue
		}

		// the channel is a match
		return token.Content().Must(), nil
	}

	return nil, nil
}

func (app *adapter) lineToMatches(line paths.Line, dependencies paths.Dependencies, script string, chansAreActivated bool) (Matches, error) {
	matches := []BlockMatch{}
	instructions := line.Instructions()
	remainingScript := script
	nextChansAreActivated := chansAreActivated
	for _, oneInstruction := range instructions {
		if oneInstruction.IsChannelSwitch() {
			nextChansAreActivated = !nextChansAreActivated
			continue
		}

		currentContainer := oneInstruction.Container()
		match, err := app.containerToBlockMatch(currentContainer, dependencies, remainingScript, chansAreActivated, nextChansAreActivated)
		if err != nil {
			return nil, err
		}

		if match == nil {
			break
		}

		// if the match is invalid:
		if !match.IsValid() {
			matches = []BlockMatch{
				match,
			}

			break
		}

		discoveries := match.Discoveries()
		if !strings.HasPrefix(remainingScript, discoveries) {
			return nil, nil
		}

		remainingScript = strings.Replace(remainingScript, discoveries, "", 1)
		matches = append(matches, match)

		if !nextChansAreActivated {
			chansAreActivated = false
		}
	}

	if len(matches) <= 0 {
		return nil, nil
	}

	return app.matchesBuilder.Create().WithList(matches).Now()
}

func (app *adapter) containerToBlockMatch(container paths.Container, dependencies paths.Dependencies, script string, chansAreActivated bool, nextChansAreActivated bool) (BlockMatch, error) {
	if script == "" {
		// if the cardinality of the token is optional, do not build a NextElement
		cardinality := container.Cardinality()
		if cardinality.IsValid(0) {
			return nil, nil
		}

		nextElement, err := app.buildNextElement(container, dependencies)
		if err != nil {
			return nil, err
		}

		return app.blockMatchBuilder.Create().WithNextElement(nextElement).Now()
	}

	if chansAreActivated {
		chanBlockMatch, err := app.findChannelBlockMatch(container, dependencies, script, nextChansAreActivated)
		if err != nil {
			return nil, err
		}

		if chanBlockMatch != nil {
			return chanBlockMatch, nil
		}
	}

	match, err := app.containerToMatch(container, dependencies, script, nextChansAreActivated)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, nil
	}

	return app.blockMatchBuilder.Create().WithContainer(match).Now()
}

func (app *adapter) buildNextElement(container paths.Container, dep paths.Dependencies) (NextElement, error) {
	if container.IsRule() {
		parsedRule := container.Rule().Base()
		return app.buildNextElementFromRule(parsedRule)
	}

	if container.IsToken() {
		name := container.Token().Path().Element().Name()
		return app.buildNextElementFromTokenName(name, dep)
	}

	recursiveTokenName := container.Recursive().Name()
	return app.buildNextElementFromTokenName(recursiveTokenName, dep)
}

func (app *adapter) buildNextElementFromTokenName(name string, dep paths.Dependencies) (NextElement, error) {
	rule, err := app.fetchFirstRuleFromTokenName(name, dep)
	if err != nil {
		return nil, err
	}

	return app.buildNextElementFromRule(rule)
}

func (app *adapter) fetchFirstRuleFromTokenName(name string, dep paths.Dependencies) (rules.Rule, error) {
	tokenPath, err := dep.Fetch(name)
	if err != nil {
		return nil, err
	}

	mustLines := tokenPath.Must()
	firstContainer, err := app.fetchFirstContainerFromLine(mustLines[0])
	if err != nil {
		return nil, err
	}

	if firstContainer.IsToken() {
		nextToken := firstContainer.Token().Path().Element().Name()
		return app.fetchFirstRuleFromTokenName(nextToken, dep)
	}

	if firstContainer.IsRule() {
		return firstContainer.Rule().Base(), nil
	}

	recursiveToken := firstContainer.Recursive().Name()
	return app.fetchFirstRuleFromTokenName(recursiveToken, dep)

}
func (app *adapter) fetchFirstContainerFromLine(line paths.Line) (paths.Container, error) {
	instructions := line.Instructions()
	for _, oneInstruction := range instructions {
		if oneInstruction.IsContainer() {
			return oneInstruction.Container(), nil
		}
	}

	return nil, nil
}

func (app *adapter) buildNextElementFromRule(rule rules.Rule) (NextElement, error) {
	builder := app.nextElementBuilder.Create()
	content := rule.Element().Content()
	if content.IsConstant() {
		constant := content.Constant()
		builder.WithConstant(constant)
	}

	if content.IsPattern() {
		subPattern := content.Pattern().SubPatterns().First()
		builder.WithPattern(subPattern)
	}

	return builder.Now()
}

func (app *adapter) findChannelBlockMatch(container paths.Container, dependencies paths.Dependencies, script string, nextChansAreActivated bool) (BlockMatch, error) {
	chanMatches, err := app.findMatchesInChannels(script)
	if err != nil {
		return nil, err
	}

	if chanMatches == nil {
		return nil, nil
	}

	matchesScript := chanMatches.Matches().Discoveries()
	if !strings.HasPrefix(script, matchesScript) {
		return nil, nil
	}

	remainingScript := strings.Replace(script, matchesScript, "", 1)
	blockMatch, err := app.containerToBlockMatch(container, dependencies, remainingScript, nextChansAreActivated, nextChansAreActivated)
	if err != nil {
		return nil, err
	}

	if blockMatch == nil {
		return nil, nil
	}

	return app.blockMatchBuilder.Create().WithBlock(blockMatch).WithChannelPrefix(chanMatches).Now()
}

func (app *adapter) containerToMatch(container paths.Container, dependencies paths.Dependencies, script string, chansAreActivated bool) (Match, error) {
	builder := app.matchBuilder.Create()
	if container.IsLines() {
		lines := container.Lines()
		lineMatch, err := app.bestLineToMatches(lines, dependencies, script, chansAreActivated)
		if err != nil {
			return nil, err
		}

		if lineMatch == nil {
			return nil, nil
		}

		builder.WithLine(lineMatch)
	}

	if container.IsToken() {
		token := container.Token()
		tokenMatch, err := app.tokenToTokenMatch(token, dependencies, script, chansAreActivated)
		if err != nil {
			return nil, err
		}

		// if the token is invalid, but the cardinality can be 0, return nil:
		isSkip := tokenMatch.Result().Cardinality().IsValid(0) && !tokenMatch.IsValid()
		if isSkip {
			skipMatch, err := app.skipMatchBuilder.Create().WithToken(tokenMatch).Now()
			if err != nil {
				return nil, err
			}

			builder.WithSkip(skipMatch)
		}

		if !isSkip {
			builder.WithToken(tokenMatch)
		}
	}

	if container.IsRule() {
		rule := container.Rule()
		ruleMatch, err := app.ruleToRuleMatch(rule, script, chansAreActivated)
		if err != nil {
			return nil, err
		}

		// if the rule is invalid, but the cardinality can be 0, return nil:
		isSkip := ruleMatch.Result().Path().Cardinality().IsValid(0) && !ruleMatch.IsValid()
		if isSkip {
			skipMatch, err := app.skipMatchBuilder.Create().WithRule(ruleMatch).Now()
			if err != nil {
				return nil, err
			}

			builder.WithSkip(skipMatch)
		}

		if !isSkip {
			builder.WithRule(ruleMatch)
		}
	}

	if container.IsRecursive() {
		recursiveToken := container.Recursive()
		name := recursiveToken.Name()
		if previousScript, ok := app.recursiveTokens[name]; ok {
			if previousScript == script {
				delete(app.recursiveTokens, name)
				return nil, nil
			}
		}

		app.recursiveTokens[name] = script
		element, err := dependencies.Fetch(name)
		if err != nil {
			str := fmt.Sprintf("the recursive token (name: %s) could not be found: %s", name, err.Error())
			return nil, errors.New(str)
		}

		tokenPathBuilder := app.pathTokenPathBuilder.Create().WithElement(element)
		if recursiveToken.HasSpecifiers() {
			specifiers := recursiveToken.Specifiers()
			tokenPathBuilder.WithSpecifiers(specifiers)
		}

		tokenPath, err := tokenPathBuilder.Now()
		if err != nil {
			return nil, err
		}

		cardinality := recursiveToken.Cardinality()
		token, err := app.pathTokenBuilder.Create().WithPath(tokenPath).WithCardinality(cardinality).Now()
		if err != nil {
			return nil, err
		}

		tokenMatch, err := app.tokenToTokenMatch(token, dependencies, script, chansAreActivated)
		if err != nil {
			return nil, err
		}

		// if the token invalid, but the cardinality can be 0, add a skip match:
		isSkip := tokenMatch.Result().Cardinality().IsValid(0) && !tokenMatch.IsValid()
		if isSkip {
			skipMatch, err := app.skipMatchBuilder.Create().WithToken(tokenMatch).Now()
			if err != nil {
				return nil, err
			}

			builder.WithSkip(skipMatch)
		}

		if !isSkip {
			builder.WithToken(tokenMatch)
		}

		return builder.Now()
	}

	return builder.Now()
}

func (app *adapter) tokenToTokenMatch(token paths.Token, dependencies paths.Dependencies, script string, chansAreActivated bool) (TokenMatch, error) {
	result, err := app.tokenToTokenResult(token, dependencies, script, chansAreActivated)
	if err != nil {
		return nil, err
	}

	path := token.Path()
	builder := app.tokenMatchBuilder.Create().WithPath(path).WithResult(result)
	if path.HasSpecifiers() {
		pathSpecifiers := path.Specifiers()
		specifiers, err := app.toSpecifiers(pathSpecifiers, result)
		if err != nil {
			return nil, err
		}

		builder.WithSpecifiers(specifiers)
	}

	return builder.Now()
}

func (app *adapter) toSpecifiers(specifics []paths.Specifier, result TokenResult) ([]Specifier, error) {
	out := []Specifier{}
	for _, oneSpecific := range specifics {
		containerName := oneSpecific.ContainerName()
		cardinality := oneSpecific.Cardinality()
		specifier, err := app.specifierBuilder.Create().WithContainerName(containerName).WithCardinality(cardinality).WithResult(result).Now()
		if err != nil {
			return nil, err
		}

		out = append(out, specifier)
	}

	return out, nil
}

func (app *adapter) tokenToTokenResult(token paths.Token, dependencies paths.Dependencies, script string, chansAreActivated bool) (TokenResult, error) {
	tokenPath := token.Path()
	cardinality := token.Cardinality()
	matches, err := app.tokenPathToTokenResultMatches(tokenPath, dependencies, script, cardinality, chansAreActivated)
	if err != nil {
		return nil, err
	}

	builder := app.tokenResultBuilder.Create().WithInput(script).WithCardinality(cardinality)
	if matches != nil {
		builder.WithMatches(matches)
	}

	return builder.Now()
}

func (app *adapter) tokenPathToTokenResultMatches(
	tokenPath paths.TokenPath,
	dependencies paths.Dependencies,
	script string,
	cardinality cardinality.Cardinality,
	chansAreActivated bool,
) (TokenResultMatches, error) {
	results := []Token{}
	element := tokenPath.Element()
	remainingScript := script
	for {
		result, err := app.elementToToken(element, dependencies, remainingScript, chansAreActivated)
		if err != nil {
			return nil, err
		}

		if !result.IsValid() {
			break
		}

		// add the result:
		results = append(results, result)

		// remove the script, from the remaining script:
		resultScript := result.Content().Must().Matches().Discoveries()
		if !strings.HasPrefix(remainingScript, resultScript) {
			break
		}

		remainingScript = strings.Replace(remainingScript, resultScript, "", 1)

		// if we reached the max amount, if any, break:
		_, pMax := cardinality.Delimiter()
		if pMax != nil {
			amountResults := uint(len(results))
			if amountResults >= *pMax {
				break
			}
		}
	}

	if len(results) <= 0 {
		return nil, nil
	}

	return app.tokenResultMatchesBuilder.Create().WithResults(results).Now()
}

func (app *adapter) ruleToRuleMatch(rulePath paths.Rule, script string, chansAreActivated bool) (RuleMatch, error) {
	remaining := script
	rule := rulePath.Base()
	scriptLength := len(script)
	list := []results.Result{}
	canContainPrefix := app.canContainPrefix && chansAreActivated
	_, pMax := rulePath.Cardinality().Delimiter()
	for {
		res, err := app.resultAdapter.ToResult(rule, remaining)
		if err != nil {
			break
		}

		if !res.HasMatches() {
			break
		}

		list = append(list, res)
		content := res.Content()

		if res.HasMatches() {
			matches := res.Matches().List()
			index := matches[0].Discoveries().Index()
			if index > 0 {
				app.canContainPrefix = false
			}

			if index > 0 && !canContainPrefix {
				break
			}
		}

		contentLength := len(content)
		if scriptLength <= contentLength {
			break
		}

		remaining = remaining[contentLength:]
		if pMax != nil && uint(len(list)) >= *pMax {
			break
		}
	}

	result, err := app.resultBuilder.Create().WithResults(list).Now()
	if err != nil {
		return nil, err
	}

	ruleMatchResult, err := app.ruleMatchResultBuilder.Create().WithResult(result).WithPath(rulePath).Now()
	if err != nil {
		return nil, err
	}

	builder := app.ruleMatchBuilder.Create().WithRule(rulePath).WithResult(ruleMatchResult)
	if app.canContainPrefix && chansAreActivated {
		builder.CanContainPrefix()
		//app.canContainPrefix = false
	}

	return builder.Now()
}
