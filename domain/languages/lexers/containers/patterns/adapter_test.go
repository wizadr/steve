package patterns

import (
	"testing"
)

func TestAdapter_withMultipleCharactersPattern_withOneRuleInstance_Success(t *testing.T) {
	patternScript := `
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		SLASH_STAR: $SLASH $STAR;
		STAR_SLASH: $STAR $SLASH;

		$SLASH: "/";
		$STAR: "*";
    `

	adapter := NewAdapter()
	patterns, err := adapter.ToPatterns(patternScript)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	comment := `/* this is some content *//* this is some other content */`
	result, err := adapter.FromPatternsToResult(patterns, comment)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if result.Discoveries().Content() != "/* this is some content */" {
		t.Errorf("the content was expected to be '%s', '%s' returned", "/* this is some content */", result.Discoveries().Content())
		return
	}
}

func TestAdapter_fromGroupToResults_withMultiCharacters_Success(t *testing.T) {
	group, err := NewGroupBuilder().Create().WithList([]string{
		"float",
	}).WithName("float").Now()
	if err != nil {
		panic(err)
	}

	content := "float"
	result, err := NewAdapter().FromGroupToResult(group, content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if result.Discoveries().Content() != content {
		t.Errorf("the discoveries were expected to be '%s', '%s' returned", content, result.Discoveries().Content())
		return
	}
}

func TestAdapter_toPatterns_Success(t *testing.T) {
	content := `
        some_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;
		quotation: QUOTATION;

        SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		QUOTATION: $QUOTATION;

        $LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
        $NUMBERS: "0,1,2,3,4,5,6,7,8,9";
        $SLASH: "/";
        $STAR: "*";
		$QUOTATION: "\"";
    `

	adapter := NewAdapter()
	patterns, err := adapter.ToPatterns(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	patternList := patterns.All()
	if len(patternList) != 2 {
		t.Errorf("%d patterns were expected, %d returned", 2, len(patternList))
		return
	}

	pattern, err := patterns.Find("some_comment")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	patternName := pattern.Name()
	if patternName != "some_comment" {
		t.Errorf("the pattern name was expected to be '%s', '%s' returned", "some_comment", patternName)
		return
	}

	patternContentLength := pattern.Content().Length()
	if patternContentLength != 3 {
		t.Errorf("the pattern content length was expected to be %d, %d returned", 3, patternContentLength)
		return
	}

	choices := pattern.Content().List()
	if len(choices) != 3 {
		t.Errorf("%d pattern choices were expected, %d returned", 3, len(choices))
		return
	}
}

func TestAdapter_fromPatternToResult_withComments_withRemaining_Success(t *testing.T) {
	patternScript := `
        comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

        SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;

        $SLASH: "/";
        $STAR: "*";
    `

	adapter := NewAdapter()
	patterns, err := adapter.ToPatterns(patternScript)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	comment := `12/*This is the first comment*/some`

	pattern, err := patterns.Find("comment")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := adapter.FromPatternToResult(pattern, comment)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Discoveries().Index()
	if index != 2 {
		t.Errorf("the index was expectedto be %d, %d returned", 2, index)
		return
	}

	discovery, err := result.Discoveries().Find("STAR_SLASH")
	if err != nil {
		t.Errorf("the index was expectedto be %d, %d returned", 2, index)
		return
	}

	expected := "This is the first comment"
	if discovery.Content() != expected {
		t.Errorf("the content was expected to be '%s', '%s' returned", expected, discovery.Content())
		return
	}

	if !result.HasRemaining() {
		t.Errorf("the result was expecting remaining result, none returned")
		return
	}

	expectedRemaining := "some"
	if result.Remaining() != expectedRemaining {
		t.Errorf("the remaining was expected to be '%s', '%s' returned", expectedRemaining, result.Remaining())
		return
	}

	if result.HasNext() {
		t.Errorf("the result was NOT expecting Next result")
		return
	}
}

func TestAdapter_fromPatternToResult_withComments_withRemaining_withNext_Success(t *testing.T) {
	patternScript := `
        comment: SLASH_STAR @STAR_SLASH STAR_SLASH SOMETHING;

        SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		SOMETHING: $SOMETHING;

        $SLASH: "/";
        $STAR: "*";
		$SOMETHING: "@";
    `

	adapter := NewAdapter()
	patterns, err := adapter.ToPatterns(patternScript)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	comment := `12/*This is the first comment*/some`

	pattern, err := patterns.Find("comment")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := adapter.FromPatternToResult(pattern, comment)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Discoveries().Index()
	if index != 2 {
		t.Errorf("the index was expectedto be %d, %d returned", 2, index)
		return
	}

	discovery, err := result.Discoveries().Find("STAR_SLASH")
	if err != nil {
		t.Errorf("the index was expectedto be %d, %d returned", 2, index)
		return
	}

	expected := "This is the first comment"
	if discovery.Content() != expected {
		t.Errorf("the content was expected to be '%s', '%s' returned", expected, discovery.Content())
		return
	}

	if !result.HasRemaining() {
		t.Errorf("the result was expecting remaining result, none returned")
		return
	}

	expectedRemaining := "some"
	if result.Remaining() != expectedRemaining {
		t.Errorf("the remaining was expected to be '%s', '%s' returned", expectedRemaining, result.Remaining())
		return
	}

	if !result.HasNext() {
		t.Errorf("the result was expecting Next result, none returned")
		return
	}

	next := result.Next()
	expectedGroupName := "$SOMETHING"
	if next.Group().Name() != expectedGroupName {
		t.Errorf("the group was expected to be '%s', '%s' returned", expectedGroupName, next.Group().Name())
		return
	}
}

func TestAdapter_fromPatternToResult_withComments_withoutRemaining_withNext_Success(t *testing.T) {
	patternScript := `
        comment: SLASH_STAR @STAR_SLASH STAR_SLASH SOMETHING;

        SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		SOMETHING: $SOMETHING;

        $SLASH: "/";
        $STAR: "*";
		$SOMETHING: "@";
    `

	adapter := NewAdapter()
	patterns, err := adapter.ToPatterns(patternScript)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	comment := `12/*This is the first comment*/`

	pattern, err := patterns.Find("comment")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := adapter.FromPatternToResult(pattern, comment)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Discoveries().Index()
	if index != 2 {
		t.Errorf("the index was expectedto be %d, %d returned", 2, index)
		return
	}

	discovery, err := result.Discoveries().Find("STAR_SLASH")
	if err != nil {
		t.Errorf("the index was expectedto be %d, %d returned", 2, index)
		return
	}

	expected := "This is the first comment"
	if discovery.Content() != expected {
		t.Errorf("the content was expected to be '%s', '%s' returned", expected, discovery.Content())
		return
	}

	if result.HasRemaining() {
		t.Errorf("the result was NOT expecting remaining result")
		return
	}

	if !result.HasNext() {
		t.Errorf("the result was expecting Next result, none returned")
		return
	}

	next := result.Next()
	expectedGroupName := "$SOMETHING"
	if next.Group().Name() != expectedGroupName {
		t.Errorf("the group was expected to be '%s', '%s' returned", expectedGroupName, next.Group().Name())
		return
	}
}

func TestAdapter_fromPatternToResult_withComments_withoutRemaining_withoutNext_Success(t *testing.T) {
	patternScript := `
        comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

        SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;

        $SLASH: "/";
        $STAR: "*";
    `

	adapter := NewAdapter()
	patterns, err := adapter.ToPatterns(patternScript)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	comment := `12/*This is the first comment*/`

	pattern, err := patterns.Find("comment")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := adapter.FromPatternToResult(pattern, comment)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Discoveries().Index()
	if index != 2 {
		t.Errorf("the index was expectedto be %d, %d returned", 2, index)
		return
	}

	discovery, err := result.Discoveries().Find("STAR_SLASH")
	if err != nil {
		t.Errorf("the index was expectedto be %d, %d returned", 2, index)
		return
	}

	expected := "This is the first comment"
	if discovery.Content() != expected {
		t.Errorf("the content was expected to be '%s', '%s' returned", expected, discovery.Content())
		return
	}

	if result.HasRemaining() {
		t.Errorf("the result was NOT expecting remaining result")
		return
	}

	if result.HasNext() {
		t.Errorf("the result was NOT expecting Next result")
		return
	}
}

func TestAdapter_fromPatternToResult_Success(t *testing.T) {
	patternScript := `
        comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

        SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;

        $SLASH: "/";
        $STAR: "*";
    `

	adapter := NewAdapter()
	patterns, err := adapter.ToPatterns(patternScript)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	comment := `
		sdfsdf
	`

	pattern, err := patterns.Find("comment")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = adapter.FromPatternToResult(pattern, comment)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_fromPatternsToResult_Success(t *testing.T) {
	patternScript := `
        min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;

        MIN_LETTERS: $MIN_LETTERS;
        MAJ_LETTERS: $MAJ_LETTERS;

        $MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
        $MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
    `

	adapter := NewAdapter()
	patterns, err := adapter.ToPatterns(patternScript)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	content := "->JohnDoe<-"
	ins, err := adapter.FromPatternsToResult(patterns, content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := "J"
	if ins.Discoveries().Content() != expected {
		t.Errorf("the content was expected to be '%s','%s' returned", expected, ins.Discoveries().Content())
		return
	}

}
