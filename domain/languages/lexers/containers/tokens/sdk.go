package tokens

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

// ToPatternStringsFn converts a token to a must and must not pattern strings
type ToPatternStringsFn func(tok Token) ([]string, []string, error)

// NewAdapterBuilder creates a new adapter builder instance
func NewAdapterBuilder() AdapterBuilder {
	builder := NewBuilder()
	tokenBuilder := NewTokenBuilder()
	testSuiteBuilder := NewTestSuiteBuilder()
	linesBuilder := NewLinesBuilder()
	lineBuilder := NewLineBuilder()
	blockBuilder := NewBlockBuilder()
	instructionBuilder := NewInstructionBuilder()
	elementBuilder := NewElementBuilder()
	cardinalityAdapter := cardinality.NewAdapter()
	specificAdapter := cardinality.NewSpecificAdapter()
	contentBuilder := NewContentBuilder()
	tokenReferenceBuilder := NewTokenReferenceBuilder()
	tokenSpecifiersBuilder := NewTokenSpecifiersBuilder()
	tokenSpecifierBuilder := NewTokenSpecifierBuilder()
	tokenSpecifierContentBuilder := NewTokenSpecifierContentBuilder()

	tokenPattern := "[a-z]{1}[a-zA-Z]*"
	anythingExcept := "[^%s]+"
	begin := ":"
	or := "|"
	end := ";"
	notDelimiter := "---"
	testDelimiter := "+++"
	testLineBegin := ":"
	testLineDelimiter := "###"
	testLineEnd := ";"
	whiteSpacePattern := "[ \t\r\n]*"
	tokenSpecifierPrefix := "->"
	tokenSpecifierSuffix := "<-"
	cardinalityZeroMultiplePattern := "[\\*]{1}"
	cardinalityMultiplePattern := "[\\+]{1}"
	cardinalityOptionalPattern := "[\\?]{1}"
	cardinalityRangeBegin := "{"
	cardinalityRangeEnd := "}"
	cardinalityRangeSeparator := ","
	switchChannelCharacter := "#"
	return createAdapterBuilder(
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
		rules.RulePattern,
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
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a new token builder instance
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewTestSuiteBuilder creates a new testSuite builder
func NewTestSuiteBuilder() TestSuiteBuilder {
	return createTestSuiteBuilder()
}

// NewLinesBuilder creates a new lines builder
func NewLinesBuilder() LinesBuilder {
	return createLinesBuilder()
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	return createBlockBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// NewTokenReferenceBuilder creates a new token reference builder
func NewTokenReferenceBuilder() TokenReferenceBuilder {
	return createTokenReferenceBuilder()
}

// NewTokenSpecifiersBuilder creates a new token specifiers builder
func NewTokenSpecifiersBuilder() TokenSpecifiersBuilder {
	return createTokenSpecifiersBuilder()
}

// NewTokenSpecifierBuilder creates a new token specifier builder
func NewTokenSpecifierBuilder() TokenSpecifierBuilder {
	return createTokenSpecifierBuilder()
}

// NewTokenSpecifierContentBuilder creates a new token specific content builder
func NewTokenSpecifierContentBuilder() TokenSpecifierContentBuilder {
	return createTokenSpecifierContentBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithRules(rules rules.Rules) AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents a token adapter
type Adapter interface {
	ToTokens(content string) (Tokens, error)
}

// Builder represents a tokens builder
type Builder interface {
	Create() Builder
	WithTokens(tokens []Token) Builder
	Now() (Tokens, error)
}

// Tokens represets tokens
type Tokens interface {
	All() []Token
	Find(name string) (Token, error)
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithBlock(block Block) TokenBuilder
	WithTestSuite(testSuite TestSuite) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Block() Block
	HasTestSuite() bool
	TestSuite() TestSuite
}

// TestSuiteBuilder represents a test suite builder
type TestSuiteBuilder interface {
	Create() TestSuiteBuilder
	WithLines(lines []string) TestSuiteBuilder
	Now() (TestSuite, error)
}

// TestSuite represents a test suite
type TestSuite interface {
	Lines() []string
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMust(must Lines) BlockBuilder
	WithNot(not Lines) BlockBuilder
	Now() (Block, error)
}

// Block represents a token block
type Block interface {
	Must() Lines
	HasNot() bool
	Not() Lines
}

// LinesBuilder represents a lines builder
type LinesBuilder interface {
	Create() LinesBuilder
	WithLines(lines []Line) LinesBuilder
	Now() (Lines, error)
}

// Lines represents lines
type Lines interface {
	All() []Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithInstructions(instructions []Instruction) LineBuilder
	Now() (Line, error)
}

// Line represents a token line
type Line interface {
	Instructions() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	IsChannelSwitch() InstructionBuilder
	WithElement(element Element) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsChannelSwitch() bool
	IsElement() bool
	Element() Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithContent(content Content) ElementBuilder
	WithCode(code string) ElementBuilder
	WithCardinality(cardinality cardinality.Cardinality) ElementBuilder
	Now() (Element, error)
}

// Element represents a token element
type Element interface {
	Content() Content
	Code() string
	HasCardinality() bool
	Cardinality() cardinality.Cardinality
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithToken(token TokenReference) ContentBuilder
	WithRule(rule rules.Rule) ContentBuilder
	Now() (Content, error)
}

// Content represents an element content
type Content interface {
	Name() string
	IsToken() bool
	Token() TokenReference
	IsRule() bool
	Rule() rules.Rule
}

// TokenReferenceBuilder represents a token reference builder
type TokenReferenceBuilder interface {
	Create() TokenReferenceBuilder
	WithName(name string) TokenReferenceBuilder
	WithSpecifiers(specifiers TokenSpecifiers) TokenReferenceBuilder
	Now() (TokenReference, error)
}

// TokenReference represents a token reference
type TokenReference interface {
	Name() string
	HasSpecifiers() bool
	Specifiers() TokenSpecifiers
}

// TokenSpecifiersBuilder represents token specifiers builder
type TokenSpecifiersBuilder interface {
	Create() TokenSpecifiersBuilder
	WithTokenSpecifiers(specifiers []TokenSpecifier) TokenSpecifiersBuilder
	Now() (TokenSpecifiers, error)
}

// TokenSpecifiers represents token specifiers
type TokenSpecifiers interface {
	All() []TokenSpecifier
	Find(name string) (TokenSpecifier, error)
}

// TokenSpecifierBuilder represents a token specifier builder
type TokenSpecifierBuilder interface {
	Create() TokenSpecifierBuilder
	WithContent(content TokenSpecifierContent) TokenSpecifierBuilder
	WithCardinality(cardinality cardinality.Specific) TokenSpecifierBuilder
	Now() (TokenSpecifier, error)
}

// TokenSpecifier represents a token specifier
type TokenSpecifier interface {
	Content() TokenSpecifierContent
	Cardinality() cardinality.Specific
}

// TokenSpecifierContentBuilder represents a token specifier content builder
type TokenSpecifierContentBuilder interface {
	Create() TokenSpecifierContentBuilder
	WithRule(rule rules.Rule) TokenSpecifierContentBuilder
	WithToken(token string) TokenSpecifierContentBuilder
	Now() (TokenSpecifierContent, error)
}

// TokenSpecifierContent represents a token specifier content
type TokenSpecifierContent interface {
	Name() string
	IsToken() bool
	Token() string
	IsRule() bool
	Rule() rules.Rule
}
