package tokens

import "errors"

type lineBuilder struct {
	instructions []Instruction
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithInstructions add instructions to the builder
func (app *lineBuilder) WithInstructions(instructions []Instruction) LineBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.instructions != nil && len(app.instructions) <= 0 {
		app.instructions = nil
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions are mandatory in order to build a Line instance")
	}

	return createLine(app.instructions), nil
}
