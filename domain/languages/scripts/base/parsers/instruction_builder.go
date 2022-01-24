package parsers

import "errors"

type instructionBuilder struct {
	assignment          Assignment
	variableDeclaration VariableDeclaration
	ifDeclaration       IfDeclaration
	forDeclaration      ForDeclaration
	voidFuncCall        VoidFuncCall
	direction           Direction
	intIncrement        IntIncrement
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		assignment:          nil,
		variableDeclaration: nil,
		ifDeclaration:       nil,
		forDeclaration:      nil,
		voidFuncCall:        nil,
		direction:           nil,
		intIncrement:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

// WithVariableDeclaration adds a variableDeclaration to the builder
func (app *instructionBuilder) WithVariableDeclaration(variableDeclaration VariableDeclaration) InstructionBuilder {
	app.variableDeclaration = variableDeclaration
	return app
}

// WithIfDeclaration adds an ifDeclaration to the builder
func (app *instructionBuilder) WithIfDeclaration(ifDeclaration IfDeclaration) InstructionBuilder {
	app.ifDeclaration = ifDeclaration
	return app
}

// WithForDeclaration adds a forDeclaration to the builder
func (app *instructionBuilder) WithForDeclaration(forDeclaration ForDeclaration) InstructionBuilder {
	app.forDeclaration = forDeclaration
	return app
}

// WithFuncCall adds a voidFuncCall to the builder
func (app *instructionBuilder) WithFuncCall(funcCall VoidFuncCall) InstructionBuilder {
	app.voidFuncCall = funcCall
	return app
}

// WithDirection adds a direction to the builder
func (app *instructionBuilder) WithDirection(direction Direction) InstructionBuilder {
	app.direction = direction
	return app
}

// WithIntIncrement adds an intIncrement instance
func (app *instructionBuilder) WithIntIncrement(intIncrement IntIncrement) InstructionBuilder {
	app.intIncrement = intIncrement
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.assignment != nil {
		return createInstructionWithAssignment(app.assignment), nil
	}

	if app.variableDeclaration != nil {
		return createInstructionWithVariableDeclaration(app.variableDeclaration), nil
	}

	if app.ifDeclaration != nil {
		return createInstructionWithIfDeclaration(app.ifDeclaration), nil
	}

	if app.forDeclaration != nil {
		return createInstructionWithForDeclaration(app.forDeclaration), nil
	}

	if app.voidFuncCall != nil {
		return createInstructionWithFuncCall(app.voidFuncCall), nil
	}

	if app.direction != nil {
		return createInstructionWithDirection(app.direction), nil
	}

	if app.intIncrement != nil {
		return createInstructionWithIntIncrement(app.intIncrement), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
