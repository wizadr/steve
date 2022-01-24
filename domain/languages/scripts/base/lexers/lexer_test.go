package lexers

import (
	"testing"

	"github.com/steve-care-software/steve/applications/languages/lexers"
	"github.com/steve-care-software/steve/domain/languages/scripts"
)

func TestParser_computableType_isFloat_Success(t *testing.T) {
	root := "operationElement"
	script := `
		true
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	if instance == nil {
		t.Errorf("the instance was NOT expected to be nil")
		return
	}

	expected := `
		true`

	code := instance.Content()
	if expected != code {
		t.Errorf("the script (\n\n\n\n\n\n'%s'\n\n\n\n\n\n) was expected to be the same as the instance code (\n\n\n\n\n\n'%s'\n\n\n\n\n\n)", expected, code)
		return
	}
}

func TestParser_Success(t *testing.T) {
	baseDir := "./../../resources"
	langDir := "base"
	lexer := lexers.NewLexerForTests(baseDir, baseDir, baseDir)
	resource := scripts.NewMultipleResourceForTests(baseDir, langDir)
	suites, err := lexer.Tests(resource)
	if err != nil {
		panic(err)
	}

	// validate suites:
	lexers.ValidateTestSuitesForTests(suites, t)
}
