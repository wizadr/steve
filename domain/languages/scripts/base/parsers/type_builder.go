package parsers

import "errors"

type typeBuilder struct {
	dept uint
	flag *uint8
}

func createTypeBuilder() TypeBuilder {
	out := typeBuilder{
		dept: 0,
		flag: nil,
	}

	return &out
}

// Create initializes the builder
func (app *typeBuilder) Create() TypeBuilder {
	return createTypeBuilder()
}

// WithDept adds a dept to the builder
func (app *typeBuilder) WithDept(dept uint) TypeBuilder {
	app.dept = dept
	return app
}

// WithFlag adds a flag to the builder
func (app *typeBuilder) WithFlag(flag uint8) TypeBuilder {
	app.flag = &flag
	return app
}

// Now builds a new Type instance
func (app *typeBuilder) Now() (Type, error) {
	if app.flag == nil {
		return nil, errors.New("the flag is mandatory in order to build a Type instance")
	}

	return createType(app.dept, *app.flag), nil
}
