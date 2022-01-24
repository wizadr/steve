package tokens

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
)

type ruleMatch struct {
	rule          paths.Rule
	result        RuleMatchResult
	canHavePrefix bool
}

func createRuleMatch(
	rule paths.Rule,
	result RuleMatchResult,
	canHavePrefix bool,
) RuleMatch {
	return createRuleMatchInternally(rule, result, canHavePrefix)
}

func createRuleMatchInternally(
	rule paths.Rule,
	result RuleMatchResult,
	canHavePrefix bool,
) RuleMatch {
	out := ruleMatch{
		rule:          rule,
		result:        result,
		canHavePrefix: canHavePrefix,
	}

	return &out
}

// IsValid returns true if the ruleMatch is valid
func (obj *ruleMatch) IsValid() bool {
	// to be valid, the cardinality must be valid:
	if !obj.result.IsCardinalityValid() {
		return false
	}

	// to be valid, there must be an exact match:
	return obj.IsExact()
}

// IsExact returns true if the match is exact, false otherwise
func (obj *ruleMatch) IsExact() bool {
	// if there is no match, the cardinality must be valid at zero:
	if !obj.result.Result().HasResults() {
		return obj.result.Path().Cardinality().IsValid(0)
	}

	// if we can't have a prefix, make sure we don't have one:
	if !obj.canHavePrefix {
		result := obj.Result().Result()
		if result.HasResults() {
			results := result.Results()
			if results[0].HasMatches() {
				list := results[0].Matches().List()
				return list[0].Discoveries().Index() <= 0
			}

			return false
		}

		return false
	}

	// this is an exact match:
	return true
}

// CanHavePrefix returns true if the ruleMatch can have prefix to be valid and exact, false otherwise
func (obj *ruleMatch) CanHavePrefix() bool {
	return obj.canHavePrefix
}

// Rule returns the rule
func (obj *ruleMatch) Rule() paths.Rule {
	return obj.rule
}

// Result returns the result, if any
func (obj *ruleMatch) Result() RuleMatchResult {
	return obj.result
}
