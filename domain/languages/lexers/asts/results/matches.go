package results

import (
	"strings"

	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

type matches struct {
	list []patterns.Result
}

func createMatches(
	list []patterns.Result,
) Matches {
	out := matches{
		list: list,
	}

	return &out
}

// Content returns the content
func (obj *matches) Content() string {
	discoveries := []string{}
	for _, oneContent := range obj.list {
		discoveries = append(discoveries, oneContent.Discoveries().Content())
	}

	return strings.Join(discoveries, "")
}

// Amount returns the amount
func (obj *matches) Amount() uint {
	return uint(len(obj.list))
}

// List returns the list of results
func (obj *matches) List() []patterns.Result {
	return obj.list
}
