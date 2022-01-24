package tokens

import (
	"testing"
)

func TestAdapter_withRuleWithConstant_withCardinalityTwo_Success(t *testing.T) {
	rulesScript := `
		SMALLER_THAN: "<";
    `

	tokensScript := `
		twoSmallerThan
	        : SMALLER_THAN{2}
	        ;
    `

	// create the adapter:
	adapter := createAdapterForTests("", rulesScript, tokensScript, "", "twoSmallerThan")

	// fetch the token:
	script := "<<"
	ins, err := adapter.ToToken(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins == nil {
		t.Errorf("the instance was expected to be valid, nil returned")
		return
	}

	if script != ins.Discoveries() {
		t.Errorf("the discoveries were expected to be '%s','%s' returned", script, ins.Discoveries())
		return
	}
}

func TestAdapter_fullName_withSpaceInChannels_Success(t *testing.T) {

	patternsScript := `
		whitespace: WHITESPACE;
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;
		single_line_comment: TWO_SLASHES @EOL EOL;
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		WHITESPACE: $WHITESPACE;
		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;
		SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		TWO_SLASHES: $TWO_SLASHES;
		EOL: $EOL;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
		$WHITESPACE: " ,	,
";
		$EOL: "
";
		$SLASH: "/";
		$TWO_SLASHES: "//";
        $STAR: "*";
    `

	rulesScript := `
		WHITESPACES: [whitespace]+;
		SINGLE_LINE_COMMENT: [single_line_comment]{1};
		MULTI_LINE_COMMENT: [multi_line_comment]{1};
		NAME_FIRST_LETTER: [maj_letters]{1};
		NAME_AFTER_FIRST_LETTER: [min_letters]+;
		SMALLER_THAN: "<";
		BIGGER_THAN: ">";
		DOUBLE_SLASHES: "//";
		SLASH_STAR: "/*";
		STAR_SLASH: "*/";
		SPACE: " ";
		COMMA: ",";
    `

	tokensScript := `
		fullName
	        : name{2}
	        ;

		name
			: # NAME_FIRST_LETTER NAME_AFTER_FIRST_LETTER #
			;
    `

	channelScript := `
		whiteSpace
			: WHITESPACES
			;

		multiLineComment
			: MULTI_LINE_COMMENT
			;

		lineComment
			: SINGLE_LINE_COMMENT
			;
    `

	// create the adapter:
	adapter := createAdapterForTests(patternsScript, rulesScript, tokensScript, channelScript, "fullName")

	// fetch the token:
	script := `
		Roger Cyr
	`

	ins, err := adapter.ToToken(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins == nil {
		t.Errorf("the instance was expected to be valid, nil returned")
		return
	}

	expected := `
		Roger Cyr`
	if expected != ins.Discoveries() {
		t.Errorf("the discoveries were expected to be '%s','%s' returned", expected, ins.Discoveries())
		return
	}
}

func TestAdapter_containsMultipleValidChoices_firstLineIsLiongest_takesFirstLine_Success(t *testing.T) {
	patternsScript := `
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;
		single_line_comment: TWO_SLASHES @EOF EOF;
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;
		SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		TWO_SLASHES: $TWO_SLASHES;
		EOF: $EOF;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
		$EOF: "\r\n,\n";
		$SLASH: "/";
		$TWO_SLASHES: "//";
        $STAR: "*";
    `

	rulesScript := `
		SINGLE_LINE_COMMENT: [single_line_comment]{1};
		MULTI_LINE_COMMENT: [multi_line_comment]{1};
		NAME_FIRST_LETTER: [maj_letters]{1};
		NAME_AFTER_FIRST_LETTER: [min_letters]+;
		SMALLER_THAN: "<";
		BIGGER_THAN: ">";
		DOUBLE_SLASHES: "//";
		SLASH_STAR: "/*";
		STAR_SLASH: "*/";
		SPACE: " ";
		COMMA: ",";
    `

	tokensScript := `
		multipleNames
			: fullName COMMA fullName
			| fullName
			;

		fullName
	        : name{1} SPACE name{1}
	        ;

		name
			: # NAME_FIRST_LETTER NAME_AFTER_FIRST_LETTER #
			;
    `

	channelScript := `
		multiLineComment
			: MULTI_LINE_COMMENT
			;

		lineComment
			: SINGLE_LINE_COMMENT
			;
    `

	// create the adapter:
	adapter := createAdapterForTests(patternsScript, rulesScript, tokensScript, channelScript, "multipleNames")

	// fetch the token:
	script := "Roger Cyr,Paul Dube"
	ins, err := adapter.ToToken(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins == nil {
		t.Errorf("the instance was expected to be valid, nil returned")
		return
	}

	if script != ins.Discoveries() {
		t.Errorf("the discoveries were expected to be the whole script")
		return
	}
}

func TestAdapter_containsMultipleValidChoices_secondLineIsLiongest_takesSecondLine_Success(t *testing.T) {
	patternsScript := `
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;
		single_line_comment: TWO_SLASHES @EOF EOF;
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;
		SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		TWO_SLASHES: $TWO_SLASHES;
		EOF: $EOF;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
		$EOF: "\r\n,\n";
		$SLASH: "/";
		$TWO_SLASHES: "//";
        $STAR: "*";
    `

	rulesScript := `
		SINGLE_LINE_COMMENT: [single_line_comment]{1};
		MULTI_LINE_COMMENT: [multi_line_comment]{1};
		NAME_FIRST_LETTER: [maj_letters]{1};
		NAME_AFTER_FIRST_LETTER: [min_letters]+;
		SMALLER_THAN: "<";
		BIGGER_THAN: ">";
		DOUBLE_SLASHES: "//";
		SLASH_STAR: "/*";
		STAR_SLASH: "*/";
		SPACE: " ";
		COMMA: ",";
    `

	tokensScript := `
		multipleNames
			: fullName
			| fullName COMMA fullName
			;

		fullName
	        : name{1} SPACE name{1}
	        ;

		name
			: # NAME_FIRST_LETTER NAME_AFTER_FIRST_LETTER #
			;
    `

	channelScript := `
		multiLineComment
			: SINGLE_LINE_COMMENT
			;

		lineComment
			: MULTI_LINE_COMMENT
			;
    `

	// create the adapter:
	adapter := createAdapterForTests(patternsScript, rulesScript, tokensScript, channelScript, "multipleNames")

	// fetch the token:
	script := "Roger Cyr,John Addams"
	ins, err := adapter.ToToken(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins == nil {
		t.Errorf("the instance was expected to be valid, nil returned")
		return
	}

	if script != ins.Discoveries() {
		t.Errorf("the discoveries were expected to be the whole script")
		return
	}
}

func TestAdapter_twoTokens_Success(t *testing.T) {
	patternsScript := `
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;
		single_line_comment: TWO_SLASHES @EOF EOF;
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;
		SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		TWO_SLASHES: $TWO_SLASHES;
		EOF: $EOF;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
		$EOF: "\r\n,\n";
		$SLASH: "/";
		$TWO_SLASHES: "//";
        $STAR: "*";
    `

	rulesScript := `
		SINGLE_LINE_COMMENT: [single_line_comment]{1};
		MULTI_LINE_COMMENT: [multi_line_comment]{1};
		NAME_FIRST_LETTER: [maj_letters]{1};
		NAME_AFTER_FIRST_LETTER: [min_letters]+;
		COMMERCIAL_A: "@";
		DOT: ".";
		SMALLER_THAN: "<";
		BIGGER_THAN: ">";
		DOUBLE_SLASHES: "//";
		SLASH_STAR: "/*";
		STAR_SLASH: "*/";
		SPACE: " ";
    `

	tokensScript := `
		emailWithName
			: fullName SPACE emailWithDelimiters
			;

		emailWithDelimiters
			: SMALLER_THAN email BIGGER_THAN
			;

		fullName
	        : name{1} SPACE name{1}
	        ;

		email
			: NAME_AFTER_FIRST_LETTER COMMERCIAL_A NAME_AFTER_FIRST_LETTER DOT NAME_AFTER_FIRST_LETTER
			;

		name
			: # NAME_FIRST_LETTER NAME_AFTER_FIRST_LETTER #
			;
    `

	channelScript := `
		multiLineComment
			: SINGLE_LINE_COMMENT
			;

		lineComment
			: MULTI_LINE_COMMENT
			;
    `

	// create the adapter:
	adapter := createAdapterForTests(patternsScript, rulesScript, tokensScript, channelScript, "emailWithName")

	// fetch the token:
	script := "Roger Cyr <roger@cyr.ca>"
	ins, err := adapter.ToToken(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins == nil {
		t.Errorf("the instance was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_singleToken_Success(t *testing.T) {
	patternsScript := `
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;
		single_line_comment: TWO_SLASHES @EOF EOF;
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;
		SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		TWO_SLASHES: $TWO_SLASHES;
		EOF: $EOF;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
		$EOF: "\r\n,\n";
		$SLASH: "/";
		$TWO_SLASHES: "//";
        $STAR: "*";
    `

	rulesScript := `
		SINGLE_LINE_COMMENT: [single_line_comment]{1};
		MULTI_LINE_COMMENT: [multi_line_comment]{1};
		NAME_FIRST_LETTER: [maj_letters]{1};
        NAME_AFTER_FIRST_LETTER: [min_letters]+;
		DOUBLE_SLASHES: "//";
		SLASH_STAR: "/*";
		STAR_SLASH: "*/";
		SPACE: " ";
    `

	tokensScript := `
        fullName
            : name{1} SPACE name{1}
            ;

		name
			: NAME_FIRST_LETTER NAME_AFTER_FIRST_LETTER
			;
    `

	channelScript := `
		multiLineComment
			: SINGLE_LINE_COMMENT
			;

		lineComment
			: MULTI_LINE_COMMENT
			;

    `

	// create the adapter:
	adapter := createAdapterForTests(patternsScript, rulesScript, tokensScript, channelScript, "fullName")

	// fetch the token:
	script := "Roger Cyr"
	ins, err := adapter.ToToken(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins == nil {
		t.Errorf("the instance was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_containsRecursivity_Success(t *testing.T) {
	patternsScript := `
		numbers: NUMBERS;
		whitespace: WHITESPACES;
		single_line_comment: TWO_SLASHES @EOF EOF;
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		NUMBERS: $NUMBERS;
		SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		TWO_SLASHES: $TWO_SLASHES;
		WHITESPACES: $WHITESPACES;
		EOF: $EOF;

		$NUMBERS: "0,1,2,3,4,5,6,7,8,9";
		$WHITESPACES: " ,\t,\r,\n";
		$EOF: "\r\n,\n";
		$SLASH: "/";
		$TWO_SLASHES: "//";
        $STAR: "*";
    `

	rulesScript := `
		INT_PATTERN: [numbers]+;
		WHITESPACE: [whitespace]{1};
		SINGLE_LINE_COMMENT: [single_line_comment]{1};
		MULTI_LINE_COMMENT: [multi_line_comment]{1};
		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";
		PLUS: "+";
		DOUBLE_SLASHES: "//";
		SLASH_STAR: "/*";
		STAR_SLASH: "*/";
    `

	tokensScript := `
		forTest
			: forElementTest
			| OPEN_PARENTHESIS forTest CLOSE_PARENTHESIS
			+++
			:
				+45
				###
				(+45)
				###
				((((+45))))
				###
			;

		forElementTest
			: PLUS INT_PATTERN
			+++
			:
				+45
				###
			;
    `

	channelScript := `
		whiteSpace
			: WHITESPACE
			;

		multiLineComment
			: SINGLE_LINE_COMMENT
			;

		lineComment
			: MULTI_LINE_COMMENT
			;

    `

	// create the adapter:
	adapter := createAdapterForTests(patternsScript, rulesScript, tokensScript, channelScript, "forTest")

	// fetch the token:
	script := `((+45))`
	ins, err := adapter.ToToken(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins == nil {
		t.Errorf("the instance was expected to be valid, nil returned")
		return
	}

	if script != ins.Discoveries() {
		t.Errorf("the discoveries were expected to be the whole script")
		return
	}

}
