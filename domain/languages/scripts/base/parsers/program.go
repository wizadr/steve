package parsers

type program struct {
	instructions Instructions
	variables    HeaderVariables
}

func createProgram(
	instructions Instructions,
) Program {
	return createProgramInternally(instructions, nil)
}

func createProgramWithVariables(
	instructions Instructions,
	variables HeaderVariables,
) Program {
	return createProgramInternally(instructions, variables)
}

func createProgramInternally(
	instructions Instructions,
	variables HeaderVariables,
) Program {
	out := program{
		instructions: instructions,
		variables:    variables,
	}

	return &out
}

// Instructions returns the instructions
func (obj *program) Instructions() Instructions {
	return obj.instructions
}

// HasVariables returns true if there is variables, false otherwise
func (obj *program) HasVariables() bool {
	return obj.variables != nil
}

// Variables returns the variables, if any
func (obj *program) Variables() HeaderVariables {
	return obj.variables
}
