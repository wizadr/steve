package tokens

import (
	"testing"

	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
)

func TestAdapter_Success(t *testing.T) {
	patternsContent := `
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;
		underscore: UNDERSCORE

		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;
		UNDERSCORE: $UNDERSCORE;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
		$UNDERSCORE: "_";
	 `

	patterns, err := patterns.NewAdapter().ToPatterns(patternsContent)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	rulesContent := `
        NAME: [min_letters maj_letters]+;
        DOMAIN: [min_letters maj_letters underscore]{2,};
        EXTENSION: [min_letters]{3,};
        COMMERCIAL_A: "@";
		DOT: ".";
        QUOTATION_MARK: "\"";
        COMMA: ",";
        PIPE: "|";
    `

	rulesAdapter := rules.NewAdapter()
	rulesIns, err := rulesAdapter.Rules(rulesContent, patterns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter, err := NewAdapterBuilder().Create().WithRules(rulesIns).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil")
		return
	}

	tokensContent := `
        sections
            : section+
            ;

        section
            : anyAmountExceptTenOrFifteen PIPE?
            ;

        anyAmountExceptTenOrFifteen
            : emailList
            ---
            : emailList -> email{10} <-
            | emailList -> email{15} <-
            ;

        atMostTenEmails
            : emailList -> email{0,10} <-
            ;

        atLeastThreeEmails
            : emailList -> email{3,} <-
            ;

        threeToNineEmails
            : emailList -> email{3,9} <-
            ;

        fiveEmailsAndForCommas
            : emailList -> email{5} COMMA{4} <-
            ;

        emailList
            : email emailWithCommaPrefix*
            ;

        emailWithCommaPrefix
            : COMMA email
			:
					,myName@myDomain.com
				  	###
				  	success
				###
            ;

        email
            : NAME COMMERCIAL_A DOMAIN DOT EXTENSION
            | QUOTATION_MARK email QUOTATION_MARK
			+++
			:
					myName@myDomain.com
				###
					"myName@myDomain.com"
				###
					this is invalid
				###
					-54
				###
			;
    `

	tokens, err := adapter.ToTokens(tokensContent)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	all := tokens.All()
	if len(all) != 10 {
		t.Errorf("%d tokens were expected, %d returned", 10, len(all))
		return
	}

}
