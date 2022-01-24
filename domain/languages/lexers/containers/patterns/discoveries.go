package patterns

import (
	"errors"
	"fmt"
	"strings"
)

type discoveries struct {
	list []Discovery
	mp   map[string]Discovery
}

func createDiscoveries(
	list []Discovery,
	mp map[string]Discovery,
) Discoveries {
	out := discoveries{
		list: list,
		mp:   mp,
	}

	return &out
}

// Index returns the index, if any
func (obj *discoveries) Index() uint {
	return obj.list[0].Index()
}

// Content returns the content
func (obj *discoveries) Content() string {
	list := []string{}
	for _, oneContent := range obj.list {
		list = append(list, oneContent.Content())
	}

	return strings.Join(list, "")
}

// Find returns the discovery by container name
func (obj *discoveries) Find(name string) (Discovery, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the container (name: %s) could not be found in the Discoveries", name)
	return nil, errors.New(str)
}

// All returns the discoveries
func (obj *discoveries) All() []Discovery {
	return obj.list
}

// Amount returns the amount of discoveries
func (obj *discoveries) Amount() uint {
	content := obj.Content()
	return uint(len(content))
}

// IsAmount returns true if the amount of discioveries is the same as the passed amount
func (obj *discoveries) IsAmount(amount uint) bool {
	return obj.Amount() == amount
}

// IsInRange returns true if in range, false otherwise
func (obj *discoveries) IsInRange(min uint, max uint) bool {
	currentAmount := len(obj.list)
	return obj.IsAtLeast(min) && int(max) >= currentAmount
}

// IsAtLeast returns true if the current amount is at least, false otherwise
func (obj *discoveries) IsAtLeast(min uint) bool {
	currentAmount := obj.Amount()
	return min <= currentAmount
}
