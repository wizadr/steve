package parsers

type ifDeclaration struct {
	condition    Operation
	instructions Instructions
}

func createIfDeclaration(
	condition Operation,
) IfDeclaration {
	return createIfDeclarationInternally(condition, nil)
}

func createIfDeclarationWithInstructions(
	condition Operation,
	instructions Instructions,
) IfDeclaration {
	return createIfDeclarationInternally(condition, instructions)
}

func createIfDeclarationInternally(
	condition Operation,
	instructions Instructions,
) IfDeclaration {
	out := ifDeclaration{
		condition:    condition,
		instructions: instructions,
	}

	return &out
}

// Condition returns the condition
func (obj *ifDeclaration) Condition() Operation {
	return obj.condition
}

// HasInstructions returns true if there is instructions, false otherwise
func (obj *ifDeclaration) HasInstructions() bool {
	return obj.instructions != nil
}

// Instructions returns the instructions, if any
func (obj *ifDeclaration) Instructions() Instructions {
	return obj.instructions
}
