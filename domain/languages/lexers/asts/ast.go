package asts

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
)

type ast struct {
	mistake mistakes.Mistake
	success Token
}

func createASTWithMistake(
	mistake mistakes.Mistake,
) AST {
	return createASTInternally(mistake, nil)
}

func createASTWithSuccess(
	success Token,
) AST {
	return createASTInternally(nil, success)
}

func createASTInternally(
	mistake mistakes.Mistake,
	success Token,
) AST {
	out := ast{
		mistake: mistake,
		success: success,
	}

	return &out
}

// IsMistake returns true if there is a mistake, false otherwise
func (obj *ast) IsMistake() bool {
	return obj.mistake != nil
}

// Mistake returns the mistake, if any
func (obj *ast) Mistake() mistakes.Mistake {
	return obj.mistake
}

// IsSuccess returns true if the AST is successful, false otherwise
func (obj *ast) IsSuccess() bool {
	return obj.success != nil
}

// Success returns teh successful Token, if any
func (obj *ast) Success() Token {
	return obj.success
}
