package containers

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a container builder
type Builder interface {
	Create() Builder
	WithPatterns(patterns string) Builder
	WithRules(rules string) Builder
	WithTokens(tokens string) Builder
	Now() (Container, error)
}

// Container represents a container resource
type Container interface {
	Patterns() string
	Rules() string
	Tokens() string
}
