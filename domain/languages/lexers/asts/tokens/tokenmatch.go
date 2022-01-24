package tokens

import "github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"

type tokenMatch struct {
	path       paths.TokenPath
	result     TokenResult
	specifiers []Specifier
}

func createTokenMatch(
	path paths.TokenPath,
	result TokenResult,
) TokenMatch {
	return createTokenMatchInternally(path, result, nil)
}

func createTokenMatchWithSpecifiers(
	path paths.TokenPath,
	result TokenResult,
	specifiers []Specifier,
) TokenMatch {
	return createTokenMatchInternally(path, result, specifiers)
}

func createTokenMatchInternally(
	path paths.TokenPath,
	result TokenResult,
	specifiers []Specifier,
) TokenMatch {
	out := tokenMatch{
		path:       path,
		result:     result,
		specifiers: specifiers,
	}

	return &out
}

// IsValid returns true if the match is valid, false otherwise
func (obj *tokenMatch) IsValid() bool {
	if obj.HasSpecifiers() {
		for _, oneSpecifier := range obj.specifiers {
			if !oneSpecifier.IsValid() {
				return false
			}
		}
	}

	return obj.Result().IsValid()
}

// IsExact returns true if the match is exact, false otherwise
func (obj *tokenMatch) IsExact() bool {
	return obj.Result().IsExact()
}

// Path returns the path
func (obj *tokenMatch) Path() paths.TokenPath {
	return obj.path
}

// Result returns the result, if any
func (obj *tokenMatch) Result() TokenResult {
	return obj.result
}

// HasSpecifiers returns true if there is specifiers, false otherwise
func (obj *tokenMatch) HasSpecifiers() bool {
	return obj.specifiers != nil
}

// Specifiers returns the specifiers, if any
func (obj *tokenMatch) Specifiers() []Specifier {
	return obj.specifiers
}
