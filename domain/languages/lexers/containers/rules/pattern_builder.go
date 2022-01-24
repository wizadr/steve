package rules

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type patternBuilder struct {
	subPatterns patterns.Patterns
	cardinality cardinality.Cardinality
}

func createPatternBuilder() PatternBuilder {
	out := patternBuilder{
		subPatterns: nil,
		cardinality: nil,
	}

	return &out
}

// Create initializes the builder
func (app *patternBuilder) Create() PatternBuilder {
	return createPatternBuilder()
}

// WithSubPatterns add sub patterns to the builder
func (app *patternBuilder) WithSubPatterns(subPatterns patterns.Patterns) PatternBuilder {
	app.subPatterns = subPatterns
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *patternBuilder) WithCardinality(cardinality cardinality.Cardinality) PatternBuilder {
	app.cardinality = cardinality
	return app
}

// Now builds a new Pattern instance
func (app *patternBuilder) Now() (Pattern, error) {
	if app.subPatterns == nil {
		return nil, errors.New("the sub Patterns are mandatory in order to build a Pattern instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a Pattern instance")
	}

	return createPattern(app.subPatterns, app.cardinality), nil
}
