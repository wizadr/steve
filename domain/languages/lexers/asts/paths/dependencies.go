package paths

import (
	"errors"
	"fmt"
)

type dependencies struct {
	list []Element
	mp   map[string]Element
}

func createDependencies(
	list []Element,
	mp map[string]Element,
) Dependencies {
	out := dependencies{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns all the dependencies
func (obj *dependencies) All() []Element {
	return obj.list
}

// IsEmpty returns true if there is no dependency, false otherwise
func (obj *dependencies) IsEmpty() bool {
	return len(obj.list) <= 0
}

// Fetch fetches a dependency by name
func (obj *dependencies) Fetch(name string) (Element, error) {
	if path, ok := obj.mp[name]; ok {
		return path, nil
	}

	str := fmt.Sprintf("there is no path dependency for token (name: %s)", name)
	return nil, errors.New(str)
}
