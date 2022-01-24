package patterns

import "strings"

type groupContent struct {
	list []string
}

func createGroupContent(
	list []string,
) GroupContent {
	out := groupContent{
		list: list,
	}

	return &out
}

// List returns the group list
func (obj *groupContent) List() []string {
	return obj.list
}

// First returns the first element of the group
func (obj *groupContent) First() string {
	return obj.list[0]
}

// Length returns the amount of elements in the group
func (obj *groupContent) Length() uint {
	return uint(len(obj.list))
}

// Contains returns the amount of match the group have against the passed content:
func (obj *groupContent) Contains(content string) uint {
	longest := 0
	for _, oneElement := range obj.list {
		if strings.HasPrefix(content, oneElement) {
			prefixLength := len(oneElement)
			if prefixLength >= longest {
				longest = prefixLength
			}
		}
	}

	return uint(longest)
}
