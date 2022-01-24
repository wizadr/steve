package asts

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	ast_tokens "github.com/steve-care-software/steve/domain/languages/lexers/asts/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

//NewAdapterBuilder creates a new adapter builder
func NewAdapterBuilder() AdapterBuilder {
	discoveriesBuilder := patterns.NewDiscoveriesBuilder()
	rulesAdapter := rules.NewAdapter()
	tokensAdapterBuilder := tokens.NewAdapterBuilder()
	astTokensAdapterBuilder := ast_tokens.NewAdapterBuilder()
	mistakeAdapter := mistakes.NewAdapter()
	builder := NewBuilder()
	tokenBuilder := NewTokenBuilder()
	lineMatchBuilder := NewLineMatchBuilder()
	matchBuilder := NewMatchBuilder()
	tokenMatchBuilder := NewTokenMatchBuilder()
	ruleMatchBuilder := NewRuleMatchBuilder()
	resultBuilder := NewResultBuilder()
	return createAdapterBuilder(
		discoveriesBuilder,
		rulesAdapter,
		tokensAdapterBuilder,
		astTokensAdapterBuilder,
		mistakeAdapter,
		builder,
		tokenBuilder,
		lineMatchBuilder,
		matchBuilder,
		tokenMatchBuilder,
		ruleMatchBuilder,
		resultBuilder,
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

// NewLineMatchBuilder creates a new lineMatch builder
func NewLineMatchBuilder() LineMatchBuilder {
	return createLineMatchBuilder()
}

// NewMatchBuilder creates a new match builder
func NewMatchBuilder() MatchBuilder {
	return createMatchBuilder()
}

// NewTokenMatchBuilder creates a new tokenMatch builder
func NewTokenMatchBuilder() TokenMatchBuilder {
	return createTokenMatchBuilder()
}

// NewRuleMatchBuilder creates a new ruleMatch instance
func NewRuleMatchBuilder() RuleMatchBuilder {
	return createRuleMatchBuilder()
}

//NewResultBuilder creates a new result builder
func NewResultBuilder() ResultBuilder {
	return createResultBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithRoot(root tokens.Token) AdapterBuilder
	WithTokens(tokens tokens.Tokens) AdapterBuilder
	WithChannels(channels tokens.Tokens) AdapterBuilder
	WithTokenNotFoundFunc(tokenNotFoundFn paths.FetchTokenNotFoundFn) AdapterBuilder
	WithTokenReplacementFunc(tokenReplacementFn paths.FetchElementReplacementFn) AdapterBuilder
	CanContainPrefix() AdapterBuilder
	CanContainSuffix() AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents an ast
type Adapter interface {
	ToAST(script string) (AST, error)
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithMistake(mistake mistakes.Mistake) Builder
	WithSuccess(ast Token) Builder
	Now() (AST, error)
}

// AST represents an AST instance
type AST interface {
	IsMistake() bool
	Mistake() mistakes.Mistake
	IsSuccess() bool
	Success() Token
}

// TokenBuilder represents a Token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithMatch(match LineMatch) TokenBuilder
	Now() (Token, error)
}

// Token represents a Token
type Token interface {
	Name() string
	Match() LineMatch
	Prefix() string
	Discovery() string
	Suffix() string
}

// LineMatchBuilder represents a lineMatch builder
type LineMatchBuilder interface {
	Create() LineMatchBuilder
	WithIndex(index uint) LineMatchBuilder
	WithMatches(matches []Match) LineMatchBuilder
	Now() (LineMatch, error)
}

// LineMatch represents a line match
type LineMatch interface {
	Index() uint
	HasMatches() bool
	Matches() []Match
}

// MatchBuilder represents a match builder
type MatchBuilder interface {
	Create() MatchBuilder
	WithToken(token TokenMatch) MatchBuilder
	WithRule(rule RuleMatch) MatchBuilder
	Now() (Match, error)
}

// Match represents a rule match
type Match interface {
	IsToken() bool
	Token() TokenMatch
	IsRule() bool
	Rule() RuleMatch
}

// TokenMatchBuilder represents a tokenMatch builder
type TokenMatchBuilder interface {
	Create() TokenMatchBuilder
	WithToken(token string) TokenMatchBuilder
	WithMatches(matches []Token) TokenMatchBuilder
	Now() (TokenMatch, error)
}

// TokenMatch represets a token match
type TokenMatch interface {
	Token() string
	Discoveries() []string
	Matches() []Token
}

// RuleMatchBuilder represents a ruleMatch builder
type RuleMatchBuilder interface {
	Create() RuleMatchBuilder
	WithRule(rule string) RuleMatchBuilder
	WithResult(result Result) RuleMatchBuilder
	Now() (RuleMatch, error)
}

// RuleMatch represents a rule match
type RuleMatch interface {
	Rule() string
	Result() Result
}

// ResultBuilder represents the result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithInput(input string) ResultBuilder
	WithDiscoveries(discoveries patterns.Discoveries) ResultBuilder
	WithPrefix(prefix string) ResultBuilder
	WithSuffix(suffix string) ResultBuilder
	Now() (Result, error)
}

// Result represents results
type Result interface {
	Input() string
	Discoveries() patterns.Discoveries
	HasPrefix() bool
	Prefix() string
	HasSuffix() bool
	Suffix() string
}
