package tokens

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type tokenResult struct {
	input       string
	cardinality cardinality.Cardinality
	matches     TokenResultMatches
}

func createTokenResult(
	input string,
	cardinality cardinality.Cardinality,
) TokenResult {
	return createTokenResultInternally(input, cardinality, nil)
}

func createTokenResultWithMatches(
	input string,
	cardinality cardinality.Cardinality,
	matches TokenResultMatches,
) TokenResult {
	return createTokenResultInternally(input, cardinality, matches)
}

func createTokenResultInternally(
	input string,
	cardinality cardinality.Cardinality,
	matches TokenResultMatches,
) TokenResult {
	out := tokenResult{
		input:       input,
		cardinality: cardinality,
		matches:     matches,
	}

	return &out
}

// IsCardinalityValid returns true if the cardinality is valid, false otherwise
func (obj *tokenResult) IsCardinalityValid() bool {
	amount := obj.Amount()
	return obj.Cardinality().IsValid(amount)
}

// Cardinality returns the cardinality
func (obj *tokenResult) Cardinality() cardinality.Cardinality {
	return obj.cardinality
}

// Amount retruns the amount of results
func (obj *tokenResult) Amount() uint {
	if !obj.HasMatches() {
		return 0
	}

	return obj.Matches().Amount()
}

// IsValid returns true if the result is valid, false otherwise
func (obj *tokenResult) IsValid() bool {
	if !obj.IsCardinalityValid() {
		return false
	}

	// the cardinality is valid, so if there is no match, return valid:
	if !obj.HasMatches() {
		return true
	}

	return obj.Matches().IsValid()
}

// IsExact returns true if the result is exact, false otherwise
func (obj *tokenResult) IsExact() bool {
	if !obj.HasMatches() {
		return false
	}

	return obj.Matches().IsExact()
}

// Input returns the input
func (obj *tokenResult) Input() string {
	return obj.input
}

// Discoveries returns the matched discoveries
func (obj *tokenResult) Discoveries() string {
	if !obj.HasMatches() {
		return ""
	}

	return obj.matches.Discoveries()
}

// HasMatches returns true if there is matches, false otherwise
func (obj *tokenResult) HasMatches() bool {
	return obj.matches != nil
}

// Matches returns the matches, if any
func (obj *tokenResult) Matches() TokenResultMatches {
	return obj.matches
}
