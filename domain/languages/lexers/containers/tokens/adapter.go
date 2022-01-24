package tokens

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

type adapter struct {
	builder                        Builder
	tokenBuilder                   TokenBuilder
	testSuiteBuilder               TestSuiteBuilder
	linesBuilder                   LinesBuilder
	lineBuilder                    LineBuilder
	blockBuilder                   BlockBuilder
	instructionBuilder             InstructionBuilder
	elementBuilder                 ElementBuilder
	cardinalityAdapter             cardinality.Adapter
	specificAdapter                cardinality.SpecificAdapter
	contentBuilder                 ContentBuilder
	tokenReferenceBuilder          TokenReferenceBuilder
	tokenSpecifiersBuilder         TokenSpecifiersBuilder
	tokenSpecifierBuilder          TokenSpecifierBuilder
	tokenSpecifierContentBuilder   TokenSpecifierContentBuilder
	tokenPattern                   string
	rulePattern                    string
	anythingExcept                 string
	begin                          string
	or                             string
	end                            string
	notDelimiter                   string
	testDelimiter                  string
	testLineBegin                  string
	testLineDelimiter              string
	testLineEnd                    string
	whiteSpacePattern              string
	tokenSpecifierPrefix           string
	tokenSpecifierSuffix           string
	cardinalityZeroMultiplePattern string
	cardinalityMultiplePattern     string
	cardinalityOptionalPattern     string
	cardinalityRangeBegin          string
	cardinalityRangeEnd            string
	cardinalityRangeSeparator      string
	switchChannelCharacter         string
	rules                          rules.Rules
}

func createAdapter(
	builder Builder,
	tokenBuilder TokenBuilder,
	testSuiteBuilder TestSuiteBuilder,
	linesBuilder LinesBuilder,
	lineBuilder LineBuilder,
	blockBuilder BlockBuilder,
	instructionBuilder InstructionBuilder,
	elementBuilder ElementBuilder,
	cardinalityAdapter cardinality.Adapter,
	specificAdapter cardinality.SpecificAdapter,
	contentBuilder ContentBuilder,
	tokenReferenceBuilder TokenReferenceBuilder,
	tokenSpecifiersBuilder TokenSpecifiersBuilder,
	tokenSpecifierBuilder TokenSpecifierBuilder,
	tokenSpecifierContentBuilder TokenSpecifierContentBuilder,
	tokenPattern string,
	rulePattern string,
	anythingExcept string,
	begin string,
	or string,
	end string,
	notDelimiter string,
	testDelimiter string,
	testLineBegin string,
	testLineDelimiter string,
	testLineEnd string,
	whiteSpacePattern string,
	tokenSpecifierPrefix string,
	tokenSpecifierSuffix string,
	cardinalityZeroMultiplePattern string,
	cardinalityMultiplePattern string,
	cardinalityOptionalPattern string,
	cardinalityRangeBegin string,
	cardinalityRangeEnd string,
	cardinalityRangeSeparator string,
	switchChannelCharacter string,
	rules rules.Rules,
) Adapter {
	return createAdapterInternally(
		builder,
		tokenBuilder,
		testSuiteBuilder,
		linesBuilder,
		lineBuilder,
		blockBuilder,
		instructionBuilder,
		elementBuilder,
		cardinalityAdapter,
		specificAdapter,
		contentBuilder,
		tokenReferenceBuilder,
		tokenSpecifiersBuilder,
		tokenSpecifierBuilder,
		tokenSpecifierContentBuilder,
		tokenPattern,
		rulePattern,
		anythingExcept,
		begin,
		or,
		end,
		notDelimiter,
		testDelimiter,
		testLineBegin,
		testLineDelimiter,
		testLineEnd,
		whiteSpacePattern,
		tokenSpecifierPrefix,
		tokenSpecifierSuffix,
		cardinalityZeroMultiplePattern,
		cardinalityMultiplePattern,
		cardinalityOptionalPattern,
		cardinalityRangeBegin,
		cardinalityRangeEnd,
		cardinalityRangeSeparator,
		switchChannelCharacter,
		rules,
	)
}

func createAdapterInternally(
	builder Builder,
	tokenBuilder TokenBuilder,
	testSuiteBuilder TestSuiteBuilder,
	linesBuilder LinesBuilder,
	lineBuilder LineBuilder,
	blockBuilder BlockBuilder,
	instructionBuilder InstructionBuilder,
	elementBuilder ElementBuilder,
	cardinalityAdapter cardinality.Adapter,
	specificAdapter cardinality.SpecificAdapter,
	contentBuilder ContentBuilder,
	tokenReferenceBuilder TokenReferenceBuilder,
	tokenSpecifiersBuilder TokenSpecifiersBuilder,
	tokenSpecifierBuilder TokenSpecifierBuilder,
	tokenSpecifierContentBuilder TokenSpecifierContentBuilder,
	tokenPattern string,
	rulePattern string,
	anythingExcept string,
	begin string,
	or string,
	end string,
	notDelimiter string,
	testDelimiter string,
	testLineBegin string,
	testLineDelimiter string,
	testLineEnd string,
	whiteSpacePattern string,
	tokenSpecifierPrefix string,
	tokenSpecifierSuffix string,
	cardinalityZeroMultiplePattern string,
	cardinalityMultiplePattern string,
	cardinalityOptionalPattern string,
	cardinalityRangeBegin string,
	cardinalityRangeEnd string,
	cardinalityRangeSeparator string,
	switchChannelCharacter string,
	rules rules.Rules,
) Adapter {
	out := adapter{
		builder:                        builder,
		tokenBuilder:                   tokenBuilder,
		testSuiteBuilder:               testSuiteBuilder,
		rulePattern:                    rulePattern,
		linesBuilder:                   linesBuilder,
		lineBuilder:                    lineBuilder,
		blockBuilder:                   blockBuilder,
		instructionBuilder:             instructionBuilder,
		elementBuilder:                 elementBuilder,
		cardinalityAdapter:             cardinalityAdapter,
		specificAdapter:                specificAdapter,
		contentBuilder:                 contentBuilder,
		tokenReferenceBuilder:          tokenReferenceBuilder,
		tokenSpecifiersBuilder:         tokenSpecifiersBuilder,
		tokenSpecifierBuilder:          tokenSpecifierBuilder,
		tokenSpecifierContentBuilder:   tokenSpecifierContentBuilder,
		tokenPattern:                   tokenPattern,
		anythingExcept:                 anythingExcept,
		begin:                          begin,
		or:                             or,
		end:                            end,
		notDelimiter:                   notDelimiter,
		testDelimiter:                  testDelimiter,
		testLineBegin:                  testLineBegin,
		testLineDelimiter:              testLineDelimiter,
		testLineEnd:                    testLineEnd,
		whiteSpacePattern:              whiteSpacePattern,
		tokenSpecifierPrefix:           tokenSpecifierPrefix,
		tokenSpecifierSuffix:           tokenSpecifierSuffix,
		cardinalityZeroMultiplePattern: cardinalityZeroMultiplePattern,
		cardinalityMultiplePattern:     cardinalityMultiplePattern,
		cardinalityOptionalPattern:     cardinalityOptionalPattern,
		cardinalityRangeBegin:          cardinalityRangeBegin,
		cardinalityRangeEnd:            cardinalityRangeEnd,
		cardinalityRangeSeparator:      cardinalityRangeSeparator,
		switchChannelCharacter:         switchChannelCharacter,
		rules:                          rules,
	}

	return &out
}

// ToTokens converts content to a Tokens instance
func (app *adapter) ToTokens(content string) (Tokens, error) {
	list, err := app.tokens(content)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithTokens(list).Now()
}

func (app *adapter) tokens(content string) ([]Token, error) {
	// find the tokens
	anythingExceptEnd := fmt.Sprintf(app.anythingExcept, app.end)
	patternStr := fmt.Sprintf(
		"(%s(%s)%s%s%s(%s)%s)",
		app.whiteSpacePattern,
		app.tokenPattern,
		app.whiteSpacePattern,
		app.begin,
		app.whiteSpacePattern,
		anythingExceptEnd,
		app.end,
	)

	tokens := []Token{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		name := oneMatch[2]
		tokenBuilder := app.tokenBuilder.Create().WithName(name)
		testSuite, blockCode, err := app.testSuite(name, oneMatch[3])
		if err != nil {
			return nil, err
		}

		if testSuite != nil {
			tokenBuilder.WithTestSuite(testSuite)
		}

		block, err := app.block(name, blockCode)
		if err != nil {
			return nil, err
		}

		token, err := tokenBuilder.WithBlock(block).Now()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (app *adapter) testSuite(tokenName string, content string) (TestSuite, string, error) {
	sections := strings.Split(content, app.testDelimiter)
	if len(sections) <= 1 {
		return nil, sections[0], nil
	}

	trimmedContent := strings.TrimSpace(sections[1])
	if !strings.HasPrefix(trimmedContent, app.begin) {
		str := fmt.Sprintf("the testSuite content (%s) of Token (%s) was expecting the begin (%s) character as prefix", sections[1], tokenName, app.begin)
		return nil, sections[0], errors.New(str)
	}

	linesContent := strings.Split(trimmedContent[1:], app.testLineDelimiter)
	if len(linesContent) <= 0 {
		str := fmt.Sprintf("the testSuite content (%s) of Token (%s) was expected to be separated by line using this delimiter: %s", trimmedContent[1:], tokenName, app.testLineDelimiter)
		return nil, sections[0], errors.New(str)
	}

	lines := []string{}
	for _, oneLineContent := range linesContent {
		if oneLineContent == "" {
			continue
		}

		lines = append(lines, oneLineContent)
	}

	suite, err := app.testSuiteBuilder.Create().WithLines(lines).Now()
	if err != nil {
		return nil, sections[0], err
	}

	return suite, sections[0], nil
}

func (app *adapter) block(tokenName string, content string) (Block, error) {
	sections := strings.Split(content, app.notDelimiter)
	if len(sections) <= 0 {
		str := fmt.Sprintf("there must be at least must Lines in order to create the Block instance of Token: %s", tokenName)
		return nil, errors.New(str)
	}

	must, err := app.lines(sections[0])
	if err != nil {
		return nil, err
	}

	builder := app.blockBuilder.Create().WithMust(must)
	if len(sections) > 1 {
		trimmedContent := strings.TrimSpace(sections[1])
		if !strings.HasPrefix(trimmedContent, app.begin) {
			str := fmt.Sprintf("the not content (%s) of Token (%s) was expecting the begin (%s) character as prefix", sections[1], tokenName, app.begin)
			return nil, errors.New(str)
		}

		notLines, err := app.lines(trimmedContent[1:])
		if err != nil {
			return nil, err
		}

		builder.WithNot(notLines)
	}

	return builder.Now()
}

func (app *adapter) lines(content string) (Lines, error) {
	list := []Line{}
	sections := strings.Split(content, app.or)
	for _, oneSection := range sections {
		line, err := app.line(oneSection)
		if err != nil {
			return nil, err
		}

		list = append(list, line)
	}

	return app.linesBuilder.Create().WithLines(list).Now()
}

func (app *adapter) line(content string) (Line, error) {
	instructions, err := app.instructions(content)
	if err != nil {
		return nil, err
	}

	return app.lineBuilder.Create().WithInstructions(instructions).Now()
}

func (app *adapter) instructions(content string) ([]Instruction, error) {
	cardinalityRangePattern := app.cardinalityRangePattern()
	cardinalityPattern := fmt.Sprintf(
		"(%s)?(%s)?(%s)?(%s)?",
		cardinalityRangePattern,
		app.cardinalityMultiplePattern,
		app.cardinalityZeroMultiplePattern,
		app.cardinalityOptionalPattern,
	)

	anythingExceptSubElementSuffix := fmt.Sprintf(app.anythingExcept, app.tokenSpecifierSuffix)
	subElementPattern := fmt.Sprintf(
		"%s(%s)%s",
		app.tokenSpecifierPrefix,
		anythingExceptSubElementSuffix,
		app.tokenSpecifierSuffix,
	)

	elementSuffixPattern := fmt.Sprintf(
		"(%s)?(%s)?",
		cardinalityPattern,
		subElementPattern,
	)

	tokenPattern := fmt.Sprintf(
		"(%s)%s(%s)",
		app.tokenPattern,
		app.whiteSpacePattern,
		elementSuffixPattern,
	)

	rulePattern := fmt.Sprintf(
		"(%s)%s(%s)",
		app.rulePattern,
		app.whiteSpacePattern,
		elementSuffixPattern,
	)

	patternStr := fmt.Sprintf(
		"(%s)?(%s)?(%s)?",
		tokenPattern,
		rulePattern,
		app.switchChannelCharacter,
	)

	out := []Instruction{}
	pattern := regexp.MustCompile(patternStr)
	trimmedContent := strings.TrimSpace(content)
	matches := pattern.FindAllStringSubmatch(trimmedContent, -1)
	for _, oneMatch := range matches {
		amount := (len(oneMatch) - 1) / 2
		index := amount + 1
		tokenMatches := oneMatch[1:index]
		ruleMatches := oneMatch[index:]

		instructionBuilder := app.instructionBuilder.Create()
		if oneMatch[0] == app.switchChannelCharacter {
			ins, err := instructionBuilder.IsChannelSwitch().Now()
			if err != nil {
				return nil, err
			}

			out = append(out, ins)
		}

		if tokenMatches[0] != "" {
			tokenElement, err := app.tokenElement(tokenMatches)
			if err != nil {
				return nil, err
			}

			ins, err := instructionBuilder.WithElement(tokenElement).Now()
			if err != nil {
				return nil, err
			}

			out = append(out, ins)
		}

		if ruleMatches[0] != "" {
			ruleElement, err := app.ruleElement(ruleMatches)
			if err != nil {
				return nil, err
			}

			ins, err := instructionBuilder.WithElement(ruleElement).Now()
			if err != nil {
				return nil, err
			}

			out = append(out, ins)
		}

	}

	return out, nil
}

func (app *adapter) tokenElement(matches []string) (Element, error) {
	tokenReference, err := app.tokenReference(matches)
	if err != nil {
		return nil, err
	}

	contentIns, err := app.contentBuilder.Create().WithToken(tokenReference).Now()
	if err != nil {
		return nil, err
	}

	return app.buildElement(contentIns, matches)
}

func (app *adapter) ruleElement(matches []string) (Element, error) {
	rule, err := app.rules.Find(matches[1])
	if err != nil {
		return nil, err
	}

	contentIns, err := app.contentBuilder.Create().WithRule(rule).Now()
	if err != nil {
		return nil, err
	}

	return app.buildElement(contentIns, matches)
}

func (app *adapter) tokenReference(matches []string) (TokenReference, error) {
	builder := app.tokenReferenceBuilder.Create().WithName(matches[1])
	specifiers, err := app.tokenSpecifiers(matches[9])
	if err != nil {
		return nil, err
	}

	if specifiers != nil {
		builder.WithSpecifiers(specifiers)
	}

	return builder.Now()
}

func (app *adapter) tokenSpecifiers(content string) (TokenSpecifiers, error) {
	trimmedContent := strings.TrimSpace(content)
	if trimmedContent == "" {
		return nil, nil
	}

	patternStr := fmt.Sprintf(
		"(%s)?(%s)?%s(%s)",
		app.tokenPattern,
		app.rulePattern,
		app.whiteSpacePattern,
		app.cardinalityRangePattern(),
	)

	list := []TokenSpecifier{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(trimmedContent, -1)
	for _, oneMatch := range matches {
		builder := app.tokenSpecifierBuilder.Create()
		if oneMatch[1] != "" {
			tokenSpecifierContent, err := app.tokenSpecifierContentBuilder.Create().WithToken(oneMatch[1]).Now()
			if err != nil {
				return nil, err
			}

			builder.WithContent(tokenSpecifierContent)
		}

		if oneMatch[2] != "" {
			rule, err := app.rules.Find(oneMatch[2])
			if err != nil {
				return nil, err
			}

			tokenSpecifierContent, err := app.tokenSpecifierContentBuilder.Create().WithRule(rule).Now()
			if err != nil {
				return nil, err
			}

			builder.WithContent(tokenSpecifierContent)
		}

		trimmedSpecificContent := strings.TrimSpace(oneMatch[4])
		if trimmedSpecificContent != "" {
			specificCardinality, err := app.specificAdapter.ToSpecific(trimmedSpecificContent)
			if err != nil {
				return nil, err
			}

			if specificCardinality != nil {
				builder.WithCardinality(specificCardinality)
			}
		}

		tokenSpecifier, err := builder.Now()
		if err != nil {
			return nil, err
		}

		list = append(list, tokenSpecifier)
	}

	return app.tokenSpecifiersBuilder.Create().WithTokenSpecifiers(list).Now()
}

func (app *adapter) buildElement(contentIns Content, matches []string) (Element, error) {
	elementBuilder := app.elementBuilder.Create().WithCode(matches[0]).WithContent(contentIns)
	if matches[3] != "" {
		cardinality, err := app.cardinalityAdapter.ToCardinality(matches[3])
		if err != nil {
			return nil, err
		}

		elementBuilder.WithCardinality(cardinality)
	}

	return elementBuilder.Now()
}

func (app *adapter) cardinalityRangePattern() string {
	anythingExceptCardinalityRangeEnd := fmt.Sprintf(app.anythingExcept, app.cardinalityRangeEnd)
	return fmt.Sprintf(
		"%s(%s)%s",
		app.cardinalityRangeBegin,
		anythingExceptCardinalityRangeEnd,
		app.cardinalityRangeEnd,
	)
}
