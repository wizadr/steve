package paths

type tokenPath struct {
	element    Element
	specifiers []Specifier
}

func createTokenPath(
	element Element,
) TokenPath {
	return createTokenPathInternally(element, nil)
}

func createTokenPathWithSpecifiers(
	element Element,
	specifiers []Specifier,
) TokenPath {
	return createTokenPathInternally(element, specifiers)
}

func createTokenPathInternally(
	element Element,
	specifiers []Specifier,
) TokenPath {
	out := tokenPath{
		element:    element,
		specifiers: specifiers,
	}

	return &out
}

// Element returns the token element
func (obj *tokenPath) Element() Element {
	return obj.element
}

// HasSpecifiers returns true if there is specifiers, false otherwise
func (obj *tokenPath) HasSpecifiers() bool {
	return obj.specifiers != nil
}

// Specifiers returns the specifiers, if any
func (obj *tokenPath) Specifiers() []Specifier {
	return obj.specifiers
}
