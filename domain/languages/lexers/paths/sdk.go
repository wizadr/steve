package paths

// NewBuilder createsa new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a paths builder
type Builder interface {
	Create() Builder
	WithRules(rules string) Builder
	WithTokens(tokens string) Builder
	WithChannels(channels string) Builder
	Now() (Paths, error)
}

// Paths represents token paths
type Paths interface {
	Rules() string
	Tokens() string
	HasChannels() bool
	Channels() string
}
