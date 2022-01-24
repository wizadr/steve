package parsers

import "errors"

type variableDeclarationsBuilder struct {
	list []VariableDeclaration
}

func createVariableDeclarationsBuilder() VariableDeclarationsBuilder {
	out := variableDeclarationsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *variableDeclarationsBuilder) Create() VariableDeclarationsBuilder {
	return createVariableDeclarationsBuilder()
}

// WithDeclarations add declarations to the builder
func (app *variableDeclarationsBuilder) WithDeclarations(declarations []VariableDeclaration) VariableDeclarationsBuilder {
	app.list = declarations
	return app
}

// Now builds a new VariableDeclarations instance
func (app *variableDeclarationsBuilder) Now() (VariableDeclarations, error) {
	if app.list == nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 VariableDeclaration in order to build a VariableDeclarations instance")
	}

	return createVariableDeclarations(app.list), nil
}
