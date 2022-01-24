package parsers

import "errors"

type programBuilder struct {
	instructions Instructions
	variables    HeaderVariables
}

func createProgramBuilder() ProgramBuilder {
	out := programBuilder{
		instructions: nil,
		variables:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *programBuilder) Create() ProgramBuilder {
	return createProgramBuilder()
}

// WithInstructions add instructions to the builder
func (app *programBuilder) WithInstructions(instructions Instructions) ProgramBuilder {
	app.instructions = instructions
	return app
}

// WithVariables add variables to the builder
func (app *programBuilder) WithVariables(variables HeaderVariables) ProgramBuilder {
	app.variables = variables
	return app
}

// Now builds a new Program instance
func (app *programBuilder) Now() (Program, error) {
	if app.instructions == nil {
		return nil, errors.New("the Instructions is mandatory in order to build a Program instance")
	}

	if app.variables != nil {
		return createProgramWithVariables(app.instructions, app.variables), nil
	}

	return createProgram(app.instructions), nil
}
