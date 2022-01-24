package mistakes

import "github.com/steve-care-software/steve/domain/languages/lexers/asts/tokens"

type containsNextElement struct {
	nextElement tokens.NextElement
}

func createContainsNextElement(
	nextElement tokens.NextElement,
) ContainsNextElement {
	out := containsNextElement{
		nextElement: nextElement,
	}

	return &out
}

// NextElement returns the next element
func (obj *containsNextElement) NextElement() tokens.NextElement {
	return obj.nextElement
}
