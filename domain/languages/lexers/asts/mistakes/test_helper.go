package mistakes

import (
	ast_tokens "github.com/steve-care-software/steve/domain/languages/lexers/asts/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
)

func createTokenForTests(rootTokenName string, script string) ast_tokens.Token {
	patternsScript := `
		numbers: NUMBERS;
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;
		whitespace: WHITESPACE;
		single_line_comment: TWO_SLASHES @EOF EOF;
		multi_line_comment: SLASH_STAR @STAR_SLASH STAR_SLASH;

		NUMBERS: $NUMBERS;
		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;
		SLASH_STAR: $SLASH $STAR;
        STAR_SLASH: $STAR $SLASH;
		TWO_SLASHES: $TWO_SLASHES;
		WHITESPACE: $WHITESPACE;
		EOF: $EOF;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
		$NUMBERS: "0,1,2,3,4,5,6,7,8,9";
		$WHITESPACE: " ,	,
";
		$EOF: "
";
		$SLASH: "/";
		$TWO_SLASHES: "//";
        $STAR: "*";
    `

	rulesScript := `
		SINGLE_LINE_COMMENT: [single_line_comment]{1};
		MULTI_LINE_COMMENT: [multi_line_comment]{1};
		WHITESPACES: [whitespace]+;
		FIRST_LETTER: [maj_letters]{1};
		MIN_LETTERS:  [min_letters]+;
		TWO_MIN_LETTERS_OR_MORE:  [min_letters]{2,};
		INT_PATTERN: [numbers]+;
		DOUBLE_DASH: "//";
		SMALLER_THAN: "<";
		BIGGER_THAN: ">";
		COMMA: ",";
		QUOTATION: """;
		MINUS: "-";
		COMMERCIAL_A: "@";
		DOT: ".";
    `

	tokensScript := `
		intValue
			: MINUS? INT_PATTERN
			;

		twoEmailsWithName
			: betweenTwoAndFiveEmailsWithName -> emailWithDelimiters{2} <-
			;

		betweenTwoAndFiveEmailsWithNameButNotThree
			: fullName emailWithDelimiters{2,5}
			---
			: fullName emailWithDelimiters{3}
			;

		betweenTwoAndFiveEmailsWithName
			: fullName emailWithDelimiters{2,5}
			;

		threeEmailsWithName
			: emailWithName{3}
			;

		threeSmallerThans
			: SMALLER_THAN{3}
			;

		emailWithName
			: fullName emailWithDelimiters
			;

		emailWithDelimiters
			: SMALLER_THAN email BIGGER_THAN
			;

		fullName
			: name{2}
			;

		name
			: # FIRST_LETTER{1} MIN_LETTERS{1} #
			;

		email
			: MIN_LETTERS COMMERCIAL_A MIN_LETTERS DOT TWO_MIN_LETTERS_OR_MORE
			;
    `

	channelScript := `
		whiteSpaces
			: WHITESPACES
			;

		multiLineComment
			: MULTI_LINE_COMMENT
			;

		lineComment
			: SINGLE_LINE_COMMENT
			;
    `

	patterns, err := patterns.NewAdapter().ToPatterns(patternsScript)
	if err != nil {
		panic(err)
	}

	rules, err := rules.NewAdapter().Rules(rulesScript, patterns)
	if err != nil {
		panic(err)
	}

	tokenAdapter, err := tokens.NewAdapterBuilder().Create().WithRules(rules).Now()
	if err != nil {
		panic(err)
	}

	tokens, err := tokenAdapter.ToTokens(tokensScript)
	if err != nil {
		panic(err)
	}

	channels, err := tokenAdapter.ToTokens(channelScript)
	if err != nil {
		panic(err)
	}

	rootToken, err := tokens.Find(rootTokenName)
	if err != nil {
		panic(err)
	}

	astTokenAdapter, err := ast_tokens.NewAdapterBuilder().Create().WithTokens(tokens).WithToken(rootToken).WithChannels(channels).CanContainPrefix().Now()
	if err != nil {
		panic(err)
	}

	astToken, err := astTokenAdapter.ToToken(script)
	if err != nil {
		panic(err)
	}

	return astToken
}
