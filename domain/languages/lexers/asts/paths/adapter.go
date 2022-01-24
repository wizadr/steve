package paths

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

type adapter struct {
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
	elementReplaceFn       FetchElementReplacementFn
	dependencies           []Element
	currentTokenStack      map[string]bool
}

func createAdapter(
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
	tokens tokens.Tokens,
) Adapter {
	return createAdapterInternally(
		builder,
		elementBuilder,
		dependenciesBuilder,
		lineBuilder,
		instructionBuilder,
		containerBuilder,
		rescursiveTokenBuilder,
		tokenBuilder,
		ruleBuilder,
		tokenPathBuilder,
		specifierBuilder,
		cardinalityBuilder,
		tokens,
		nil,
		nil,
	)
}

func createAdapterWithTokenNotFoundFunc(
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
	tokens tokens.Tokens,
	tokenNotFoundFn FetchTokenNotFoundFn,
) Adapter {
	return createAdapterInternally(
		builder,
		elementBuilder,
		dependenciesBuilder,
		lineBuilder,
		instructionBuilder,
		containerBuilder,
		rescursiveTokenBuilder,
		tokenBuilder,
		ruleBuilder,
		tokenPathBuilder,
		specifierBuilder,
		cardinalityBuilder,
		tokens,
		tokenNotFoundFn,
		nil,
	)
}

func createAdapterWithTokenReplaceFunc(
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
	tokens tokens.Tokens,
	elementReplaceFn FetchElementReplacementFn,
) Adapter {
	return createAdapterInternally(
		builder,
		elementBuilder,
		dependenciesBuilder,
		lineBuilder,
		instructionBuilder,
		containerBuilder,
		rescursiveTokenBuilder,
		tokenBuilder,
		ruleBuilder,
		tokenPathBuilder,
		specifierBuilder,
		cardinalityBuilder,
		tokens,
		nil,
		elementReplaceFn,
	)
}

func createAdapterWithTokenNotFoundAndTokenReplaceFunc(
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
	tokens tokens.Tokens,
	tokenNotFoundFn FetchTokenNotFoundFn,
	elementReplaceFn FetchElementReplacementFn,
) Adapter {
	return createAdapterInternally(
		builder,
		elementBuilder,
		dependenciesBuilder,
		lineBuilder,
		instructionBuilder,
		containerBuilder,
		rescursiveTokenBuilder,
		tokenBuilder,
		ruleBuilder,
		tokenPathBuilder,
		specifierBuilder,
		cardinalityBuilder,
		tokens,
		tokenNotFoundFn,
		elementReplaceFn,
	)
}

func createAdapterInternally(
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
	tokens tokens.Tokens,
	tokenNotFoundFn FetchTokenNotFoundFn,
	elementReplaceFn FetchElementReplacementFn,
) Adapter {
	out := adapter{
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
		tokens:                 tokens,
		tokenNotFoundFn:        tokenNotFoundFn,
		elementReplaceFn:       elementReplaceFn,
		dependencies:           nil,
		currentTokenStack:      map[string]bool{},
	}

	return &out
}

// ToPath converts a Token to Path
func (app *adapter) ToPath(token tokens.Token) (Path, error) {
	if app.dependencies == nil {
		dep, err := app.buildDependencies()
		if err != nil {
			return nil, err
		}

		app.dependencies = dep
	}

	app.currentTokenStack = map[string]bool{}
	element, externalDepElements, err := app.toElement(token)
	if err != nil {
		return nil, err
	}

	allElements := app.dependencies
	for _, oneElement := range externalDepElements {
		allElements = append(allElements, oneElement)
	}

	allElements = append(allElements, element)
	dep, err := app.dependenciesBuilder.Create().WithDependencies(allElements).Now()
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithElement(element).WithDependencies(dep).Now()
}

func (app *adapter) buildDependencies() ([]Element, error) {
	dependencyList := []Element{}
	allTokens := app.tokens.All()
	for _, oneToken := range allTokens {
		dep, extDepElements, err := app.toElement(oneToken)
		if err != nil {
			return nil, err
		}

		dependencyList = append(dependencyList, dep)
		dependencyList = append(dependencyList, extDepElements...)
	}

	return dependencyList, nil
}

func (app *adapter) toElement(token tokens.Token) (Element, []Element, error) {
	name := token.Name()
	block := token.Block()
	externalDepElements := []Element{}
	mustProcessedLines := block.Must().All()
	mustLines, mustExtElements, err := app.processLines(name, mustProcessedLines)
	if err != nil {
		return nil, nil, err
	}

	externalDepElements = append(externalDepElements, mustExtElements...)
	builder := app.elementBuilder.Create().WithName(name).WithMust(mustLines)
	if block.HasNot() {
		notProcessedLines := block.Not().All()
		notLines, notExtElements, err := app.processLines(name, notProcessedLines)
		if err != nil {
			return nil, nil, err
		}

		externalDepElements = append(externalDepElements, notExtElements...)
		builder.WithNot(notLines)
	}

	element, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return element, externalDepElements, nil
}

func (app *adapter) toElementEncapsulated(token tokens.Token) (Element, []Element, error) {
	name := token.Name()
	app.currentTokenStack[name] = true
	rulePath, extDepElements, err := app.toElement(token)
	if _, ok := app.currentTokenStack[name]; ok {
		delete(app.currentTokenStack, name)
	}

	return rulePath, extDepElements, err
}

func (app *adapter) processLines(parentTokenName string, parsedLines []tokens.Line) ([]Line, []Element, error) {
	lines := []Line{}
	extDepElements := []Element{}
	for index, oneParsedLine := range parsedLines {
		line, lineExtDepElements, err := app.processLine(parentTokenName, oneParsedLine, uint(index), parsedLines)
		if err != nil {
			return nil, nil, err
		}

		if line == nil {
			continue
		}

		lines = append(lines, line)
		extDepElements = append(extDepElements, lineExtDepElements...)
	}

	return lines, extDepElements, nil
}

func (app *adapter) processLine(parentTokenName string, line tokens.Line, lineIndex uint, blockLines []tokens.Line) (Line, []Element, error) {
	amountElements := 0
	lineInstructions := line.Instructions()
	for _, oneInstruction := range lineInstructions {
		if !oneInstruction.IsElement() {
			continue
		}

		amountElements++
	}

	instructions := []Instruction{}
	extDepElements := []Element{}
	for _, oneInstruction := range lineInstructions {
		builder := app.instructionBuilder.Create()
		if oneInstruction.IsChannelSwitch() {
			builder.IsChannelSwitch()
		} else {
			oneElement := oneInstruction.Element()
			container, containerDepElements, err := app.elementToContainer(parentTokenName, oneElement, uint(amountElements), lineIndex, blockLines)
			if err != nil {
				return nil, nil, err
			}

			if container == nil {
				return nil, nil, nil
			}

			if containerDepElements != nil {
				extDepElements = append(extDepElements, containerDepElements...)
			}

			builder.WithContainer(container)
		}

		ins, err := builder.Now()
		if err != nil {
			return nil, nil, err
		}

		instructions = append(instructions, ins)
	}

	procLined, err := app.lineBuilder.Create().WithInstructions(instructions).Now()
	if err != nil {
		return nil, nil, err
	}

	return procLined, extDepElements, nil
}

func (app *adapter) elementToContainer(parentTokenName string, element tokens.Element, amountElementsCurrentLine uint, lineIndex uint, blockLines []tokens.Line) (Container, []Element, error) {
	// if there is no cardinality (default), it means the cardinality is mandatory:
	cardinality := element.Cardinality()
	if !element.HasCardinality() {
		card, err := app.cardinalityBuilder.Create().IsMandatory().Now()
		if err != nil {
			return nil, nil, err
		}

		cardinality = card
	}

	content := element.Content()
	builder := app.containerBuilder.Create()
	if content.IsToken() {
		tokenReference := content.Token()
		tokenReferenceName := tokenReference.Name()

		// if the token is recursive:
		if _, ok := app.currentTokenStack[tokenReferenceName]; ok {
			if parentTokenName == tokenReferenceName {
				// if there is just 1 element in the line, that means we have a recursivity without exit possibility, return an error:
				if amountElementsCurrentLine <= 1 {
					str := fmt.Sprintf("the token (name: %s) containers a recursivity occurence without exit possibility", tokenReferenceName)
					return nil, nil, errors.New(str)
				}
			}

			recursiveTokenBuilder := app.rescursiveTokenBuilder.Create().WithName(tokenReferenceName).WithCardinality(cardinality)
			if tokenReference.HasSpecifiers() {
				parsedSpecifiers := tokenReference.Specifiers()
				specifiers, err := app.buildSpecifiers(parsedSpecifiers)
				if err != nil {
					return nil, nil, err
				}

				recursiveTokenBuilder.WithSpecifiers(specifiers)
			}

			recursiveToken, err := recursiveTokenBuilder.Now()
			if err != nil {
				return nil, nil, err
			}

			container, err := builder.WithRecursive(recursiveToken).Now()
			if err != nil {
				return nil, nil, err
			}

			return container, nil, nil
		}

		// build the token:
		tokenPath, path, err := app.tokenReferenceToTokenPath(tokenReference)
		if err != nil {
			return nil, nil, err
		}

		token, err := app.tokenBuilder.Create().WithCardinality(cardinality).WithPath(tokenPath).Now()
		if err != nil {
			return nil, nil, err
		}

		container, err := builder.WithToken(token).Now()
		if err != nil {
			return nil, nil, err
		}

		return container, path, nil
	}

	if content.IsRule() {
		parsedRule := content.Rule()
		rule, err := app.ruleBuilder.Create().WithCardinality(cardinality).WithBase(parsedRule).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithRule(rule)
	}

	container, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return container, nil, nil
}

func (app *adapter) tokenReferenceToTokenPath(tokenReference tokens.TokenReference) (TokenPath, []Element, error) {
	extDepElements := []Element{}
	name := tokenReference.Name()
	builder := app.tokenPathBuilder.Create()

	var fetchedPath Path
	hasFetchdPath := false
	if app.elementReplaceFn != nil {
		fetchedPath, refExtDep := app.elementReplaceFn(name)
		if fetchedPath != nil {
			hasFetchdPath = true
			extDepElements = append(extDepElements, refExtDep...)
			builder.WithElement(fetchedPath)
		}
	}

	if !hasFetchdPath {
		referencedToken, err := app.tokens.Find(name)
		if err != nil {
			if app.tokenNotFoundFn == nil {
				return nil, nil, err
			}

			fetchedPath = app.tokenNotFoundFn(name)
			if fetchedPath != nil {
				hasFetchdPath = true
				element := fetchedPath.Element()
				extDepElements = append(extDepElements, fetchedPath.CombinedElements()...)
				builder.WithElement(element)
			}
		}

		if !hasFetchdPath {
			if referencedToken == nil {
				str := fmt.Sprintf("the referenced token (name: %s) is not declared", name)
				return nil, nil, errors.New(str)
			}

			element, refExtElements, err := app.toElementEncapsulated(referencedToken)
			if err != nil {
				return nil, nil, err
			}

			extDepElements = append(extDepElements, refExtElements...)
			builder.WithElement(element)
		}
	}

	if tokenReference.HasSpecifiers() {
		parsedSpecifiers := tokenReference.Specifiers()
		specifiers, err := app.buildSpecifiers(parsedSpecifiers)
		if err != nil {
			return nil, nil, err
		}

		builder.WithSpecifiers(specifiers)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, extDepElements, nil
}

func (app *adapter) buildSpecifiers(specifiers tokens.TokenSpecifiers) ([]Specifier, error) {
	out := []Specifier{}
	list := specifiers.All()
	for _, oneSpecifier := range list {
		cardinality := oneSpecifier.Cardinality()
		containerName := oneSpecifier.Content().Name()
		specifier, err := app.specifierBuilder.Create().WithCardinality(cardinality).WithContainerName(containerName).Now()
		if err != nil {
			return nil, err
		}

		out = append(out, specifier)
	}

	return out, nil
}
