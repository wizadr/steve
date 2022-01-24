package parsers

import "errors"

type instructionsBuilder struct {
	list []Instruction
}

func createInstructionsBuilder() InstructionsBuilder {
	out := instructionsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionsBuilder) Create() InstructionsBuilder {
	return createInstructionsBuilder()
}

// WithInstructions add instructions to the builder
func (app *instructionsBuilder) WithInstructions(instructions []Instruction) InstructionsBuilder {
	app.list = instructions
	return app
}

// Now builds a new Instructions instance
func (app *instructionsBuilder) Now() (Instructions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Instruction in order to build an Instructions instance")
	}

	return createInstructions(app.list), nil
}
