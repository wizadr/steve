package results

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

func toResultForTests(patternScript string, ruleScript string, script string) Result {
	var ptrns patterns.Patterns
	if patternScript != "" {
		patterns, err := patterns.NewAdapter().ToPatterns(patternScript)
		if err != nil {
			panic(err)
		}

		ptrns = patterns
	}

	rulesAdapter := rules.NewAdapter()
	rules, err := rulesAdapter.Rules(ruleScript, ptrns)
	if err != nil {
		panic(err)
	}

	rulesList := rules.All()
	result, err := NewAdapter().ToResult(rulesList[0], script)
	if err != nil {
		panic(err)
	}

	return result
}
