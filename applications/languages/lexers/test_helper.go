package lexers

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/steve-care-software/steve/domain/languages/lexers/contents"
	"github.com/steve-care-software/steve/domain/languages/lexers/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"
	"github.com/steve-care-software/steve/domain/languages/lexers/resources/multiples"
	"github.com/steve-care-software/steve/domain/languages/lexers/resources/singles"
	"github.com/steve-care-software/steve/domain/languages/lexers/suites"
)

// NewContentForTests creates a new content for tests
func NewContentForTests(content string) contents.Content {
	ins, err := contents.NewBuilder().Create().WithContent(content).IsPrefixLegal().IsSuffixLegal().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSingleResourceForTests creates a new single resource for tests
func NewSingleResourceForTests(rootToken string, patternFile string, ruleFile string, tokenFile string, channelFile *string) singles.Single {
	multiple := NewMultipleResourceForTests(patternFile, ruleFile, tokenFile, channelFile)
	ins, err := singles.NewBuilder().Create().WithRootToken(rootToken).WithMultiple(multiple).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewMultipleResourceForTests creates a new multiple resource for tests
func NewMultipleResourceForTests(patternFile string, ruleFile string, tokenFile string, channelFile *string) multiples.Multiple {
	container := NewContainerResourceForTests(patternFile, ruleFile, tokenFile)
	multipleBuilder := multiples.NewBuilder().Create().WithContainer(container)
	if channelFile != nil {
		channels, err := ioutil.ReadFile(*channelFile)
		if err != nil {
			panic(err)
		}

		multipleBuilder.WithChannels(string(channels))
	}

	ins, err := multipleBuilder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContainerResourceForTests creates a new container resource for tests
func NewContainerResourceForTests(patternFile string, ruleFile string, tokenFile string) containers.Container {
	patterns, err := ioutil.ReadFile(patternFile)
	if err != nil {
		panic(err)
	}

	rules, err := ioutil.ReadFile(ruleFile)
	if err != nil {
		panic(err)
	}

	tokens, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		panic(err)
	}

	ins, err := containers.NewBuilder().WithPatterns(string(patterns)).WithRules(string(rules)).WithTokens(string(tokens)).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLexerForTests creates a new lexer for tests
func NewLexerForTests(ruleDir string, tokenDir string, channelDir string) Application {
	path, err := paths.NewBuilder().WithRules(ruleDir).WithTokens(tokenDir).WithChannels(channelDir).Now()
	if err != nil {
		panic(err)
	}

	lexer, err := NewBuilder().Create().WithBasePaths(path).Now()
	if err != nil {
		panic(err)
	}

	return lexer
}

// ValidateTestSuitesForTests validates test suites for tests
func ValidateTestSuitesForTests(suites []suites.Suite, t *testing.T) {
	for _, oneSuite := range suites {
		name := oneSuite.Name()
		lines := oneSuite.Lines()
		for _, oneLine := range lines {
			index := oneLine.Index()
			if !oneLine.IsSuccess() {
				t.Errorf("the test (token: %s. index: %d) is invalid", name, index)
				continue
			}

			fmt.Printf("valid: %s:%d\n", name, index)
		}

		fmt.Printf("-----\n")
	}
}
