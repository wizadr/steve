package mistakes

type pathContainer struct {
	rule  string
	token string
}

func createPathContainerWithRule(
	rule string,
) PathContainer {
	return createPathContainerInternally(rule, "")
}

func createPathContainerWithToken(
	token string,
) PathContainer {
	return createPathContainerInternally("", token)
}

func createPathContainerInternally(
	rule string,
	token string,
) PathContainer {
	out := pathContainer{
		rule:  rule,
		token: token,
	}

	return &out
}

// IsRule returns true if there is a rule, false otherwise
func (obj *pathContainer) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *pathContainer) Rule() string {
	return obj.rule
}

// IsToken returns true if there is a token, false otherwise
func (obj *pathContainer) IsToken() bool {
	return obj.token != ""
}

// Token returns the token, if any
func (obj *pathContainer) Token() string {
	return obj.token
}
