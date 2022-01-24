package mistakes

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	indexBuilder := NewIndexBuilder()
	pathBuilder := NewPathBuilder()
	contentBuilder := NewContentBuilder()
	containsValidNotMatchBuilder := NewContainsValidNotMatchBuilder()
	specifierDoNotMatchBuilder := NewSpecifierDoNotMatchBuilder()
	containsNextElementBuilder := NewContainsNextElementBuilder()
	cardinalityIsInvalidBuilder := NewCardinalityIsInvalidBuilder()
	containsPrefixBuilder := NewContainsPrefixBuilder()
	newLineCharacter := "\n"
	return createAdapter(
		builder,
		indexBuilder,
		pathBuilder,
		contentBuilder,
		containsValidNotMatchBuilder,
		specifierDoNotMatchBuilder,
		containsNextElementBuilder,
		cardinalityIsInvalidBuilder,
		containsPrefixBuilder,
		newLineCharacter,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewIndexBuilder creates a new index builder
func NewIndexBuilder() IndexBuilder {
	return createIndexBuilder()
}

// NewPathBuilder creates a new path builder instance
func NewPathBuilder() PathBuilder {
	return createPathBuilder()
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// NewContainsValidNotMatchBuilder creates a new containsValidNotMatchBuilder builder
func NewContainsValidNotMatchBuilder() ContainsValidNotMatchBuilder {
	return createContainsValidNotMatchBuilder()
}

// NewSpecifierDoNotMatchBuilder creates a new specifierDoNotMatchBuilder builder
func NewSpecifierDoNotMatchBuilder() SpecifierDoNotMatchBuilder {
	return createSpecifierDoNotMatchBuilder()
}

// NewContainsNextElementBuilder creates a new containsNextElementBuilder builder
func NewContainsNextElementBuilder() ContainsNextElementBuilder {
	return createContainsNextElementBuilder()
}

// NewCardinalityIsInvalidBuilder creates a new cardinalityIsInvalidBuilder builder
func NewCardinalityIsInvalidBuilder() CardinalityIsInvalidBuilder {
	return createCardinalityIsInvalidBuilder()
}

// NewContainsPrefixBuilder creates a new containsPrefixBuilder builder
func NewContainsPrefixBuilder() ContainsPrefixBuilder {
	return createContainsPrefixBuilder()
}

// Adapter represents a mistake adapter
type Adapter interface {
	ToMistake(result tokens.Token, canContainPrefix bool) (Mistake, error)
}

// Builder represents a mistake builder
type Builder interface {
	Create() Builder
	WithIndex(index Index) Builder
	WithPath(path Path) Builder
	WithContent(content Content) Builder
	Now() (Mistake, error)
}

// Mistake represents a mistake
type Mistake interface {
	Index() Index
	Path() Path
	Content() Content
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	IsTokenHasNoMatch() ContentBuilder
	WithContainsValidNotMatch(containsValidNotMatch ContainsValidNotMatch) ContentBuilder
	WithSpecifierDoNotMatch(specifierDoNotMatch SpecifierDoNotMatch) ContentBuilder
	WithContainsNextElement(containsNextElement ContainsNextElement) ContentBuilder
	WithCardinalityIsInvalid(cardinalityIsInvalid CardinalityIsInvalid) ContentBuilder
	WithContainsPrefix(containsPrefix ContainsPrefix) ContentBuilder
	Now() (Content, error)
}

// Content represents the mistake content
type Content interface {
	IsTokenHasNoMatch() bool
	IsContainsValidNotMatch() bool
	ContainsValidNotMatch() ContainsValidNotMatch
	IsSpecifierDoNotMatch() bool
	SpecifierDoNotMatch() SpecifierDoNotMatch
	IsContainsNextElement() bool
	ContainsNextElement() ContainsNextElement
	IsCardinalityIsInvalid() bool
	CardinalityIsInvalid() CardinalityIsInvalid
	IsContainsPrefix() bool
	ContainsPrefix() ContainsPrefix
}

// ContainsValidNotMatchBuilder represents a containsValidNotMatch builder
type ContainsValidNotMatchBuilder interface {
	Create() ContainsValidNotMatchBuilder
	WithLine(line uint) ContainsValidNotMatchBuilder
	Now() (ContainsValidNotMatch, error)
}

// ContainsValidNotMatch represents a contains valid not match mistake
type ContainsValidNotMatch interface {
	Line() uint
}

// SpecifierDoNotMatchBuilder represents the specifierDoNotMatch builder
type SpecifierDoNotMatchBuilder interface {
	Create() SpecifierDoNotMatchBuilder
	WithContainerName(containerName string) SpecifierDoNotMatchBuilder
	WithCardinality(cardinality cardinality.Specific) SpecifierDoNotMatchBuilder
	WithAmount(amount uint) SpecifierDoNotMatchBuilder
	Now() (SpecifierDoNotMatch, error)
}

// SpecifierDoNotMatch represents a token specifier do not match mistake
type SpecifierDoNotMatch interface {
	ContainerName() string
	Cardinality() cardinality.Specific
	Amount() uint
}

// ContainsNextElementBuilder represents a containsNextElement builder
type ContainsNextElementBuilder interface {
	Create() ContainsNextElementBuilder
	WithNextElement(nextElement tokens.NextElement) ContainsNextElementBuilder
	Now() (ContainsNextElement, error)
}

// ContainsNextElement represents a contains next element mistake
type ContainsNextElement interface {
	NextElement() tokens.NextElement
}

// CardinalityIsInvalidBuilder represents a cardinalityIsInvalid builder
type CardinalityIsInvalidBuilder interface {
	Create() CardinalityIsInvalidBuilder
	WithCardinality(cardinality cardinality.Cardinality) CardinalityIsInvalidBuilder
	WithAmount(amount uint) CardinalityIsInvalidBuilder
	Now() (CardinalityIsInvalid, error)
}

// CardinalityIsInvalid represents a cardinality is invalid mistake
type CardinalityIsInvalid interface {
	Cardinality() cardinality.Cardinality
	Amount() uint
}

// ContainsPrefixBuilder represents a containsPrefix builder
type ContainsPrefixBuilder interface {
	Create() ContainsPrefixBuilder
	WithPrefix(prefix string) ContainsPrefixBuilder
	Now() (ContainsPrefix, error)
}

// ContainsPrefix represents a contains prefix mistake
type ContainsPrefix interface {
	Prefix() string
}

// IndexBuilder represents an index builder
type IndexBuilder interface {
	Create() IndexBuilder
	WithIndex(index uint) IndexBuilder
	WithLine(line uint) IndexBuilder
	WithColumn(column uint) IndexBuilder
	Now() (Index, error)
}

// Index represents the mistake index
type Index interface {
	Index() uint
	Line() uint
	Column() uint
}

// PathBuilder represents a path builder
type PathBuilder interface {
	Create() PathBuilder
	WithParents(parents []string) PathBuilder
	WithRule(rule string) PathBuilder
	WithToken(token string) PathBuilder
	Now() (Path, error)
}

// Path represents a path
type Path interface {
	Container() PathContainer
	HasParents() bool
	Parents() []string
}

// PathContainer represents a path container
type PathContainer interface {
	IsRule() bool
	Rule() string
	IsToken() bool
	Token() string
}
