package results

import "github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"

type result struct {
	input   rules.Rule
	matches Matches
}

func createResult(
	input rules.Rule,
) Result {
	return createResultInternally(input, nil)
}

func createResultWithMatches(
	input rules.Rule,
	matches Matches,
) Result {
	return createResultInternally(input, matches)
}

func createResultInternally(
	input rules.Rule,
	matches Matches,
) Result {
	out := result{
		input:   input,
		matches: matches,
	}

	return &out
}

// Input returns the input
func (obj *result) Input() rules.Rule {
	return obj.input
}

// Content returns the content
func (obj *result) Content() string {
	if obj.HasMatches() {
		return obj.matches.Content()
	}

	return ""
}

// Amount returns the amount
func (obj *result) Amount() uint {
	if obj.HasMatches() {
		return obj.matches.Amount()
	}

	return uint(0)
}

// HasMatches returns true if there is matches, false otherwise
func (obj *result) HasMatches() bool {
	return obj.matches != nil
}

// Matches returns the matches
func (obj *result) Matches() Matches {
	return obj.matches
}
