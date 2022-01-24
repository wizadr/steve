package tokens

type matchContent struct {
	token TokenMatch
	rule  RuleMatch
	skip  SkipMatch
	line  LineMatch
}

func createMatchContentWithToken(
	token TokenMatch,
) MatchContent {
	return createMatchContentInternally(token, nil, nil, nil)
}

func createMatchContentWithRule(
	rule RuleMatch,
) MatchContent {
	return createMatchContentInternally(nil, rule, nil, nil)
}

func createMatchContentWithSkip(
	skip SkipMatch,
) MatchContent {
	return createMatchContentInternally(nil, nil, skip, nil)
}

func createMatchContentWithLine(
	line LineMatch,
) MatchContent {
	return createMatchContentInternally(nil, nil, nil, line)
}

func createMatchContentInternally(
	token TokenMatch,
	rule RuleMatch,
	skip SkipMatch,
	line LineMatch,
) MatchContent {
	out := matchContent{
		token: token,
		rule:  rule,
		skip:  skip,
		line:  line,
	}

	return &out
}

// IsToken returns true if there is a token, false otherwise
func (obj *matchContent) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *matchContent) Token() TokenMatch {
	return obj.token
}

// IsRule returns true if there is a rule, false otherwise
func (obj *matchContent) IsRule() bool {
	return obj.rule != nil
}

// Rule returns the rule, if any
func (obj *matchContent) Rule() RuleMatch {
	return obj.rule
}

// IsSkip returns true if there is a skip, false otherwise
func (obj *matchContent) IsSkip() bool {
	return obj.skip != nil
}

// Skip returns the skip, if any
func (obj *matchContent) Skip() SkipMatch {
	return obj.skip
}

// IsLine returns true if there is a line, false otherwise
func (obj *matchContent) IsLine() bool {
	return obj.line != nil
}

// Line returns the line, if any
func (obj *matchContent) Line() LineMatch {
	return obj.line
}
