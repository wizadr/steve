package parsers

import "errors"

type ifDeclarationBuilder struct {
	condition    Operation
	instructions Instructions
}

func createIfDeclarationBuilder() IfDeclarationBuilder {
	out := ifDeclarationBuilder{
		condition:    nil,
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ifDeclarationBuilder) Create() IfDeclarationBuilder {
	return createIfDeclarationBuilder()
}

// WithCondition adds a condition to the builder
func (app *ifDeclarationBuilder) WithCondition(condition Operation) IfDeclarationBuilder {
	app.condition = condition
	return app
}

// WithInstructions add instructions to the builder
func (app *ifDeclarationBuilder) WithInstructions(instructions Instructions) IfDeclarationBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new IfDeclaration instance
func (app *ifDeclarationBuilder) Now() (IfDeclaration, error) {
	if app.condition == nil {
		return nil, errors.New("the condition is mandatory in order to build an IfDeclaration instance")
	}

	if app.instructions != nil {
		return createIfDeclarationWithInstructions(app.condition, app.instructions), nil
	}

	return createIfDeclaration(app.condition), nil
}
