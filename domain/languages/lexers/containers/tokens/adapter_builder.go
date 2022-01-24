package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

type adapterBuilder struct {
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

func createAdapterBuilder(
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
) AdapterBuilder {
	out := adapterBuilder{
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
		rules:                          nil,
	}

	return &out
}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(
		app.builder,
		app.tokenBuilder,
		app.testSuiteBuilder,
		app.linesBuilder,
		app.lineBuilder,
		app.blockBuilder,
		app.instructionBuilder,
		app.elementBuilder,
		app.cardinalityAdapter,
		app.specificAdapter,
		app.contentBuilder,
		app.tokenReferenceBuilder,
		app.tokenSpecifiersBuilder,
		app.tokenSpecifierBuilder,
		app.tokenSpecifierContentBuilder,
		app.tokenPattern,
		app.rulePattern,
		app.anythingExcept,
		app.begin,
		app.or,
		app.end,
		app.notDelimiter,
		app.testDelimiter,
		app.testLineBegin,
		app.testLineDelimiter,
		app.testLineEnd,
		app.whiteSpacePattern,
		app.tokenSpecifierPrefix,
		app.tokenSpecifierSuffix,
		app.cardinalityZeroMultiplePattern,
		app.cardinalityMultiplePattern,
		app.cardinalityOptionalPattern,
		app.cardinalityRangeBegin,
		app.cardinalityRangeEnd,
		app.cardinalityRangeSeparator,
		app.switchChannelCharacter,
	)
}

// WithRules add rules to the builder
func (app *adapterBuilder) WithRules(rules rules.Rules) AdapterBuilder {
	app.rules = rules
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.rules == nil {
		return nil, errors.New("the rules are mandatory in order to build an Adapter instance")
	}

	return createAdapter(
		app.builder,
		app.tokenBuilder,
		app.testSuiteBuilder,
		app.linesBuilder,
		app.lineBuilder,
		app.blockBuilder,
		app.instructionBuilder,
		app.elementBuilder,
		app.cardinalityAdapter,
		app.specificAdapter,
		app.contentBuilder,
		app.tokenReferenceBuilder,
		app.tokenSpecifiersBuilder,
		app.tokenSpecifierBuilder,
		app.tokenSpecifierContentBuilder,
		app.tokenPattern,
		app.rulePattern,
		app.anythingExcept,
		app.begin,
		app.or,
		app.end,
		app.notDelimiter,
		app.testDelimiter,
		app.testLineBegin,
		app.testLineDelimiter,
		app.testLineEnd,
		app.whiteSpacePattern,
		app.tokenSpecifierPrefix,
		app.tokenSpecifierSuffix,
		app.cardinalityZeroMultiplePattern,
		app.cardinalityMultiplePattern,
		app.cardinalityOptionalPattern,
		app.cardinalityRangeBegin,
		app.cardinalityRangeEnd,
		app.cardinalityRangeSeparator,
		app.switchChannelCharacter,
		app.rules,
	), nil
}
