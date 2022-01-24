package rules

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type pattern struct {
	subPatterns patterns.Patterns
	cardinality cardinality.Cardinality
}

func createPattern(
	subPatterns patterns.Patterns,
	cardinality cardinality.Cardinality,
) Pattern {
	out := pattern{
		subPatterns: subPatterns,
		cardinality: cardinality,
	}

	return &out
}

// SubPatterns returns the sub patterns
func (obj *pattern) SubPatterns() patterns.Patterns {
	return obj.subPatterns
}

// Cardinality returns the cardinality
func (obj *pattern) Cardinality() cardinality.Cardinality {
	return obj.cardinality
}
