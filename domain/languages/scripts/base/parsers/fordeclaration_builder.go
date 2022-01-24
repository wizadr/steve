package parsers

import "errors"

type forDeclarationBuilder struct {
	statement    ForStatement
	instructions Instructions
}

func createForDeclarationBuilder() ForDeclarationBuilder {
	out := forDeclarationBuilder{
		statement:    nil,
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *forDeclarationBuilder) Create() ForDeclarationBuilder {
	return createForDeclarationBuilder()
}

// WithStatement adds a statement to the builder
func (app *forDeclarationBuilder) WithStatement(statement ForStatement) ForDeclarationBuilder {
	app.statement = statement
	return app
}

// WithInstructions add instructions to the builder
func (app *forDeclarationBuilder) WithInstructions(instructions Instructions) ForDeclarationBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new ForDeclaration instance
func (app *forDeclarationBuilder) Now() (ForDeclaration, error) {
	if app.statement == nil {
		return nil, errors.New("the ForStatement is mandatory in order to build a ForDeclaration instance")
	}

	if app.instructions != nil {
		return createForDeclarationWithInstructions(app.statement, app.instructions), nil
	}

	return createForDeclaration(app.statement), nil
}
