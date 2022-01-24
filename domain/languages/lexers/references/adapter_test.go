package references

import (
	"testing"
)

func TestAdapter_withReplacements_Success(t *testing.T) {
	content := `
        include myLocalRef "./my/rules.alex;./my/tokens.alex"
		include notAdded "./my/notadded/rules.alex;./my/notadded/tokens.alex"
        replace myLocalRef.mySubToken localToken
    `

	adapter := NewAdapter()
	references, err := adapter.ToReferences(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if references.HasReferences() {
		t.Errorf("the References instance was NOT expected to contain references")
		return
	}

	if !references.HasReplacements() {
		t.Errorf("the References instance was expected to contain replacements")
		return
	}

	localToken, err := references.Replacements().FetchByLocalToken("localToken")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if localToken.LocalToken() != "localToken" {
		t.Errorf("the localToken was expected to be %s, %s returned", "localToken", localToken.LocalToken())
		return
	}

	if localToken.ReferenceToken() != "mySubToken" {
		t.Errorf("the referenceToken was expected to be %s, %s returned", "mySubToken", localToken.ReferenceToken())
		return
	}

	include := localToken.Include()
	if include.Name() != "myLocalRef" {
		t.Errorf("the name was expected to be %s, %s returned", "myLocalRef", include.Name())
		return
	}

	paths := include.Paths()
	if paths.Rules() != "./my/rules.alex" {
		t.Errorf("the path was expected to be %s, %s returned", "./my/rules.alex", paths.Rules())
		return
	}

	if paths.Tokens() != "./my/tokens.alex" {
		t.Errorf("the path was expected to be %s, %s returned", "./my/tokens.alex", paths.Tokens())
		return
	}
}

func TestAdapter_withReferences_Success(t *testing.T) {
	content := `
        include myLocalRef "./my/rules.alex;./my/tokens.alex"
		include notAdded "./my/notadded/rules.alex;./my/notadded/tokens.alex"
        reference localToken myLocalRef.mySubToken
    `

	adapter := NewAdapter()
	references, err := adapter.ToReferences(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !references.HasReferences() {
		t.Errorf("the References instance was expected to contain references")
		return
	}

	if references.HasReplacements() {
		t.Errorf("the References instance was NOT expected to contain replacements")
		return
	}

	localToken, err := references.References().FetchByLocalToken("localToken")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if localToken.LocalToken() != "localToken" {
		t.Errorf("the localToken was expected to be %s, %s returned", "localToken", localToken.LocalToken())
		return
	}

	if localToken.ReferenceToken() != "mySubToken" {
		t.Errorf("the referenceToken was expected to be %s, %s returned", "mySubToken", localToken.ReferenceToken())
		return
	}

	include := localToken.Include()
	if include.Name() != "myLocalRef" {
		t.Errorf("the name was expected to be %s, %s returned", "myLocalRef", include.Name())
		return
	}

	paths := include.Paths()
	if paths.Rules() != "./my/rules.alex" {
		t.Errorf("the path was expected to be %s, %s returned", "./my/rules.alex", paths.Rules())
		return
	}

	if paths.Tokens() != "./my/tokens.alex" {
		t.Errorf("the path was expected to be %s, %s returned", "./my/tokens.alex", paths.Tokens())
		return
	}
}

func TestAdapter_withReferences_withChannels_Success(t *testing.T) {
	content := `
        include myLocalRef "./my/rules.alex;./my/tokens.alex;./my/channels.alex"
		include notAdded "./my/notadded/rules.alex;./my/notadded/tokens.alex"
        reference localToken myLocalRef.mySubToken
    `

	adapter := NewAdapter()
	references, err := adapter.ToReferences(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !references.HasReferences() {
		t.Errorf("the References instance was expected to contain references")
		return
	}

	if references.HasReplacements() {
		t.Errorf("the References instance was NOT expected to contain replacements")
		return
	}

	localToken, err := references.References().FetchByLocalToken("localToken")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if localToken.LocalToken() != "localToken" {
		t.Errorf("the localToken was expected to be %s, %s returned", "localToken", localToken.LocalToken())
		return
	}

	if localToken.ReferenceToken() != "mySubToken" {
		t.Errorf("the referenceToken was expected to be %s, %s returned", "mySubToken", localToken.ReferenceToken())
		return
	}

	include := localToken.Include()
	if include.Name() != "myLocalRef" {
		t.Errorf("the name was expected to be %s, %s returned", "myLocalRef", include.Name())
		return
	}

	paths := include.Paths()
	if paths.Rules() != "./my/rules.alex" {
		t.Errorf("the path was expected to be %s, %s returned", "./my/rules.alex", paths.Rules())
		return
	}

	if paths.Tokens() != "./my/tokens.alex" {
		t.Errorf("the path was expected to be %s, %s returned", "./my/tokens.alex", paths.Tokens())
		return
	}

	if paths.Channels() != "./my/channels.alex" {
		t.Errorf("the path was expected to be %s, %s returned", "./my/channels.alex", paths.Channels())
		return
	}
}
