package parsers

import "errors"

type headerVariableBuilder struct {
	isMandatory bool
	isInput     bool
	declaration VariableDeclaration
	keyname     string
}

func createHeaderVariableBuilder() HeaderVariableBuilder {
	out := headerVariableBuilder{
		isMandatory: false,
		isInput:     false,
		declaration: nil,
		keyname:     "",
	}

	return &out
}

// Create initializes the builder
func (app *headerVariableBuilder) Create() HeaderVariableBuilder {
	return createHeaderVariableBuilder()
}

// WithDeclaration adds a declaration to the builder
func (app *headerVariableBuilder) WithDeclaration(declaration VariableDeclaration) HeaderVariableBuilder {
	app.declaration = declaration
	return app
}

// WithKeyname adds a keyname to the builder
func (app *headerVariableBuilder) WithKeyname(keyname string) HeaderVariableBuilder {
	app.keyname = keyname
	return app
}

// IsMandatory flags the builder as mandatory
func (app *headerVariableBuilder) IsMandatory() HeaderVariableBuilder {
	app.isMandatory = true
	return app
}

// IsInput flags the builder as input
func (app *headerVariableBuilder) IsInput() HeaderVariableBuilder {
	app.isInput = true
	return app
}

// Now builds a new HeaderVariable instance
func (app *headerVariableBuilder) Now() (HeaderVariable, error) {
	if app.declaration == nil {
		return nil, errors.New("the VariableDeclaration is mandatory in order to build an HeaderVariable instance")
	}

	if app.keyname == "" {
		return nil, errors.New("the keyname is mandatory in order to build an HeaderVariable instance")
	}

	return createHeaderVariable(app.isMandatory, app.isInput, app.declaration, app.keyname), nil
}
