package suites

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
)

func mistakeToError(mistake mistakes.Mistake) string {
	content := mistake.Content()
	if content.IsTokenHasNoMatch() {
		return tokenHasNoMatch
	}

	if content.IsContainsValidNotMatch() {
		return containsValidNotMatch
	}

	if content.IsSpecifierDoNotMatch() {
		return specifierDoNotMatch
	}

	if content.IsContainsNextElement() {
		return containsNextElement
	}

	if content.IsCardinalityIsInvalid() {
		return cardinalityIsInvalid
	}

	// contains prefix:
	return containsPrefix
}
