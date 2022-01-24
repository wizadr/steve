package parsers

import "errors"

type managerBuilder struct {
	resultBuilder ResultBuilder
	events        map[string][]EventFunc
}

func createManagerBuilder(
	resultBuilder ResultBuilder,
) ManagerBuilder {
	out := managerBuilder{
		resultBuilder: resultBuilder,
		events:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *managerBuilder) Create() ManagerBuilder {
	return createManagerBuilder(
		app.resultBuilder,
	)
}

// WithEvents addeventsto the builder
func (app *managerBuilder) WithEvents(events map[string][]EventFunc) ManagerBuilder {
	app.events = events
	return app
}

// Now builds a new manager instance
func (app *managerBuilder) Now() (Manager, error) {
	if app.events != nil && len(app.events) <= 0 {
		app.events = nil
	}

	if app.events == nil {
		return nil, errors.New("there must be at least 1 container that contains at least 1 Event in order to build a Manager instance")
	}

	return createManager(app.resultBuilder, app.events), nil
}
