package paths

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type container struct {
	token     Token
	rule      Rule
	recursive RecursiveToken
	lines     []Line
}

func createContainerWithToken(
	token Token,
) Container {
	return createContainerInternally(token, nil, nil, nil)
}

func createContainerWithRule(
	rule Rule,
) Container {
	return createContainerInternally(nil, rule, nil, nil)
}

func createContainerWithRecursiveToken(
	recursive RecursiveToken,
) Container {
	return createContainerInternally(nil, nil, recursive, nil)
}

func createContainerWithLines(
	lines []Line,
) Container {
	return createContainerInternally(nil, nil, nil, lines)
}

func createContainerInternally(
	token Token,
	rule Rule,
	recursive RecursiveToken,
	lines []Line,
) Container {
	out := container{
		token:     token,
		rule:      rule,
		recursive: recursive,
		lines:     lines,
	}

	return &out
}

// Cardinality returns the cardinality
func (obj *container) Cardinality() cardinality.Cardinality {
	if obj.IsToken() {
		return obj.token.Cardinality()
	}

	if obj.IsRule() {
		return obj.rule.Cardinality()
	}

	return obj.recursive.Cardinality()
}

// IsToken returns true if there is a token, false otherwise
func (obj *container) IsToken() bool {
	return obj.token != nil
}

// Token returns a token, if any
func (obj *container) Token() Token {
	return obj.token
}

// IsRule returns true if there is a rule, false otherwise
func (obj *container) IsRule() bool {
	return obj.rule != nil
}

// Rule returns a rule, if any
func (obj *container) Rule() Rule {
	return obj.rule
}

// IsRecursive returns true if there is a recursiveToken, false otherwise
func (obj *container) IsRecursive() bool {
	return obj.recursive != nil
}

// Recursive returns a recursiveToken, if any
func (obj *container) Recursive() RecursiveToken {
	return obj.recursive
}

// IsLines returns true if there is lines, false otherwise
func (obj *container) IsLines() bool {
	return obj.lines != nil
}

// Lines returns the lines, if any
func (obj *container) Lines() []Line {
	return obj.lines
}
