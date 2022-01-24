package tokens

type tokenReference struct {
	name       string
	specifiers TokenSpecifiers
}

func createTokenReference(
    name string,
) TokenReference {
    return createTokenReferenceInternally(name, nil)
}

func createTokenReferenceWithSpecifiers(
    name string,
    specifiers TokenSpecifiers,
) TokenReference {
    return createTokenReferenceInternally(name, specifiers)
}

func createTokenReferenceInternally(
	name string,
	specifiers TokenSpecifiers,
) TokenReference {
	out := tokenReference{
		name:       name,
		specifiers: specifiers,
	}

	return &out
}

// Name returns the name
func (obj *tokenReference) Name() string {
	return obj.name
}

// HasSpecifiers returns true if there is specifiers, false otherwise
func (obj *tokenReference) HasSpecifiers() bool {
	return obj.specifiers != nil
}

// Specifiers returns the specifiers, if any
func (obj *tokenReference) Specifiers() TokenSpecifiers {
	return obj.specifiers
}
