package tokens

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type nextElement struct {
	constant string
	pattern  patterns.Pattern
}

func createNextElementWithConstant(
	constant string,
) NextElement {
	return createNextElementInternally(constant, nil)
}

func createNextElementWithPattern(
	pattern patterns.Pattern,
) NextElement {
	return createNextElementInternally("", pattern)
}

func createNextElementInternally(
	constant string,
	pattern patterns.Pattern,
) NextElement {
	out := nextElement{
		constant: constant,
		pattern:  pattern,
	}

	return &out
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *nextElement) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *nextElement) Constant() string {
	return obj.constant
}

// IsPattern returns true if there is a pattern, false otherwise
func (obj *nextElement) IsPattern() bool {
	return obj.pattern != nil
}

// Pattern returns the pattern, if any
func (obj *nextElement) Pattern() patterns.Pattern {
	return obj.pattern
}
