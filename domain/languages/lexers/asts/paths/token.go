package paths

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type token struct {
	path        TokenPath
	cardinality cardinality.Cardinality
}

func createToken(
	path TokenPath,
	cardinality cardinality.Cardinality,
) Token {
	out := token{
		path:        path,
		cardinality: cardinality,
	}

	return &out
}

// Path returns the path
func (obj *token) Path() TokenPath {
	return obj.path
}

// Cardinality returns the cardinality
func (obj *token) Cardinality() cardinality.Cardinality {
	return obj.cardinality
}
