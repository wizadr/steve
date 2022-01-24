package paths

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type recursiveToken struct {
	name        string
	cardinality cardinality.Cardinality
	specifiers  []Specifier
}

func createRecursiveToken(
	name string,
	cardinality cardinality.Cardinality,
) RecursiveToken {
	return createRecursiveTokenInternally(name, cardinality, nil)
}

func createRecursiveTokenWithSpecifiers(
	name string,
	cardinality cardinality.Cardinality,
	specifiers []Specifier,
) RecursiveToken {
	return createRecursiveTokenInternally(name, cardinality, specifiers)
}

func createRecursiveTokenInternally(
	name string,
	cardinality cardinality.Cardinality,
	specifiers []Specifier,
) RecursiveToken {
	out := recursiveToken{
		name:        name,
		cardinality: cardinality,
		specifiers:  specifiers,
	}

	return &out
}

// Name returns the name
func (obj *recursiveToken) Name() string {
	return obj.name
}

// Cardinality returns the cardinality
func (obj *recursiveToken) Cardinality() cardinality.Cardinality {
	return obj.cardinality
}

// HasSpecifiers returns true if there is specifiers, false otherwise
func (obj *recursiveToken) HasSpecifiers() bool {
	return obj.specifiers != nil
}

// Specifiers returns the specifiers, if any
func (obj *recursiveToken) Specifiers() []Specifier {
	return obj.specifiers
}
