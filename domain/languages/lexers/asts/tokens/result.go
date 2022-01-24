package tokens

import (
	"strings"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts/results"
)

type result struct {
	list []results.Result
}

func createResult() Result {
	return createResultInternally(nil)
}

func createResultWithList(
	list []results.Result,
) Result {
	return createResultInternally(list)
}

func createResultInternally(
	list []results.Result,
) Result {
	out := result{
		list: list,
	}

	return &out
}

// Amount returns the amount
func (obj *result) Amount() uint {
	if !obj.HasResults() {
		return 0
	}

	return uint(len(obj.list))
}

// Discoveries returns the discoveries
func (obj *result) Discoveries() string {
	if !obj.HasResults() {
		return ""
	}

	discoveries := []string{}
	for _, oneResult := range obj.list {
		discoveries = append(discoveries, oneResult.Content())
	}

	return strings.Join(discoveries, "")
}

// HasResults returns true if there is results, false otherwise
func (obj *result) HasResults() bool {
	return obj.list != nil
}

// Results returns the results, if any
func (obj *result) Results() []results.Result {
	return obj.list
}
