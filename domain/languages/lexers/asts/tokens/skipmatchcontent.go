package tokens

type skipMatchContent struct {
	token TokenMatch
	rule  RuleMatch
}

func createSkipMatchContentWithToken(
	token TokenMatch,
) SkipMatchContent {
	return createSkipMatchContentInternally(token, nil)
}

func createSkipMatchContentWithRule(
	rule RuleMatch,
) SkipMatchContent {
	return createSkipMatchContentInternally(nil, rule)
}

func createSkipMatchContentInternally(
	token TokenMatch,
	rule RuleMatch,
) SkipMatchContent {
	out := skipMatchContent{
		token: token,
		rule:  rule,
	}

	return &out
}

// IsToken returns true if there is a token, false otherwise
func (obj *skipMatchContent) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *skipMatchContent) Token() TokenMatch {
	return obj.token
}

// IsRule returns true if there is a rule, false otherwise
func (obj *skipMatchContent) IsRule() bool {
	return obj.rule != nil
}

// Rule returns the rule, if any
func (obj *skipMatchContent) Rule() RuleMatch {
	return obj.rule
}
