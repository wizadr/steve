package singles

import "github.com/steve-care-software/steve/domain/languages/lexers/resources/multiples"

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a single resource builder
type Builder interface {
	Create() Builder
	WithRootToken(rootToken string) Builder
	WithMultiple(multiple multiples.Multiple) Builder
	Now() (Single, error)
}

// Single represents a single resource
type Single interface {
	RootToken() string
	Multiple() multiples.Multiple
}
