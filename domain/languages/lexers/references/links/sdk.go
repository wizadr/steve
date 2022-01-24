package links

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/paths"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewLinkBuilder creates a new link builder
func NewLinkBuilder() LinkBuilder {
	return createLinkBuilder()
}

// NewIncludeBuilder represents an include builder
func NewIncludeBuilder() IncludeBuilder {
	return createIncludeBuilder()
}

// Adapter represents the links adapter
type Adapter interface {
	ToLinks(content string) (Links, error)
}

// Builder represents a references builder
type Builder interface {
	Create() Builder
	WithLinks(links []Link) Builder
	Now() (Links, error)
}

// Links represents links
type Links interface {
	All() []Link
	FetchByLocalToken(localToken string) (Link, error)
	FetchByReferenceToken(includeName string, refToken string) (Link, error)
}

// LinkBuilder represents a link builder
type LinkBuilder interface {
	Create() LinkBuilder
	WithLocalToken(local string) LinkBuilder
	WithReferenceToken(reference string) LinkBuilder
	WithInclude(include Include) LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	LocalToken() string
	ReferenceToken() string
	Include() Include
}

// IncludeBuilder represents an include builder
type IncludeBuilder interface {
	Create() IncludeBuilder
	WithName(name string) IncludeBuilder
	WithPaths(paths paths.Paths) IncludeBuilder
	Now() (Include, error)
}

// Include represents an include
type Include interface {
	Name() string
	Paths() paths.Paths
}
