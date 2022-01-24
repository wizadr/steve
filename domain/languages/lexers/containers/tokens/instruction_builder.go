package tokens

import "errors"

type instructionBuilder struct {
	isChannelSwitch bool
	element         Element
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		isChannelSwitch: false,
		element:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// IsChannelSwitch flags the builder as a chanel switch
func (app *instructionBuilder) IsChannelSwitch() InstructionBuilder {
	app.isChannelSwitch = true
	return app
}

// WithElement adds an element to the builder
func (app *instructionBuilder) WithElement(element Element) InstructionBuilder {
	app.element = element
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.element != nil {
		return createInstructionWithElement(app.element), nil
	}

	if app.isChannelSwitch {
		return createInstructionWithChannelSwitch(), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
