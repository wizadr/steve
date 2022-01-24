package tokens

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type tokenSpecifier struct {
	content     TokenSpecifierContent
	cardinality cardinality.Specific
}

func createTokenSpecifier(
	content TokenSpecifierContent,
	cardinality cardinality.Specific,
) TokenSpecifier {
	out := tokenSpecifier{
		content:     content,
		cardinality: cardinality,
	}

	return &out
}

// Content returns the content
func (obj *tokenSpecifier) Content() TokenSpecifierContent {
	return obj.content
}

// Cardinality returns the cardinality
func (obj *tokenSpecifier) Cardinality() cardinality.Specific {
	return obj.cardinality
}
