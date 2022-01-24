package paths

type line struct {
	instructions []Instruction
}

func createLine(
	instructions []Instruction,
) Line {
	out := line{
		instructions: instructions,
	}

	return &out
}

// Instructions returns the instructions
func (obj *line) Instructions() []Instruction {
	return obj.instructions
}
