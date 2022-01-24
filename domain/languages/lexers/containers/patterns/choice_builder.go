package patterns

import "errors"

type choiceBuilder struct {
	serie     Serie
	isReverse bool
}

func createChoiceBuilder() ChoiceBuilder {
	out := choiceBuilder{
		serie:     nil,
		isReverse: false,
	}

	return &out
}

// Create initializes the builder
func (app *choiceBuilder) Create() ChoiceBuilder {
	return createChoiceBuilder()
}

// WithSerie adds a serie to the builder
func (app *choiceBuilder) WithSerie(serie Serie) ChoiceBuilder {
	app.serie = serie
	return app
}

// IsReverse flags the builder as reverse, false otherwise
func (app *choiceBuilder) IsReverse() ChoiceBuilder {
	app.isReverse = true
	return app
}

// Now builds a new Choice instance
func (app *choiceBuilder) Now() (Choice, error) {
	if app.serie == nil {
		return nil, errors.New("the serie is mandatory in order to build a Choice instance")
	}

	return createChoice(app.serie, app.isReverse), nil
}
