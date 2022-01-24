package asts

type match struct {
	token TokenMatch
	rule  RuleMatch
}

func createMatchWithToken(
	token TokenMatch,
) Match {
	return createMatchInternally(token, nil)
}

func createMatchWithRule(
	rule RuleMatch,
) Match {
	return createMatchInternally(nil, rule)
}

func createMatchInternally(
	token TokenMatch,
	rule RuleMatch,
) Match {
	out := match{
		token: token,
		rule:  rule,
	}

	return &out
}

// IsToken returns true if there is a token match, false otherwise
func (obj *match) IsToken() bool {
	return obj.token != nil
}

// Token returns the token match, if any
func (obj *match) Token() TokenMatch {
	return obj.token
}

// IsRule returns true if there is a rule match, false otherwise
func (obj *match) IsRule() bool {
	return obj.rule != nil
}

// Rule returns the rule match, if any
func (obj *match) Rule() RuleMatch {
	return obj.rule
}
