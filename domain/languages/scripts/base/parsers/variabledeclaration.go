package parsers

type variableDeclaration struct {
	typ  Type
	name string
}

func createVariableDeclaration(
	typ Type,
	name string,
) VariableDeclaration {
	out := variableDeclaration{
		typ:  typ,
		name: name,
	}

	return &out
}

// Type returns the type
func (obj *variableDeclaration) Type() Type {
	return obj.typ
}

// Name returns the name
func (obj *variableDeclaration) Name() string {
	return obj.name
}
