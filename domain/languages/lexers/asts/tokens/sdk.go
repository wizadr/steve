package tokens

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/results"
	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

// NewAdapterBuilder creates a new adapter builder instance
func NewAdapterBuilder() AdapterBuilder {
	builder := NewBuilder()
	tokenContentBuilder := NewTokenContentBuilder()
	lineMatchBuilder := NewLineMatchBuilder()
	blockMatchBuilder := NewBlockMatchBuilder()
	nextElementBuilder := NewNextElementBuilder()
	matchesBuilder := NewMatchesBuilder()
	matchBuilder := NewMatchBuilder()
	skipMatchBuilder := NewSkipMatchBuilder()
	tokenMatchBuilder := NewTokenMatchBuilder()
	specifierBuilder := NewSpecifierBuilder()
	tokenResultBuilder := NewTokenResultBuilder()
	tokenResultMatchesBuilder := NewTokenResultMatchesBuilder()
	ruleMatchBuilder := NewRuleMatchBuilder()
	ruleMatchResultBuilder := NewRuleMatchResultBuilder()
	resultBuilder := NewResultBuilder()
	resultAdapter := results.NewAdapter()
	pathTokenPathBuilder := paths.NewTokenPathBuilder()
	pathTokenBuilder := paths.NewTokenBuilder()
	pathBuilder := paths.NewBuilder()
	pathAdapterBuilder := paths.NewAdapterBuilder()
	return createAdapterBuilder(
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
		pathAdapterBuilder,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenContentBuilder creates a new token content builder
func NewTokenContentBuilder() TokenContentBuilder {
	return createTokenContentBuilder()
}

// NewLineMatchBuilder creates a new line match builder
func NewLineMatchBuilder() LineMatchBuilder {
	return createLineMatchBuilder()
}

// NewBlockMatchBuilder creates a new block match builder
func NewBlockMatchBuilder() BlockMatchBuilder {
	return createBlockMatchBuilder()
}

// NewNextElementBuilder creates a new next element builder
func NewNextElementBuilder() NextElementBuilder {
	return createNextElementBuilder()
}

// NewMatchesBuilder creates a new matches builder
func NewMatchesBuilder() MatchesBuilder {
	return createMatchesBuilder()
}

// NewMatchBuilder creates a new match builder
func NewMatchBuilder() MatchBuilder {
	return createMatchBuilder()
}

// NewSkipMatchBuilder initializes the builder
func NewSkipMatchBuilder() SkipMatchBuilder {
	return createSkipMatchBuilder()
}

// NewTokenMatchBuilder creates a new token match builder
func NewTokenMatchBuilder() TokenMatchBuilder {
	return createTokenMatchBuilder()
}

// NewSpecifierBuilder creates a new specifier builder
func NewSpecifierBuilder() SpecifierBuilder {
	return createSpecifierBuilder()
}

// NewTokenResultBuilder creates a new token result builder
func NewTokenResultBuilder() TokenResultBuilder {
	return createTokenResultBuilder()
}

// NewTokenResultMatchesBuilder creates a new token result matches builder
func NewTokenResultMatchesBuilder() TokenResultMatchesBuilder {
	return createTokenResultMatchesBuilder()
}

// NewRuleMatchBuilder creates a new rule match builder
func NewRuleMatchBuilder() RuleMatchBuilder {
	return createRuleMatchBuilder()
}

// NewRuleMatchResultBuilder creates a new rule match result builder
func NewRuleMatchResultBuilder() RuleMatchResultBuilder {
	return createRuleMatchResultBuilder()
}

// NewResultBuilder creates a new result builder
func NewResultBuilder() ResultBuilder {
	return createResultBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithToken(parsed tokens.Token) AdapterBuilder
	WithTokens(tokens tokens.Tokens) AdapterBuilder
	WithChannels(channels tokens.Tokens) AdapterBuilder
	WithTokenNotFoundFunc(tokenNotFoundFn paths.FetchTokenNotFoundFn) AdapterBuilder
	WithTokenReplacementFunc(tokenReplacementFn paths.FetchElementReplacementFn) AdapterBuilder
	CanContainPrefix() AdapterBuilder
	CanContainSuffix() AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents a token adapter
type Adapter interface {
	ToToken(script string) (Token, error)
}

// Builder represents a token builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPath(path paths.Path) Builder
	WithContent(content TokenContent) Builder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	IsValid() bool
	IsExact() bool
	Discoveries() string
	Path() paths.Path
	HasContent() bool
	Content() TokenContent
}

// TokenContentBuilder represents a token content builder
type TokenContentBuilder interface {
	Create() TokenContentBuilder
	WithMust(matches LineMatch) TokenContentBuilder
	WithNot(not LineMatch) TokenContentBuilder
	Now() (TokenContent, error)
}

// TokenContent represents token content
type TokenContent interface {
	IsValid() bool
	IsExact() bool
	Discoveries() string
	Must() LineMatch
	HasNot() bool
	Not() LineMatch
}

// LineMatchBuilder represents a line match builder
type LineMatchBuilder interface {
	Create() LineMatchBuilder
	WithIndex(index uint) LineMatchBuilder
	WithMatches(matches Matches) LineMatchBuilder
	Now() (LineMatch, error)
}

// LineMatch represents a line match
type LineMatch interface {
	Index() uint
	Matches() Matches
}

// MatchesBuilder represents a matches builder
type MatchesBuilder interface {
	Create() MatchesBuilder
	WithList(matches []BlockMatch) MatchesBuilder
	Now() (Matches, error)
}

// Matches represents matches
type Matches interface {
	BlockMatches() []BlockMatch
	Input() string
	IsValid() bool
	IsExact() bool
	Discoveries() string
}

// BlockMatchBuilder represents a block match builder
type BlockMatchBuilder interface {
	Create() BlockMatchBuilder
	WithContainer(container Match) BlockMatchBuilder
	WithBlock(block BlockMatch) BlockMatchBuilder
	WithNextElement(nextElement NextElement) BlockMatchBuilder
	WithChannelPrefix(channelPrefix LineMatch) BlockMatchBuilder
	Now() (BlockMatch, error)
}

// BlockMatch represents a block match
type BlockMatch interface {
	Input() string
	IsValid() bool
	IsExact() bool
	Discoveries() string
	Content() BlockMatchContent
	HasChannelPrefix() bool
	ChannelPrefix() LineMatch
}

// BlockMatchContent represents a block match content
type BlockMatchContent interface {
	IsContainer() bool
	Container() Match
	IsBlock() bool
	Block() BlockMatch
	IsNextElement() bool
	NextElement() NextElement
}

// NextElementBuilder represents the next element builder
type NextElementBuilder interface {
	Create() NextElementBuilder
	WithConstant(constant string) NextElementBuilder
	WithPattern(pattern patterns.Pattern) NextElementBuilder
	Now() (NextElement, error)
}

// NextElement represents the next element
type NextElement interface {
	IsConstant() bool
	Constant() string
	IsPattern() bool
	Pattern() patterns.Pattern
}

// MatchBuilder represents a match builder
type MatchBuilder interface {
	Create() MatchBuilder
	WithToken(token TokenMatch) MatchBuilder
	WithRule(rule RuleMatch) MatchBuilder
	WithSkip(skip SkipMatch) MatchBuilder
	WithLine(lineMatch LineMatch) MatchBuilder
	Now() (Match, error)
}

// Match represents a rule match
type Match interface {
	Input() string
	IsValid() bool
	IsExact() bool
	Discoveries() string
	Content() MatchContent
}

// MatchContent represents a match content
type MatchContent interface {
	IsToken() bool
	Token() TokenMatch
	IsRule() bool
	Rule() RuleMatch
	IsSkip() bool
	Skip() SkipMatch
	IsLine() bool
	Line() LineMatch
}

// SkipMatchBuilder represents a skip match builder
type SkipMatchBuilder interface {
	Create() SkipMatchBuilder
	WithToken(token TokenMatch) SkipMatchBuilder
	WithRule(rule RuleMatch) SkipMatchBuilder
	Now() (SkipMatch, error)
}

// SkipMatch represents a skip match
type SkipMatch interface {
	Name() string
	Input() string
	Content() SkipMatchContent
}

// SkipMatchContent representsa skip matchcontent:
type SkipMatchContent interface {
	IsToken() bool
	Token() TokenMatch
	IsRule() bool
	Rule() RuleMatch
}

// TokenMatchBuilder represents a token match builder
type TokenMatchBuilder interface {
	Create() TokenMatchBuilder
	WithSpecifiers(specifiers []Specifier) TokenMatchBuilder
	WithPath(path paths.TokenPath) TokenMatchBuilder
	WithResult(result TokenResult) TokenMatchBuilder
	Now() (TokenMatch, error)
}

// TokenMatch represets a token match
type TokenMatch interface {
	IsValid() bool
	IsExact() bool
	Path() paths.TokenPath
	Result() TokenResult
	HasSpecifiers() bool
	Specifiers() []Specifier
}

// SpecifierBuilder represents a specifier builder
type SpecifierBuilder interface {
	Create() SpecifierBuilder
	WithContainerName(containerName string) SpecifierBuilder
	WithCardinality(cardinality cardinality.Specific) SpecifierBuilder
	WithResult(result TokenResult) SpecifierBuilder
	Now() (Specifier, error)
}

// Specifier represents a specifier
type Specifier interface {
	ContainerName() string
	Cardinality() cardinality.Specific
	Result() TokenResult
	Amount() uint
	IsValid() bool
}

// TokenResultBuilder represents a token result builder
type TokenResultBuilder interface {
	Create() TokenResultBuilder
	WithInput(input string) TokenResultBuilder
	WithCardinality(cardinality cardinality.Cardinality) TokenResultBuilder
	WithMatches(matches TokenResultMatches) TokenResultBuilder
	Now() (TokenResult, error)
}

// TokenResult represents token result
type TokenResult interface {
	Input() string
	Cardinality() cardinality.Cardinality
	IsCardinalityValid() bool
	Amount() uint
	Discoveries() string
	IsValid() bool
	IsExact() bool
	HasMatches() bool
	Matches() TokenResultMatches
}

// TokenResultMatchesBuilder represents token result matches builder
type TokenResultMatchesBuilder interface {
	Create() TokenResultMatchesBuilder
	WithResults(results []Token) TokenResultMatchesBuilder
	Now() (TokenResultMatches, error)
}

// TokenResultMatches represents token result matches
type TokenResultMatches interface {
	All() []Token
	Amount() uint
	IsValid() bool
	IsExact() bool
	Discoveries() string
	IsAmount(amount uint) bool
	IsInRange(min uint, max uint) bool
	IsAtLeast(min uint) bool
}

// RuleMatchBuilder represents a rule match builder
type RuleMatchBuilder interface {
	Create() RuleMatchBuilder
	WithRule(rule paths.Rule) RuleMatchBuilder
	WithResult(result RuleMatchResult) RuleMatchBuilder
	CanContainPrefix() RuleMatchBuilder
	Now() (RuleMatch, error)
}

// RuleMatch represents a rule match
type RuleMatch interface {
	IsValid() bool
	IsExact() bool
	CanHavePrefix() bool
	Rule() paths.Rule
	Result() RuleMatchResult
}

// RuleMatchResultBuilder represents the rule match result builder
type RuleMatchResultBuilder interface {
	Create() RuleMatchResultBuilder
	WithPath(path paths.Rule) RuleMatchResultBuilder
	WithResult(result Result) RuleMatchResultBuilder
	Now() (RuleMatchResult, error)
}

// RuleMatchResult represents the rule match result
type RuleMatchResult interface {
	IsCardinalityValid() bool
	Path() paths.Rule
	Result() Result
}

// ResultBuilder represents a result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithResults(results []results.Result) ResultBuilder
	Now() (Result, error)
}

// Result represents result
type Result interface {
	Amount() uint
	Discoveries() string
	HasResults() bool
	Results() []results.Result
}
