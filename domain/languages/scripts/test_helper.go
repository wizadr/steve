package scripts

import (
	"path/filepath"

	applications "github.com/steve-care-software/steve/applications/languages/lexers"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
	"github.com/steve-care-software/steve/domain/languages/lexers/resources/multiples"
	"github.com/steve-care-software/steve/domain/languages/lexers/resources/singles"
	"github.com/steve-care-software/steve/domain/languages/lexers/roots"
)

// ExecuteLexerForTests executes the lexer for tests
func ExecuteLexerForTests(rootToken string, script string, resourcesDir string, langDir string) (roots.Root, mistakes.Mistake) {
	lexer := applications.NewLexerForTests(resourcesDir, resourcesDir, resourcesDir)
	resource := NewSingleResourceForTests(rootToken, resourcesDir, langDir)
	content := applications.NewContentForTests(script)
	root, mistake, err := lexer.Root(resource, content)
	if err != nil {
		panic(err)
	}

	return root, mistake
}

// NewSingleResourceForTests creates a new single resource for tests
func NewSingleResourceForTests(rootToken string, resourcesDir string, langDir string) singles.Single {
	patternFile, ruleFile, tokenFile, channelFile := createPathsForTests(resourcesDir, langDir)
	return applications.NewSingleResourceForTests(rootToken, patternFile, ruleFile, tokenFile, &channelFile)
}

// NewMultipleResourceForTests creates a new multiple resource for tests
func NewMultipleResourceForTests(resourcesDir string, langDir string) multiples.Multiple {
	patternFile, ruleFile, tokenFile, channelFile := createPathsForTests(resourcesDir, langDir)
	return applications.NewMultipleResourceForTests(patternFile, ruleFile, tokenFile, &channelFile)
}

func createPathsForTests(resourcesDir string, langDir string) (string, string, string, string) {
	patternsFile := filepath.Join(resourcesDir, langDir, "patterns.alex")
	ruleFile := filepath.Join(resourcesDir, langDir, "rules.alex")
	tokenFile := filepath.Join(resourcesDir, langDir, "tokens.alex")
	channelFile := filepath.Join(resourcesDir, langDir, "channels.alex")
	return patternsFile, ruleFile, tokenFile, channelFile
}
