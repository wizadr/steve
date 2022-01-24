package parsers

type forDeclaration struct {
	statement    ForStatement
	instructions Instructions
}

func createForDeclaration(
	statement ForStatement,
) ForDeclaration {
	return createForDeclarationInternally(statement, nil)
}

func createForDeclarationWithInstructions(
	statement ForStatement,
	instructions Instructions,
) ForDeclaration {
	return createForDeclarationInternally(statement, instructions)
}

func createForDeclarationInternally(
	statement ForStatement,
	instructions Instructions,
) ForDeclaration {
	out := forDeclaration{
		statement:    statement,
		instructions: instructions,
	}

	return &out
}

// Statement returns the statement
func (obj *forDeclaration) Statement() ForStatement {
	return obj.statement
}

// HasInstructions returns true if there is instructions, false otherwise
func (obj *forDeclaration) HasInstructions() bool {
	return obj.instructions != nil
}

// Instructions returns the instructions, if any
func (obj *forDeclaration) Instructions() Instructions {
	return obj.instructions
}
