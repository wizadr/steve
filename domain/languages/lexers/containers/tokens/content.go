package tokens

import "github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"

type content struct {
	token TokenReference
	rule  rules.Rule
}

func createContentWithToken(
	token TokenReference,
) Content {
	return createContentInternally(token, nil)
}

func createContentWithRule(
	rule rules.Rule,
) Content {
	return createContentInternally(nil, rule)
}

func createContentInternally(
	token TokenReference,
	rule rules.Rule,
) Content {
	out := content{
		token: token,
		rule:  rule,
	}

	return &out
}

// Name returns the name
func (obj *content) Name() string {
	if obj.IsToken() {
		return obj.Token().Name()
	}

	return obj.Rule().Name()
}

// IsToken returns true if there is a token, false otherwise
func (obj *content) IsToken() bool {
	return obj.token != nil
}

// Token returns the token reference
func (obj *content) Token() TokenReference {
	return obj.token
}

// IsRule returns true if there is a rule, false otherwise
func (obj *content) IsRule() bool {
	return obj.rule != nil
}

// Rule returns the rule
func (obj *content) Rule() rules.Rule {
	return obj.rule
}
