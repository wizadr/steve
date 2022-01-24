package tokens

import (
	"errors"
	"fmt"
)

type tokenSpecifiers struct {
	list []TokenSpecifier
	mp   map[string]TokenSpecifier
}

func createTokenSpecifiers(
	list []TokenSpecifier,
	mp map[string]TokenSpecifier,
) TokenSpecifiers {
	out := tokenSpecifiers{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns the sub elements
func (obj *tokenSpecifiers) All() []TokenSpecifier {
	return obj.list
}

// Find returns the sub element by name
func (obj *tokenSpecifiers) Find(name string) (TokenSpecifier, error) {
	if el, ok := obj.mp[name]; ok {
		return el, nil
	}

	str := fmt.Sprintf("the tokenSpecifier (name: %s) is not declared", name)
	return nil, errors.New(str)
}
