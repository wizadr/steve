package paths

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

type rule struct {
	base        rules.Rule
	cardinality cardinality.Cardinality
}

func createRule(
	base rules.Rule,
	cardinality cardinality.Cardinality,
) Rule {
	out := rule{
		base:        base,
		cardinality: cardinality,
	}

	return &out
}

// Base returns the base rule
func (obj *rule) Base() rules.Rule {
	return obj.base
}

// Cardinality retruns the cardinality
func (obj *rule) Cardinality() cardinality.Cardinality {
	return obj.cardinality
}
