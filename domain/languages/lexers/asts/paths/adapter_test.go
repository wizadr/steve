package paths

import (
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	patternsScript := `
		min_letters: MIN_LETTERS;
		maj_letters: MAJ_LETTERS;

		MIN_LETTERS: $MIN_LETTERS;
		MAJ_LETTERS: $MAJ_LETTERS;

		$MIN_LETTERS: "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
		$MAJ_LETTERS: "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z";
    `

	rulesScript := `
        NAME: [min_letters maj_letters]+;
		MIN_LETTERS: [min_letters]+;
		TWO_MIN_LETTERS: [min_letters]{2};
		COMMERCIAL_A: "@";
        SMALLER_THAN: "<";
        BIGGER_THAN: ">";
        COMMA: ",";
		DOT: ".";
		QUOTATION: """;
		VOILA: "Voila";
		ROGER: "Roger";
		PLUS: "+";
		MINUS: "-";
    `

	tokensScript := `
		betweenFiveAndTenEmailsButNotNine
			: nameWithEmails -> nameWithEmail{5,10} <- plusAndMinus -> PLUS{2,} MINUS{3,5} <-
			---
			: nameWithEmails -> nameWithEmail{9} <-
			;

		plusAndMinus
			: PLUS+ MINUS{2,}
			;

        nameWithEmails
            : nameWithEmail nameWithEmailWithComma* VOILA?
			| QUOTATION nameWithEmailsInside -> nameWithEmail{1} <- QUOTATION
            ;

		nameWithEmailsInside
			: nameWithEmail nameWithEmailWithComma* VOILA?
			;

        nameWithEmailWithComma
            : COMMA nameWithEmail
            ;

        nameWithEmail
            : fullName SMALLER_THAN email BIGGER_THAN
            ;

		email
			: NAME COMMERCIAL_A MIN_LETTERS DOT TWO_MIN_LETTERS
			;

		fullName
	        : NAME{2}
			---
			: ROGER NAME
	        ;
    `

	path := toPathForTests(patternsScript, rulesScript, tokensScript, "betweenFiveAndTenEmailsButNotNine")
	if path == nil {
		t.Errorf("the path was expected to be vaid")
	}
}
