package patterns

import (
	"errors"
	"fmt"
)

type groups struct {
	list []Group
	mp   map[string]Group
}

func createGroups(
	list []Group,
	mp map[string]Group,
) Groups {
	out := groups{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns the groups
func (obj *groups) All() []Group {
	return obj.list
}

// Find finds a group by name
func (obj *groups) Find(name string) (Group, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the group (name: %s) does not exists", name)
	return nil, errors.New(str)
}
