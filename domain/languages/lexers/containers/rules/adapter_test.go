package rules

import (
	"testing"

	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
)

func TestAdapter_Success(t *testing.T) {

	patternsContent := `
		numbers: NUMBERS;
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;
		single_line_comment: DOUBLE_SLASH @NEWLINE NEWLINE
		multi_line_comment: STAR_SLASH @SLASH_STAR SLASH_STAR

		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;
		NUMBERS: $NUMBERS;
		NEWLINE: $NEWLINE;
		STAR_SLASH: $STAR $SLASH;
		SLASH_STAR: $SLASH $STAR;
		DOUBLE_SLASH: $SLASH $SLASH;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
		$NUMBERS: "0,1,2,3,4,5,6,7,8,9";
		$NEWLINE: " \r\n";
		$STAR: "*";
		$SLASH: "/";
	 `

	patterns, err := patterns.NewAdapter().ToPatterns(patternsContent)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	content := `
		NUMBER: [numbers]{1};
		NAME: [min_letters maj_letters]+;
		TWO_THREE_NUMBERS: [numbers]{2,3};
		FOURTH: "y\"ay";
		ONE_LETTER_MAJ: [min_letters]{1};
		ARROW: "-->";
		COMMERCIAL_A: "@";
    `

	adapter := NewAdapter()
	rules, err := adapter.Rules(content, patterns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := rules.All()
	if len(list) != 7 {
		t.Errorf("%d rules were expected, %d returned", 7, len(list))
		return
	}

	if list[3].Element().Content().Constant() != "y\"ay" {
		t.Errorf("the code was expected to be '%s','%s' returned", "y\"ay", list[3].Element().Content().Constant())
		return
	}
}
