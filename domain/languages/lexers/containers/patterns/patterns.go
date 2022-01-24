package patterns

import (
	"errors"
	"fmt"
)

type patterns struct {
	list []Pattern
	mp   map[string]Pattern
}

func createPatterns(
	list []Pattern,
	mp map[string]Pattern,
) Patterns {
	out := patterns{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns the patterns
func (obj *patterns) All() []Pattern {
	return obj.list
}

// First returns the first pattern
func (obj *patterns) First() Pattern {
	return obj.list[0]
}

// Find finds a pattern by name
func (obj *patterns) Find(name string) (Pattern, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("there is no Pattern with the given name: %s", name)
	return nil, errors.New(str)
}
