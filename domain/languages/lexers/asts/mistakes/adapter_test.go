package mistakes

import (
	"testing"
)

func TestAdapter_containsValidNotMatch_returnsMistake(t *testing.T) {

	rootTokenName := "betweenTwoAndFiveEmailsWithNameButNotThree"
	script := `
		Roger Cyr <roger@cyr.ca> <first@last.com> <third@name.com>
	`

	token := createTokenForTests(rootTokenName, script)
	mistake, err := NewAdapter().ToMistake(token, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !mistake.Content().IsContainsValidNotMatch() {
		t.Errorf("the error was expected to be 'contains valid not match', returned: %v", mistake.Content())
		return
	}
}

func TestAdapter_tokenHasNoMatch_returnsMistake(t *testing.T) {

	rootTokenName := "twoEmailsWithName"
	// no match - twoEmailsWithName
	script := `
		this is invalid
	`

	token := createTokenForTests(rootTokenName, script)
	mistake, err := NewAdapter().ToMistake(token, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !mistake.Content().IsTokenHasNoMatch() {
		t.Errorf("the error was expected to be 'token has no match', returned: %v", mistake.Content())
		return
	}
}

func TestAdapter_specifierDoNotMatch_returnsMistake(t *testing.T) {

	rootTokenName := "twoEmailsWithName"
	script := `
		Roger Cyr <roger@cyr.ca> <second@google.ca> <third@google.ca>
	`

	token := createTokenForTests(rootTokenName, script)
	mistake, err := NewAdapter().ToMistake(token, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !mistake.Content().IsSpecifierDoNotMatch() {
		t.Errorf("the error was expected to be 'specifier do not match', returned: %v", mistake.Content())
		return
	}
}

func TestAdapter_containsNextElement_returnsMistake(t *testing.T) {

	rootTokenName := "emailWithName"
	script := `
		Roger Cyr
	`

	token := createTokenForTests(rootTokenName, script)
	mistake, err := NewAdapter().ToMistake(token, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !mistake.Content().IsContainsNextElement() {
		t.Errorf("the error was expected to be 'contains next element', returned: %v", mistake.Content())
		return
	}
}

func TestAdapter_tokenCardinalityIsInvalid_returnsMistake(t *testing.T) {

	rootTokenName := "threeEmailsWithName"
	script := `
		Roger Cyr <roger@cyr.ca>
		First Last <first@last.google.ca>
	`

	token := createTokenForTests(rootTokenName, script)
	mistake, err := NewAdapter().ToMistake(token, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !mistake.Content().IsCardinalityIsInvalid() {
		t.Errorf("the error was expected to be 'cardinality is invalid', returned: %v", mistake.Content())
		return
	}
}

func TestAdapter_withoutMistake_returnsNil(t *testing.T) {

	rootTokenName := "emailWithName"
	script := `
		Roger Cyr <roger@cyr.ca>
	`

	token := createTokenForTests(rootTokenName, script)
	mistake, err := NewAdapter().ToMistake(token, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}
}

func TestAdapter_ruleCardinalityIsInvalid_returnsMistake(t *testing.T) {

	rootTokenName := "threeSmallerThans"
	script := `
		<<
	`

	token := createTokenForTests(rootTokenName, script)
	mistake, err := NewAdapter().ToMistake(token, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !mistake.Content().IsCardinalityIsInvalid() {
		t.Errorf("the error was expected to be 'cardinality is invalid', returned: %v", mistake.Content())
		return
	}
}
