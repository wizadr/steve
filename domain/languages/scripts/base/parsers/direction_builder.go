package parsers

import "errors"

type directionBuilder struct {
	isBreak    bool
	isContinue bool
}

func createDirectionBuilder() DirectionBuilder {
	out := directionBuilder{
		isBreak:    false,
		isContinue: false,
	}

	return &out
}

// Create initializes the builder
func (app *directionBuilder) Create() DirectionBuilder {
	return createDirectionBuilder()
}

// IsBreak flags the builder as break
func (app *directionBuilder) IsBreak() DirectionBuilder {
	app.isBreak = true
	return app
}

// IsContinue flags the builder as continue
func (app *directionBuilder) IsContinue() DirectionBuilder {
	app.isContinue = true
	return app
}

// Now builds a new Direction instance
func (app *directionBuilder) Now() (Direction, error) {
	if app.isBreak {
		return createDirectionWithBreak(), nil
	}

	if app.isContinue {
		return createDirectionWithContinue(), nil
	}

	return nil, errors.New("the Direction is invalid")
}
