package cardinality

import "errors"

type builder struct {
	isOptional        bool
	isMandatory       bool
	isNonZeroMultiple bool
	isZeroMultiple    bool
	specific          Specific
}

func createBuilder() Builder {
	out := builder{
		isOptional:        false,
		isMandatory:       false,
		isNonZeroMultiple: false,
		isZeroMultiple:    false,
		specific:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsOptional flags the builder as optional
func (app *builder) IsOptional() Builder {
	app.isOptional = true
	return app
}

// IsMandatory flags the builder as mandatory
func (app *builder) IsMandatory() Builder {
	app.isMandatory = true
	return app
}

// IsNonZeroMultiple flags the builder as a non-zero multiple
func (app *builder) IsNonZeroMultiple() Builder {
	app.isNonZeroMultiple = true
	return app
}

// IsZeroMultiple flags the builder as a zero multiple
func (app *builder) IsZeroMultiple() Builder {
	app.isZeroMultiple = true
	return app
}

// WithSpecific adds a specific cardinality to the builder
func (app *builder) WithSpecific(specific Specific) Builder {
	app.specific = specific
	return app
}

// Now builds a new Cardinality instance
func (app *builder) Now() (Cardinality, error) {
	if app.isOptional {
		return createCardinalityWithOptional(), nil
	}

	if app.isMandatory {
		return createCardinalityWithMandatory(), nil
	}

	if app.isNonZeroMultiple {
		return createCardinalityWithNonZeroMultiple(), nil
	}

	if app.isZeroMultiple {
		return createCardinalityWithZeroMultiple(), nil
	}

	if app.specific != nil {
		return createCardinalityWithSpecific(app.specific), nil
	}

	return nil, errors.New("the cardinality instance is invalid")
}
