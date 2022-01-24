package paths

import "github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a resource path builder
type Builder interface {
	Create() Builder
	WithRootToken(rootToken string) Builder
	WithContainer(container containers.Container) Builder
	Now() (Path, error)
}

// Path represents a path resource
type Path interface {
	RootToken() string
	Container() containers.Container
}
