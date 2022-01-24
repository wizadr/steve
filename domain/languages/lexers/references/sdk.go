package references

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/references/links"
)

// NewAdapter createsa new adapter instance
func NewAdapter() Adapter {
	includeBuilder := links.NewIncludeBuilder()
	pathsBuilder := paths.NewBuilder()
	linkBuilder := links.NewLinkBuilder()
	linksBuilder := links.NewBuilder()
	builder := NewBuilder()
	includeKeyword := "include"
	referenceKeyword := "reference"
	replacementKeyword := "replace"
	quoteChar := "\""
	dotChar := "."
	tokenPattern := "[a-z]{1}[a-zA-Z]*"
	anythingExcept := "[^%s]*"
	spacePattern := "[ \t\r\n]*"
	pathSeparator := ";"
	return createAdapter(
		includeBuilder,
		pathsBuilder,
		linkBuilder,
		linksBuilder,
		builder,
		includeKeyword,
		quoteChar,
		dotChar,
		referenceKeyword,
		replacementKeyword,
		tokenPattern,
		anythingExcept,
		spacePattern,
		pathSeparator,
	)
}

// NewBuilder creates a new builderinstance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the references adapter
type Adapter interface {
	ToReferences(content string) (References, error)
}

// Builder represents a references builder
type Builder interface {
	Create() Builder
	WithReferences(references links.Links) Builder
	WithReplacements(replacements links.Links) Builder
	Now() (References, error)
}

// References represents references
type References interface {
	HasReferences() bool
	References() links.Links
	HasReplacements() bool
	Replacements() links.Links
}
