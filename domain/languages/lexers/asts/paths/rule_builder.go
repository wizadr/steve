package paths

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/cardinality"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

type ruleBuilder struct {
	base        rules.Rule
	cardinality cardinality.Cardinality
}

func createRuleBuilder() RuleBuilder {
	out := ruleBuilder{
		base:        nil,
		cardinality: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleBuilder) Create() RuleBuilder {
	return createRuleBuilder()
}

// WithBase adds a base rule to the builder
func (app *ruleBuilder) WithBase(base rules.Rule) RuleBuilder {
	app.base = base
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *ruleBuilder) WithCardinality(cardinality cardinality.Cardinality) RuleBuilder {
	app.cardinality = cardinality
	return app
}

// Now builds a new Rule instance
func (app *ruleBuilder) Now() (Rule, error) {
	if app.base == nil {
		return nil, errors.New("the base rule is mandatory in order to build a Rule instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a Rule instance")
	}

	return createRule(app.base, app.cardinality), nil
}
