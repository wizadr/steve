package asts

type ruleMatch struct {
	rule   string
	result Result
}

func createRuleMatch(
	rule string,
	result Result,
) RuleMatch {
	out := ruleMatch{
		rule:   rule,
		result: result,
	}

	return &out
}

// Rule returns the rule
func (obj *ruleMatch) Rule() string {
	return obj.rule
}

// Result returns the result
func (obj *ruleMatch) Result() Result {
	return obj.result
}
