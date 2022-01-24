package singles

import "github.com/steve-care-software/steve/domain/languages/lexers/resources/multiples"

type single struct {
	rootToken string
	multiple  multiples.Multiple
}

func createSingle(
	rootToken string,
	multiple multiples.Multiple,
) Single {
	out := single{
		rootToken: rootToken,
		multiple:  multiple,
	}

	return &out
}

// RootToken returns the rootToken
func (obj *single) RootToken() string {
	return obj.rootToken
}

// Multiple returns the multiple
func (obj *single) Multiple() multiples.Multiple {
	return obj.multiple
}
