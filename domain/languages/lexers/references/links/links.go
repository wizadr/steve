package links

import (
	"errors"
	"fmt"
)

type links struct {
	list          []Link
	mpLocalTokens map[string]Link
	mpRefTokens   map[string]map[string]Link
}

func createLinks(
	list []Link,
	mpLocalTokens map[string]Link,
	mpRefTokens map[string]map[string]Link,
) Links {
	out := links{
		list:          list,
		mpLocalTokens: mpLocalTokens,
		mpRefTokens:   mpRefTokens,
	}

	return &out
}

// All returns the references
func (obj *links) All() []Link {
	return obj.list
}

// FetchByLocalToken fetches the reference by local token, if any
func (obj *links) FetchByLocalToken(localToken string) (Link, error) {
	if ref, ok := obj.mpLocalTokens[localToken]; ok {
		return ref, nil
	}

	str := fmt.Sprintf("the local token (%s) is not present in any of the Link instances", localToken)
	return nil, errors.New(str)
}

// FetchByReferenceToken fetches the reference by reference token, if any
func (obj *links) FetchByReferenceToken(includeName string, refToken string) (Link, error) {
	if list, ok := obj.mpRefTokens[includeName]; ok {
		if ref, ok := list[refToken]; ok {
			return ref, nil
		}

		str := fmt.Sprintf("the include (name: %s) does not contain the requested reference token (name: %s)", includeName, refToken)
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the include (name: %s) is not present in any of the Link instances", includeName)
	return nil, errors.New(str)
}
