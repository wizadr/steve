package tokens

import (
	"strings"
)

type matches struct {
	list []BlockMatch
}

func createMatches(
	list []BlockMatch,
) Matches {
	out := matches{
		list: list,
	}

	return &out
}

// BlockMatches return the []BlockMatch, if any
func (obj *matches) BlockMatches() []BlockMatch {
	return obj.list
}

// Input returns the input
func (obj *matches) Input() string {
	list := []string{}
	for _, oneMatch := range obj.list {
		list = append(list, oneMatch.Input())
	}

	return strings.Join(list, "")
}

// IsValid returns true if the matches is valid, false otherwise
func (obj *matches) IsValid() bool {
	for _, oneMatch := range obj.list {
		if !oneMatch.IsValid() {
			return false
		}
	}

	return true
}

// IsExact returns true if the matches is exact, false otherwise
func (obj *matches) IsExact() bool {
	for _, oneMatch := range obj.list {
		if !oneMatch.IsExact() {
			return false
		}
	}

	return true
}

// Discoveries returns the matched discoveries
func (obj *matches) Discoveries() string {
	if !obj.IsValid() {
		return ""
	}

	list := []string{}
	for _, oneMatch := range obj.list {
		list = append(list, oneMatch.Discoveries())
	}

	return strings.Join(list, "")
}
