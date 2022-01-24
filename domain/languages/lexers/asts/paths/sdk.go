package paths

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

// FetchTokenNotFoundFn represents a func to search token that was not found in the original content
type FetchTokenNotFoundFn func(name string) Path

// FetchElementReplacementFn represents a func to search an element that needs to be replaced
type FetchElementReplacementFn func(name string) (Element, []Element)

// NewAdapterBuilder creates a new adapter builder
func NewAdapterBuilder() AdapterBuilder {
	builder := NewBuilder()
	elementBuilder := NewElementBuilder()
	dependenciesBuilder := NewDependenciesBuilder()
	lineBuilder := NewLineBuilder()
	instructionBuilder := NewInstructionBuilder()
	containerBuilder := NewContainerBuilder()
	recursiveTokenBuilder := NewRecursiveTokenBuilder()
	tokenBuilder := NewTokenBuilder()
	ruleBuilder := NewRuleBuilder()
	tokenPathBuilder := NewTokenPathBuilder()
	specifierBuilder := NewSpecifierBuilder()
	cardinalityBuilder := cardinality.NewBuilder()
	return createAdapterBuilder(
		builder,
		elementBuilder,
		dependenciesBuilder,
		lineBuilder,
		instructionBuilder,
		containerBuilder,
		recursiveTokenBuilder,
		tokenBuilder,
		ruleBuilder,
		tokenPathBuilder,
		specifierBuilder,
		cardinalityBuilder,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewDependenciesBuilder creates a new dependencies builder instance
func NewDependenciesBuilder() DependenciesBuilder {
	return createDependenciesBuilder()
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// NewInstructionBuilder creates a new instruction builder instance
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewContainerBuilder creates a new container builder
func NewContainerBuilder() ContainerBuilder {
	return createContainerBuilder()
}

// NewRecursiveTokenBuilder creates a new recursive token builder
func NewRecursiveTokenBuilder() RecursiveTokenBuilder {
	return createRecursiveTokenBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewRuleBuilder creates a new rule builder
func NewRuleBuilder() RuleBuilder {
	return createRuleBuilder()
}

// NewTokenPathBuilder creates a new token path builder
func NewTokenPathBuilder() TokenPathBuilder {
	return createTokenPathBuilder()
}

// NewSpecifierBuilder creates a new specifier builder
func NewSpecifierBuilder() SpecifierBuilder {
	return createSpecifierBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithTokens(tokens tokens.Tokens) AdapterBuilder
	WithTokenNotFoundFunc(tokenNotFoundFn FetchTokenNotFoundFn) AdapterBuilder
	WithTokenReplacementFunc(tokenReplacementFn FetchElementReplacementFn) AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents a path adapter
type Adapter interface {
	ToPath(token tokens.Token) (Path, error)
}

// Builder represents a path builder
type Builder interface {
	Create() Builder
	WithElement(element Element) Builder
	WithDependencies(dep Dependencies) Builder
	Now() (Path, error)
}

// Path represents a path
type Path interface {
	Element() Element
	Dependencies() Dependencies
	CombinedElements() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithName(name string) ElementBuilder
	WithMust(must []Line) ElementBuilder
	WithNot(not []Line) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	Must() []Line
	HasNot() bool
	Not() []Line
}

// DependenciesBuilder represents dependencies builder
type DependenciesBuilder interface {
	Create() DependenciesBuilder
	WithDependencies(dependencies []Element) DependenciesBuilder
	Now() (Dependencies, error)
}

// Dependencies represents dependencies
type Dependencies interface {
	All() []Element
	IsEmpty() bool
	Fetch(name string) (Element, error)
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithInstructions(instructions []Instruction) LineBuilder
	Now() (Line, error)
}

// Line represents a line
type Line interface {
	Instructions() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	IsChannelSwitch() InstructionBuilder
	WithContainer(container Container) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsChannelSwitch() bool
	IsContainer() bool
	Container() Container
}

// ContainerBuilder represents a container builder
type ContainerBuilder interface {
	Create() ContainerBuilder
	WithToken(token Token) ContainerBuilder
	WithRule(rule Rule) ContainerBuilder
	WithRecursive(recursive RecursiveToken) ContainerBuilder
	WithLines(lines []Line) ContainerBuilder
	Now() (Container, error)
}

// Container represents a token or element
type Container interface {
	Cardinality() cardinality.Cardinality
	IsToken() bool
	Token() Token
	IsRule() bool
	Rule() Rule
	IsRecursive() bool
	Recursive() RecursiveToken
	IsLines() bool
	Lines() []Line
}

// RecursiveTokenBuilder represents a recursive token builder
type RecursiveTokenBuilder interface {
	Create() RecursiveTokenBuilder
	WithName(name string) RecursiveTokenBuilder
	WithCardinality(cardinality cardinality.Cardinality) RecursiveTokenBuilder
	WithSpecifiers(specifiers []Specifier) RecursiveTokenBuilder
	Now() (RecursiveToken, error)
}

// RecursiveToken represents a recursive token
type RecursiveToken interface {
	Name() string
	Cardinality() cardinality.Cardinality
	HasSpecifiers() bool
	Specifiers() []Specifier
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithPath(path TokenPath) TokenBuilder
	WithCardinality(cardinality cardinality.Cardinality) TokenBuilder
	Now() (Token, error)
}

// Token represents a rule path token
type Token interface {
	Path() TokenPath
	Cardinality() cardinality.Cardinality
}

// TokenPathBuilder represents a token path builder
type TokenPathBuilder interface {
	Create() TokenPathBuilder
	WithElement(element Element) TokenPathBuilder
	WithSpecifiers(specifiers []Specifier) TokenPathBuilder
	Now() (TokenPath, error)
}

// TokenPath represents a token path
type TokenPath interface {
	Element() Element
	HasSpecifiers() bool
	Specifiers() []Specifier
}

// SpecifierBuilder represents a specifier builder
type SpecifierBuilder interface {
	Create() SpecifierBuilder
	WithContainerName(containerName string) SpecifierBuilder
	WithCardinality(cardinality cardinality.Specific) SpecifierBuilder
	Now() (Specifier, error)
}

// Specifier represents a specifier
type Specifier interface {
	ContainerName() string
	Cardinality() cardinality.Specific
}

// RuleBuilder represents a rule builder
type RuleBuilder interface {
	Create() RuleBuilder
	WithBase(base rules.Rule) RuleBuilder
	WithCardinality(cardinality cardinality.Cardinality) RuleBuilder
	Now() (Rule, error)
}

// Rule represents a rule
type Rule interface {
	Base() rules.Rule
	Cardinality() cardinality.Cardinality
}
