package results

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

type adapter struct {
	patternAdapter patterns.Adapter
	groupBuilder   patterns.GroupBuilder
	builder        Builder
}

func createAdapter(
	patternAdapter patterns.Adapter,
	groupBuilder patterns.GroupBuilder,
	builder Builder,
) Adapter {
	out := adapter{
		patternAdapter: patternAdapter,
		groupBuilder:   groupBuilder,
		builder:        builder,
	}

	return &out
}

// ToResult converts an input with script to a result instance
func (app *adapter) ToResult(input rules.Rule, script string) (Result, error) {
	content := input.Element().Content()
	builder := app.builder.Create().WithInput(input)
	if content.IsConstant() {
		constant := content.Constant()
		group, err := app.groupBuilder.WithName(constant).WithList([]string{
			constant,
		}).Now()

		if err != nil {
			return nil, err
		}

		ins, _ := app.patternAdapter.FromGroupToResult(group, script)
		if ins != nil {
			builder.WithMatches([]patterns.Result{
				ins,
			})
		}
	}

	if content.IsPattern() {
		pattern := content.Pattern()
		subPatterns := pattern.SubPatterns()
		cardinality := pattern.Cardinality()
		_, pMaxAmount := cardinality.Delimiter()

		list := []patterns.Result{}
		remaining := script

		for {

			if remaining == "" {
				break
			}

			if pMaxAmount != nil && len(list) >= int(*pMaxAmount) {
				break
			}

			result, err := app.patternAdapter.FromPatternsToResult(subPatterns, remaining)
			if err != nil {
				break
			}

			if len(list) >= 0 && result.Discoveries().Index() > 0 {
				break
			}

			list = append(list, result)
			amount := result.Discoveries().Amount()
			remaining = remaining[amount:]
		}

		if len(list) > 0 {
			builder.WithMatches(list)
		}

	}

	return builder.Now()
}
