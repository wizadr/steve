package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
)

type ruleMatchResultBuilder struct {
	path   paths.Rule
	result Result
}

func createRuleMatchResultBuilder() RuleMatchResultBuilder {
	out := ruleMatchResultBuilder{
		path:   nil,
		result: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleMatchResultBuilder) Create() RuleMatchResultBuilder {
	return createRuleMatchResultBuilder()
}

// WithPath adds a path to the builder
func (app *ruleMatchResultBuilder) WithPath(path paths.Rule) RuleMatchResultBuilder {
	app.path = path
	return app
}

// WithResult adds a result to the builder
func (app *ruleMatchResultBuilder) WithResult(result Result) RuleMatchResultBuilder {
	app.result = result
	return app
}

// Now builds a new RuleMatchResult instance
func (app *ruleMatchResultBuilder) Now() (RuleMatchResult, error) {
	if app.path == nil {
		return nil, errors.New("the rule path is mandatory in order to build a RuleMatchResult instance")
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build a RuleMatchResult instance")
	}

	return createRuleMatchResult(app.path, app.result), nil
}
