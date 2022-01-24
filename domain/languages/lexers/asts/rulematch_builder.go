package asts

import "errors"

type ruleMatchBuilder struct {
	rule   string
	result Result
}

func createRuleMatchBuilder() RuleMatchBuilder {
	out := ruleMatchBuilder{
		rule:   "",
		result: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleMatchBuilder) Create() RuleMatchBuilder {
	return createRuleMatchBuilder()
}

// WithRule adds a rule to the builder
func (app *ruleMatchBuilder) WithRule(rule string) RuleMatchBuilder {
	app.rule = rule
	return app
}

// WithResult adds a result to the builder
func (app *ruleMatchBuilder) WithResult(result Result) RuleMatchBuilder {
	app.result = result
	return app
}

// Now builds a new RuleMatch instance
func (app *ruleMatchBuilder) Now() (RuleMatch, error) {
	if app.rule == "" {
		return nil, errors.New("the rule is mandatory in order to build a RuleMatch instance")
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build a RuleMatch instance")
	}

	return createRuleMatch(app.rule, app.result), nil
}
