package tokens

import "github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"

type tokenSpecifierContent struct {
	token string
	rule  rules.Rule
}

func createTokenSpecifierContentWithToken(
	token string,
) TokenSpecifierContent {
	return createTokenSpecifierContentInternally(token, nil)
}

func createTokenSpecifierContentWithRule(
	rule rules.Rule,
) TokenSpecifierContent {
	return createTokenSpecifierContentInternally("", rule)
}

func createTokenSpecifierContentInternally(
	token string,
	rule rules.Rule,
) TokenSpecifierContent {
	out := tokenSpecifierContent{
		token: token,
		rule:  rule,
	}

	return &out
}

// Name returns the name of the content
func (obj *tokenSpecifierContent) Name() string {
	if obj.IsToken() {
		return obj.token
	}

	return obj.rule.Name()
}

// IsToken returns true if there is a token, false otherwise
func (obj *tokenSpecifierContent) IsToken() bool {
	return obj.token != ""
}

// Token returns the token, if any
func (obj *tokenSpecifierContent) Token() string {
	return obj.token
}

// IsRule returns true if there is a rule, false otherwise
func (obj *tokenSpecifierContent) IsRule() bool {
	return obj.rule != nil
}

// Rule returns the rule, if any
func (obj *tokenSpecifierContent) Rule() rules.Rule {
	return obj.rule
}
