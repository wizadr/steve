package results

import (
	"testing"
)

func TestAdapter_withComment_withRuleWithOneInstance_Success(t *testing.T) {
	patternsScript := `
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		SLASH_STAR: $SLASH $STAR;
		STAR_SLASH: $STAR $SLASH;

		$SLASH: "/";
		$STAR: "*";
	`

	rulesScript := `
        COMMENT: [multi_line_comment]{1};
    `

	script := "/* this is a comment *//* this is another comment */"
	result := toResultForTests(patternsScript, rulesScript, script)
	if !result.HasMatches() {
		t.Errorf("the result was expected to contain matches")
		return
	}

	if result.Matches().Amount() != 1 {
		t.Errorf("%d results were expected, %d returned", 1, result.Matches().Amount())
		return
	}

	matches := result.Matches().List()
	if matches[0].Discoveries().Content() != "/* this is a comment */" {
		t.Errorf("the comment was expected to be '%s', '%s' returned", "/* this is a comment */", matches[0].Discoveries().Content())
		return
	}
}

func TestAdapter_withComment_withRuleWithTwoInstances_Success(t *testing.T) {
	patternsScript := `
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		SLASH_STAR: $SLASH $STAR;
		STAR_SLASH: $STAR $SLASH;

		$SLASH: "/";
		$STAR: "*";
	`

	rulesScript := `
        COMMENT: [multi_line_comment]{2};
    `

	script := "/* this is a comment *//* this is another comment */"
	result := toResultForTests(patternsScript, rulesScript, script)
	if !result.HasMatches() {
		t.Errorf("the result was expected to contain matches")
		return
	}

	if result.Matches().Amount() != 2 {
		t.Errorf("%d results were expected, %d returned", 2, result.Matches().Amount())
		return
	}

	matches := result.Matches().List()
	if matches[0].Discoveries().Content() != "/* this is a comment */" {
		t.Errorf("the comment was expected to be '%s', '%s' returned", "/* this is a comment */", matches[0].Discoveries().Content())
		return
	}

	if matches[1].Discoveries().Content() != "/* this is another comment */" {
		t.Errorf("the comment was expected to be '%s', '%s' returned", "/* this is another comment */", matches[1].Discoveries().Content())
		return
	}
}

func TestAdapter_witPrefix_noMatch_Success(t *testing.T) {
	patternsScript := `
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		SLASH_STAR: $SLASH $STAR;
		STAR_SLASH: $STAR $SLASH;

		$SLASH: "/";
		$STAR: "*";
	`

	rulesScript := `
        COMMENT: [multi_line_comment]{2};
    `

	script := "prefix /* this is a comment *//* this is another comment */"
	result := toResultForTests(patternsScript, rulesScript, script)
	if result.HasMatches() {
		t.Errorf("the result was NOT expected to contain matches")
		return
	}
}

func TestAdapter_withRuleConstant_Success(t *testing.T) {
	rulesScript := `
        STAR: "*";
    `

	script := "***"
	result := toResultForTests("", rulesScript, script)
	if !result.HasMatches() {
		t.Errorf("the result was expected to contain matches")
		return
	}

	if result.Matches().Amount() != 1 {
		t.Errorf("%d results were expected, %d returned", 1, result.Matches().Amount())
		return
	}

	matches := result.Matches().List()
	if matches[0].Discoveries().Content() != "*" {
		t.Errorf("the comment was expected to be '%s', '%s' returned", "*", matches[0].Discoveries().Content())
		return
	}
}

func TestAdapter_withPatternWithSingleCharacter_Success(t *testing.T) {
	patternsScript := `
		star: STAR;

		STAR: $STAR;

		$STAR: "*";
	`

	rulesScript := `
        STARS: [star]+;
    `

	script := "**********"
	result := toResultForTests(patternsScript, rulesScript, script)
	if !result.HasMatches() {
		t.Errorf("the result was expected to contain matches")
		return
	}

	if result.Matches().Amount() != 10 {
		t.Errorf("%d results were expected, %d returned", 10, result.Matches().Amount())
		return
	}

	matches := result.Matches().List()
	for _, oneMatch := range matches {
		if oneMatch.Discoveries().Content() != "*" {
			t.Errorf("the comment was expected to be '%s', '%s' returned", "*", oneMatch.Discoveries().Content())
			return
		}
	}
}
