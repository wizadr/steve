package multiples

import "github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a multiple resource builder
type Builder interface {
	Create() Builder
	WithContainer(container containers.Container) Builder
	WithChannels(channels string) Builder
	Now() (Multiple, error)
}

// Multiple represents a multiple resource
type Multiple interface {
	Container() containers.Container
	HasChannels() bool
	Channels() string
}
