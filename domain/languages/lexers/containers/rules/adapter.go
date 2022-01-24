package rules

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type adapter struct {
	builder                     Builder
	ruleBuilder                 RuleBuilder
	elementBuilder              ElementBuilder
	contentBuilder              ContentBuilder
	patternBuilder              PatternBuilder
	patternsBuilder             patterns.Builder
	cardinalityAdapter          cardinality.Adapter
	space                       string
	whiteSpacePattern           string
	rulesPossibilitiesDelimiter string
	anythingExcept              string
	begin                       string
	end                         string
	patternSeparatorDelimiter   string
	possibilityDelimiter        string
	rulePattern                 string
	contentDelimiter            string
	openRegexPatternDelimiter   string
	closeRegexPatternDelimiter  string
	escapeReplacementCode       string
}

func createAdapter(
	builder Builder,
	ruleBuilder RuleBuilder,
	elementBuilder ElementBuilder,
	contentBuilder ContentBuilder,
	patternBuilder PatternBuilder,
	patternsBuilder patterns.Builder,
	cardinalityAdapter cardinality.Adapter,
	space string,
	whiteSpacePattern string,
	rulesPossibilitiesDelimiter string,
	anythingExcept string,
	begin string,
	end string,
	patternSeparatorDelimiter string,
	possibilityDelimiter string,
	rulePattern string,
	contentDelimiter string,
	openRegexPatternDelimiter string,
	closeRegexPatternDelimiter string,
	escapeReplacementCode string,
) Adapter {
	out := adapter{
		builder:                     builder,
		ruleBuilder:                 ruleBuilder,
		elementBuilder:              elementBuilder,
		contentBuilder:              contentBuilder,
		patternBuilder:              patternBuilder,
		patternsBuilder:             patternsBuilder,
		cardinalityAdapter:          cardinalityAdapter,
		space:                       space,
		whiteSpacePattern:           whiteSpacePattern,
		rulesPossibilitiesDelimiter: rulesPossibilitiesDelimiter,
		anythingExcept:              anythingExcept,
		begin:                       begin,
		end:                         end,
		patternSeparatorDelimiter:   patternSeparatorDelimiter,
		possibilityDelimiter:        possibilityDelimiter,
		rulePattern:                 rulePattern,
		contentDelimiter:            contentDelimiter,
		openRegexPatternDelimiter:   openRegexPatternDelimiter,
		closeRegexPatternDelimiter:  closeRegexPatternDelimiter,
		escapeReplacementCode:       escapeReplacementCode,
	}

	return &out
}

// Rules converts content to a Rules instance
func (app *adapter) Rules(content string, patterns patterns.Patterns) (Rules, error) {
	rules, err := app.rulesList(content, patterns)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithRules(rules).Now()
}

func (app *adapter) rulesList(content string, patterns patterns.Patterns) ([]Rule, error) {
	// find the rules:
	anythingExceptionEnd := fmt.Sprintf(app.anythingExcept, app.end)
	patternStr := fmt.Sprintf(
		"(%s)%s(%s)%s",
		app.rulePattern,
		app.begin,
		anythingExceptionEnd,
		app.end,
	)

	rules := []Rule{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		name := strings.TrimSpace(oneMatch[1])
		element, err := app.element(name, oneMatch[2], patterns)
		if err != nil {
			return nil, err
		}

		rule, err := app.ruleBuilder.Create().WithName(name).WithCode(oneMatch[0]).WithElement(element).Now()
		if err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

func (app *adapter) element(ruleName string, content string, patterns patterns.Patterns) (Element, error) {
	anythingExceptConstantDelimiter := fmt.Sprintf(app.anythingExcept, app.contentDelimiter)
	patternStr := fmt.Sprintf(
		"(%s)?(%s(%s)%s)?",
		anythingExceptConstantDelimiter,
		app.contentDelimiter,
		anythingExceptConstantDelimiter,
		app.contentDelimiter,
	)

	escaped := fmt.Sprintf("%s%s", "\\", app.contentDelimiter)
	content = strings.Replace(content, escaped, app.escapeReplacementCode, -1)

	pattern := regexp.MustCompile(patternStr)
	match := pattern.FindStringSubmatch(content)

	patternContent := strings.TrimSpace(match[1])
	if patternContent != "" {
		pattern, err := app.pattern(ruleName, patternContent, patterns)
		if err != nil {
			return nil, err
		}

		content, err := app.contentBuilder.Create().WithPattern(pattern).Now()
		if err != nil {
			return nil, err
		}

		element, err := app.elementBuilder.Create().WithCode(match[0]).WithContent(content).Now()
		if err != nil {
			return nil, err
		}

		return element, nil
	}

	constantContent := strings.Replace(match[3], app.escapeReplacementCode, app.contentDelimiter, -1)
	contentIns, err := app.contentBuilder.Create().WithConstant(constantContent).Now()
	if err != nil {
		return nil, err
	}

	code := strings.Replace(match[0], app.escapeReplacementCode, app.contentDelimiter, -1)
	return app.elementBuilder.Create().WithCode(code).WithContent(contentIns).Now()
}

func (app *adapter) pattern(ruleName string, content string, ptrns patterns.Patterns) (Pattern, error) {
	anythingExceptCloseRegexDelimiter := fmt.Sprintf(app.anythingExcept, app.closeRegexPatternDelimiter)
	patternStr := fmt.Sprintf(
		"%s%s(%s)%s%s",
		app.openRegexPatternDelimiter,
		app.whiteSpacePattern,
		anythingExceptCloseRegexDelimiter,
		app.whiteSpacePattern,
		app.closeRegexPatternDelimiter,
	)

	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindStringSubmatch(content)
	cardinalityContent := strings.Replace(content, matches[0], "", 1)
	cardinality, err := app.cardinalityAdapter.ToCardinality(cardinalityContent)
	if err != nil {
		return nil, err
	}

	list := []patterns.Pattern{}
	patternNames := strings.Split(matches[1], app.space)
	for _, onePattern := range patternNames {
		if ptrns == nil {
			str := fmt.Sprintf("the rule (name: %s) contains a pattern (name: %s) that is not defined", ruleName, onePattern)
			return nil, errors.New(str)
		}

		pattern, err := ptrns.Find(onePattern)
		if err != nil {
			return nil, err
		}

		list = append(list, pattern)
	}

	subPatterns, err := app.patternsBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, err
	}

	return app.patternBuilder.Create().WithCardinality(cardinality).WithSubPatterns(subPatterns).Now()
}
