package tokens

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
)

type ruleMatchResult struct {
	path   paths.Rule
	result Result
}

func createRuleMatchResult(
	path paths.Rule,
	result Result,
) RuleMatchResult {
	out := ruleMatchResult{
		path:   path,
		result: result,
	}

	return &out
}

// IsCardinalityValid returns true if the cardinality is valid, false otherwise
func (obj *ruleMatchResult) IsCardinalityValid() bool {
	amount := obj.result.Amount()
	return obj.Path().Cardinality().IsValid(amount)
}

// Path returns the path
func (obj *ruleMatchResult) Path() paths.Rule {
	return obj.path
}

// Result returns the result
func (obj *ruleMatchResult) Result() Result {
	return obj.result
}
