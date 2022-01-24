package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
)

type ruleMatchBuilder struct {
	rule             paths.Rule
	result           RuleMatchResult
	canContainPrefix bool
}

func createRuleMatchBuilder() RuleMatchBuilder {
	out := ruleMatchBuilder{
		rule:             nil,
		result:           nil,
		canContainPrefix: false,
	}

	return &out
}

// Create initializes the builder
func (app *ruleMatchBuilder) Create() RuleMatchBuilder {
	return createRuleMatchBuilder()
}

// CanContainPrefix flags the builder as having a potential prefix
func (app *ruleMatchBuilder) CanContainPrefix() RuleMatchBuilder {
	app.canContainPrefix = true
	return app
}

// WithRule adds a rule to the builder
func (app *ruleMatchBuilder) WithRule(rule paths.Rule) RuleMatchBuilder {
	app.rule = rule
	return app
}

// WithResult adds a result to the builder
func (app *ruleMatchBuilder) WithResult(result RuleMatchResult) RuleMatchBuilder {
	app.result = result
	return app
}

// Now builds a new RuleMatch instance
func (app *ruleMatchBuilder) Now() (RuleMatch, error) {
	if app.rule == nil {
		return nil, errors.New("the rule is mandatory in order to build a RuleMatch instance")
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build a RuleMatch instance")
	}

	return createRuleMatch(app.rule, app.result, app.canContainPrefix), nil
}
