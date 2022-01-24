package asts

import "strings"

type token struct {
	name  string
	match LineMatch
}

func createToken(
	name string,
	match LineMatch,
) Token {
	return createTokenInternally(name, match)
}

func createTokenInternally(
	name string,
	match LineMatch,
) Token {
	out := token{
		name:  name,
		match: match,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Match returns the match
func (obj *token) Match() LineMatch {
	return obj.match
}

// Prefix returns the prefix
func (obj *token) Prefix() string {
	/*if !obj.match.HasMatches() {
		return ""
	}

	matches := obj.match.Matches()
	firstMatch := matches[0]
	if firstMatch.IsToken() {
		tokenMatches := firstMatch.Token().Matches()
		return tokenMatches[0].Prefix()
	}

	return firstMatch.Rule().Result().Prefix()*/
	return ""
}

// Discovery returns the discovery
func (obj *token) Discovery() string {
	if !obj.match.HasMatches() {
		return ""
	}

	discoveries := []string{}
	matches := obj.match.Matches()
	for _, oneMatch := range matches {
		if oneMatch.IsToken() {
			tokenDiscoveries := oneMatch.Token().Discoveries()
			discoveries = append(discoveries, tokenDiscoveries...)
		}

		if oneMatch.IsRule() {
			ruleDiscoveriesContent := oneMatch.Rule().Result().Discoveries().Content()
			discoveries = append(discoveries, ruleDiscoveriesContent)
		}
	}

	return strings.Join(discoveries, "")
}

// Suffix returns the suffix, if any
func (obj *token) Suffix() string {
	return ""
}
