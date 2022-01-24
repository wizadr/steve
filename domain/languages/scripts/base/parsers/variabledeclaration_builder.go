package parsers

import "errors"

type variableDeclarationBuilder struct {
	typ  Type
	name string
}

func createVariableDeclarationBuilder() VariableDeclarationBuilder {
	out := variableDeclarationBuilder{
		typ:  nil,
		name: "",
	}

	return &out
}

// Create initializes the builder
func (app *variableDeclarationBuilder) Create() VariableDeclarationBuilder {
	return createVariableDeclarationBuilder()
}

// WithType adds a type to the builder
func (app *variableDeclarationBuilder) WithType(typ Type) VariableDeclarationBuilder {
	app.typ = typ
	return app
}

// WithName adds a name to the builder
func (app *variableDeclarationBuilder) WithName(name string) VariableDeclarationBuilder {
	app.name = name
	return app
}

// Now builds a new VariableDeclaration instance
func (app *variableDeclarationBuilder) Now() (VariableDeclaration, error) {
	if app.typ == nil {
		return nil, errors.New("the type is mandatory in order to build a VariableDeclaration instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a VariableDeclaration instance")
	}

	return createVariableDeclaration(app.typ, app.name), nil
}
