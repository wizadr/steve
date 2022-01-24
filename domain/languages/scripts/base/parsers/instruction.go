package parsers

type instruction struct {
	assignment          Assignment
	variableDeclaration VariableDeclaration
	ifDeclaration       IfDeclaration
	forDeclaration      ForDeclaration
	voidFuncCall        VoidFuncCall
	direction           Direction
	intIncrement        IntIncrement
}

func createInstructionWithAssignment(
	assignment Assignment,
) Instruction {
	return createInstructionInternally(assignment, nil, nil, nil, nil, nil, nil)
}

func createInstructionWithVariableDeclaration(
	variableDeclaration VariableDeclaration,
) Instruction {
	return createInstructionInternally(nil, variableDeclaration, nil, nil, nil, nil, nil)
}

func createInstructionWithIfDeclaration(
	ifDeclaration IfDeclaration,
) Instruction {
	return createInstructionInternally(nil, nil, ifDeclaration, nil, nil, nil, nil)
}

func createInstructionWithForDeclaration(
	forDeclaration ForDeclaration,
) Instruction {
	return createInstructionInternally(nil, nil, nil, forDeclaration, nil, nil, nil)
}

func createInstructionWithFuncCall(
	voidFuncCall VoidFuncCall,
) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, voidFuncCall, nil, nil)
}

func createInstructionWithDirection(
	direction Direction,
) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, direction, nil)
}

func createInstructionWithIntIncrement(
	intIncrement IntIncrement,
) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, intIncrement)
}

func createInstructionInternally(
	assignment Assignment,
	variableDeclaration VariableDeclaration,
	ifDeclaration IfDeclaration,
	forDeclaration ForDeclaration,
	voidFuncCall VoidFuncCall,
	direction Direction,
	intIncrement IntIncrement,
) Instruction {
	out := instruction{
		assignment:          assignment,
		variableDeclaration: variableDeclaration,
		ifDeclaration:       ifDeclaration,
		forDeclaration:      forDeclaration,
		voidFuncCall:        voidFuncCall,
		direction:           direction,
		intIncrement:        intIncrement,
	}

	return &out
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() Assignment {
	return obj.assignment
}

// IsVariableDeclaration returns true if there is a variableDeclaration, false otherwise
func (obj *instruction) IsVariableDeclaration() bool {
	return obj.variableDeclaration != nil
}

// VariableDeclaration returns the varialeDeclaration, if any
func (obj *instruction) VariableDeclaration() VariableDeclaration {
	return obj.variableDeclaration
}

// IsIfDeclaration returns true if there is an ifDeclaration, false otherwise
func (obj *instruction) IsIfDeclaration() bool {
	return obj.ifDeclaration != nil
}

// IfDeclaration returns the ifDeclaration, if any
func (obj *instruction) IfDeclaration() IfDeclaration {
	return obj.ifDeclaration
}

// IsForDeclaration returns true if there is a forDeclaration, false otherwise
func (obj *instruction) IsForDeclaration() bool {
	return obj.forDeclaration != nil
}

// ForDeclaration returns the forDeclaration, if any
func (obj *instruction) ForDeclaration() ForDeclaration {
	return obj.forDeclaration
}

// IsFuncCall returns true if there is a voidFuncCall, false otherwise
func (obj *instruction) IsFuncCall() bool {
	return obj.voidFuncCall != nil
}

// FuncCall returns the voidFuncCall, if any
func (obj *instruction) FuncCall() VoidFuncCall {
	return obj.voidFuncCall
}

// IsDirection returns true if there is a direction, false otherwise
func (obj *instruction) IsDirection() bool {
	return obj.direction != nil
}

// Direction returns the direction, if any
func (obj *instruction) Direction() Direction {
	return obj.direction
}

// IsIntIncrement adds an intIncrement to the builder
func (obj *instruction) IsIntIncrement() bool {
	return obj.intIncrement != nil
}

// IntIncrement returns the intIncrement, if any
func (obj *instruction) IntIncrement() IntIncrement {
	return obj.intIncrement
}
