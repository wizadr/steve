package parsers

type variableDeclarations struct {
	list []VariableDeclaration
}

func createVariableDeclarations(
	list []VariableDeclaration,
) VariableDeclarations {
	out := variableDeclarations{
		list: list,
	}

	return &out
}

// All returns the declarations
func (obj *variableDeclarations) All() []VariableDeclaration {
	return obj.list
}
