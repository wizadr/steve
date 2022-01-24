package mistakes

type content struct {
	isTokenHasNoMatch     bool
	containsValidNotMatch ContainsValidNotMatch
	specifierDoNotMatch   SpecifierDoNotMatch
	containsNextElement   ContainsNextElement
	cardinalityIsInvalid  CardinalityIsInvalid
	containsPrefix        ContainsPrefix
}

func createContentWithTokenHasNoMatch() Content {
	return createContentInternally(
		true,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createContentWithContainsValidNotMatch(
	containsValidNotMatch ContainsValidNotMatch,
) Content {
	return createContentInternally(
		false,
		containsValidNotMatch,
		nil,
		nil,
		nil,
		nil,
	)
}

func createContentWithSpecifierDoNotMatch(
	specifierDoNotMatch SpecifierDoNotMatch,
) Content {
	return createContentInternally(
		false,
		nil,
		specifierDoNotMatch,
		nil,
		nil,
		nil,
	)
}

func createContentWithContainsNextElement(
	containsNextElement ContainsNextElement,
) Content {
	return createContentInternally(
		false,
		nil,
		nil,
		containsNextElement,
		nil,
		nil,
	)
}

func createContentWithCardinalityIsInvalid(
	cardinalityIsInvalid CardinalityIsInvalid,
) Content {
	return createContentInternally(
		false,
		nil,
		nil,
		nil,
		cardinalityIsInvalid,
		nil,
	)
}

func createContentWithContainsPrefix(
	containsPrefix ContainsPrefix,
) Content {
	return createContentInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		containsPrefix,
	)
}

func createContentInternally(
	isTokenHasNoMatch bool,
	containsValidNotMatch ContainsValidNotMatch,
	specifierDoNotMatch SpecifierDoNotMatch,
	containsNextElement ContainsNextElement,
	cardinalityIsInvalid CardinalityIsInvalid,
	containsPrefix ContainsPrefix,
) Content {
	out := content{
		isTokenHasNoMatch:     isTokenHasNoMatch,
		containsValidNotMatch: containsValidNotMatch,
		specifierDoNotMatch:   specifierDoNotMatch,
		containsNextElement:   containsNextElement,
		cardinalityIsInvalid:  cardinalityIsInvalid,
		containsPrefix:        containsPrefix,
	}

	return &out
}

// IsTokenHasNoMatch returns true if the mistake is tokenHasNoMatch, false otherwise
func (obj *content) IsTokenHasNoMatch() bool {
	return obj.isTokenHasNoMatch
}

// IsContainsValidNotMatch returns true if the mistake is containsValidNotMatch, false otherwise
func (obj *content) IsContainsValidNotMatch() bool {
	return obj.containsValidNotMatch != nil
}

// ContainsValidNotMatch returns the containsValidNotMatch, if any
func (obj *content) ContainsValidNotMatch() ContainsValidNotMatch {
	return obj.containsValidNotMatch
}

// IsSpecifierDoNotMatch returns true if the mistake is specifierDoNotMatch, false otherwise
func (obj *content) IsSpecifierDoNotMatch() bool {
	return obj.specifierDoNotMatch != nil
}

// SpecifierDoNotMatch returns the specifierDoNotMatch, if any
func (obj *content) SpecifierDoNotMatch() SpecifierDoNotMatch {
	return obj.specifierDoNotMatch
}

// IsContainsNextElement returns true if the mistake is containsNextElement, false otherwise
func (obj *content) IsContainsNextElement() bool {
	return obj.containsNextElement != nil
}

// ContainsNextElement returns the containsNextElement, if any
func (obj *content) ContainsNextElement() ContainsNextElement {
	return obj.containsNextElement
}

// IsCardinalityIsInvalid returns true if the mistake is cardinalityIsInvalid, false otherwise
func (obj *content) IsCardinalityIsInvalid() bool {
	return obj.cardinalityIsInvalid != nil
}

// CardinalityIsInvalid returns the containsNextElement, if any
func (obj *content) CardinalityIsInvalid() CardinalityIsInvalid {
	return obj.cardinalityIsInvalid
}

// IsContainsPrefix returns true if the mistake is containsPrefix, false otherwise
func (obj *content) IsContainsPrefix() bool {
	return obj.containsPrefix != nil
}

// ContainsPrefix returns the containsPrefix, if any
func (obj *content) ContainsPrefix() ContainsPrefix {
	return obj.containsPrefix
}
