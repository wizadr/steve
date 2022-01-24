package paths

import "errors"

type instructionBuilder struct {
	isChannelSwitch bool
	container       Container
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		isChannelSwitch: false,
		container:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// IsChannelSwitch flags the builder as a channel switch
func (app *instructionBuilder) IsChannelSwitch() InstructionBuilder {
	app.isChannelSwitch = true
	return app
}

// WithContainer adds a container to the builder
func (app *instructionBuilder) WithContainer(container Container) InstructionBuilder {
	app.container = container
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.isChannelSwitch {
		return createInstructionWithChannelSwitch(), nil
	}

	if app.container != nil {
		return createInstructionWithContainer(app.container), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
