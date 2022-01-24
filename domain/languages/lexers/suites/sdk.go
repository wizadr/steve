package suites

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

const tokenHasNoMatch = "tokenHasNoMatch"
const containsValidNotMatch = "containsValidNotMatch"
const specifierDoNotMatch = "specifierDoNotMatch"
const containsNextElement = "containsNextElement"
const cardinalityIsInvalid = "cardinalityIsInvalid"
const containsPrefix = "containsPrefix"

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder() ApplicationBuilder {
	builder := NewBuilder()
	lineBuilder := NewLineBuilder()
	astAdapterBuilder := asts.NewAdapterBuilder()
	return createApplicationBuilder(
		builder,
		lineBuilder,
		astAdapterBuilder,
	)
}

// NewBuilder creates a new suite builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithTokens(tokens tokens.Tokens) ApplicationBuilder
	WithChannels(channels tokens.Tokens) ApplicationBuilder
	WithTokenNotFoundFunc(tokenNotFoundFn paths.FetchTokenNotFoundFn) ApplicationBuilder
	WithTokenReplacementFunc(tokenReplacementFn paths.FetchElementReplacementFn) ApplicationBuilder
	Now() (Application, error)
}

// Application represents a test suite application
type Application interface {
	Execute() ([]Suite, error)
}

// Builder represents an executed test suite builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithLines(lines []Line) Builder
	Now() (Suite, error)
}

// Suite represents an executed test suite
type Suite interface {
	Name() string
	Lines() []Line
}

// LineBuilder represents an executed test line builder
type LineBuilder interface {
	Create() LineBuilder
	WithIndex(index uint) LineBuilder
	IsSuccessful() LineBuilder
	Now() (Line, error)
}

// Line represents an executed test line
type Line interface {
	Index() uint
	IsSuccess() bool
}
