package rules

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

// RulePattern represents the rule pattern
const RulePattern = "[A-Z\\_]+"

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	ruleBuilder := NewRuleBuilder()
	elementBuilder := NewElementBuilder()
	contentBuilder := NewContentBuilder()
	patternBuilder := NewPatternBuilder()
	patternsBuilder := patterns.NewBuilder()
	cardinalityAdapter := cardinality.NewAdapter()
	space := " "
	whiteSpacePattern := "[ \t\r\n]*"
	rulesPossibilitiesDelimiter := "---"
	anythingExcept := "[^%s]*"
	begin := ":"
	end := ";"
	patternSeparatorDelimiter := "->"
	possibilityDelimiter := ","
	contentDelimiter := "\""
	openRegexPatternDelimiter := "\\["
	closeRegexPatternDelimiter := "\\]"
	escapeReplacementCode := "###"
	return createAdapter(
		builder,
		ruleBuilder,
		elementBuilder,
		contentBuilder,
		patternBuilder,
		patternsBuilder,
		cardinalityAdapter,
		space,
		whiteSpacePattern,
		rulesPossibilitiesDelimiter,
		anythingExcept,
		begin,
		end,
		patternSeparatorDelimiter,
		possibilityDelimiter,
		RulePattern,
		contentDelimiter,
		openRegexPatternDelimiter,
		closeRegexPatternDelimiter,
		escapeReplacementCode,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewRuleBuilder creates a new rule builder instance
func NewRuleBuilder() RuleBuilder {
	return createRuleBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// NewPatternBuilder creates a new master pattern builder
func NewPatternBuilder() PatternBuilder {
	return createPatternBuilder()
}

// Adapter represents a rule adapter
type Adapter interface {
	Rules(content string, patterns patterns.Patterns) (Rules, error)
}

// Builder represents a rules builder
type Builder interface {
	Create() Builder
	WithRules(rules []Rule) Builder
	Now() (Rules, error)
}

// Rules represents rules
type Rules interface {
	All() []Rule
	Find(name string) (Rule, error)
}

// RuleBuilder represents a rule builder
type RuleBuilder interface {
	Create() RuleBuilder
	WithName(name string) RuleBuilder
	WithCode(code string) RuleBuilder
	WithElement(element Element) RuleBuilder
	Now() (Rule, error)
}

// Rule represents a rule
type Rule interface {
	Name() string
	Code() string
	Element() Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithContent(content Content) ElementBuilder
	WithCode(code string) ElementBuilder
	Now() (Element, error)
}

// Element represents a rule element
type Element interface {
	Content() Content
	Code() string
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithConstant(constant string) ContentBuilder
	WithPattern(pattern Pattern) ContentBuilder
	Now() (Content, error)
}

// Content represents an element content
type Content interface {
	IsConstant() bool
	Constant() string
	IsPattern() bool
	Pattern() Pattern
}

// PatternBuilder represents a master pattern builder
type PatternBuilder interface {
	Create() PatternBuilder
	WithSubPatterns(subPatterns patterns.Patterns) PatternBuilder
	WithCardinality(cardinality cardinality.Cardinality) PatternBuilder
	Now() (Pattern, error)
}

// Pattern represents a master pattern
type Pattern interface {
	SubPatterns() patterns.Patterns
	Cardinality() cardinality.Cardinality
}
