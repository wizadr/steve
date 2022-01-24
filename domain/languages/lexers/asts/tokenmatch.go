package asts

type tokenMatch struct {
	token   string
	matches []Token
}

func createTokenMatch(
	token string,
	matches []Token,
) TokenMatch {
	out := tokenMatch{
		token:   token,
		matches: matches,
	}

	return &out
}

// Token returns the token
func (obj *tokenMatch) Token() string {
	return obj.token
}

// Discoveries returns the discoveries
func (obj *tokenMatch) Discoveries() []string {
	out := []string{}
	for _, oneMatch := range obj.matches {
		out = append(out, oneMatch.Discovery())
	}

	return out
}

// Matches returns the matches
func (obj *tokenMatch) Matches() []Token {
	return obj.matches
}
