package parsers

type instructions struct {
	list []Instruction
}

func createInstructions(
	list []Instruction,
) Instructions {
	out := instructions{
		list: list,
	}

	return &out
}

// All returns the instructions
func (obj *instructions) All() []Instruction {
	return obj.list
}
