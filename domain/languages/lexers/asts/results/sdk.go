package results

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	patternAdapter := patterns.NewAdapter()
	groupBuilder := patterns.NewGroupBuilder()
	builder := NewBuilder()
	return createAdapter(patternAdapter, groupBuilder, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a rule adapter
type Adapter interface {
	ToResult(input rules.Rule, script string) (Result, error)
}

// Builder represents a result builder
type Builder interface {
	Create() Builder
	WithInput(input rules.Rule) Builder
	WithMatches(matches []patterns.Result) Builder
	Now() (Result, error)
}

// Result represents a result
type Result interface {
	Amount() uint
	Input() rules.Rule
	Content() string
	HasMatches() bool
	Matches() Matches
}

// Matches represents a rule matches
type Matches interface {
	Content() string
	Amount() uint
	List() []patterns.Result
}
