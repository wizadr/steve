package parsers

import "errors"

type valueBuilder struct {
	typeBuilder TypeBuilder
	typ         Type
	element     interface{}
	uintValue   *uint
	intValue    *int
	floatValue  *float64
	boolValue   *bool
	str         *string
}

func createValueBuilder(
	typeBuilder TypeBuilder,
) ValueBuilder {
	out := valueBuilder{
		typeBuilder: typeBuilder,
		typ:         nil,
		element:     nil,
		uintValue:   nil,
		intValue:    nil,
		floatValue:  nil,
		boolValue:   nil,
		str:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder(
		app.typeBuilder,
	)
}

// WithType adds a type to the builder
func (app *valueBuilder) WithType(typ Type) ValueBuilder {
	app.typ = typ
	return app
}

// WithElement adds an element to the builder
func (app *valueBuilder) WithElement(element interface{}) ValueBuilder {
	app.element = element
	return app
}

// WithUint adds a uint to the builder
func (app *valueBuilder) WithUint(uintValue uint) ValueBuilder {
	app.uintValue = &uintValue
	return app
}

// WithInt adds an int to the builder
func (app *valueBuilder) WithInt(intValue int) ValueBuilder {
	app.intValue = &intValue
	return app
}

// WithFloat adds a float to the builder
func (app *valueBuilder) WithFloat(floatValue float64) ValueBuilder {
	app.floatValue = &floatValue
	return app
}

// WithBool adds a bool to the builder
func (app *valueBuilder) WithBool(boolValue bool) ValueBuilder {
	app.boolValue = &boolValue
	return app
}

// WithString adds a string to the builder
func (app *valueBuilder) WithString(str string) ValueBuilder {
	app.str = &str
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.typ == nil {
		typeBuilder := app.typeBuilder.Create().WithDept(0)
		if app.uintValue != nil {
			flag := TypeIntegerFlag | TypeUnsignedFlag
			typeBuilder.WithFlag(flag)
			app.element = *app.uintValue
		}

		if app.intValue != nil {
			flag := TypeIntegerFlag
			typeBuilder.WithFlag(flag)
			app.element = *app.intValue
		}

		if app.floatValue != nil {
			flag := TypeFloatFlag
			typeBuilder.WithFlag(flag)
			app.element = *app.floatValue
		}

		if app.boolValue != nil {
			flag := TypeBoolFlag
			typeBuilder.WithFlag(flag)
			app.element = *app.boolValue
		}

		if app.str != nil {
			flag := TypeStringFlag
			typeBuilder.WithFlag(flag)
			app.element = *app.str
		}

		typ, err := typeBuilder.Now()
		if err != nil {
			return nil, err
		}

		app.typ = typ
	}

	if app.typ == nil {
		return nil, errors.New("the type is mandatory in order to build a Value instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Value instance")
	}

	return createValue(app.typ, app.element), nil
}
