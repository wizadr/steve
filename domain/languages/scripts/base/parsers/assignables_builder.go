package parsers

type assignablesBuilder struct {
	list []Assignable
}

func createAssignablesBuilder() AssignablesBuilder {
	out := assignablesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignablesBuilder) Create() AssignablesBuilder {
	return createAssignablesBuilder()
}

// WithAssignables add assignables to the builder
func (app *assignablesBuilder) WithAssignables(assignables []Assignable) AssignablesBuilder {
	app.list = assignables
	return app
}

// Now builds a new Assignables instance
func (app *assignablesBuilder) Now() (Assignables, error) {
	if app.list == nil {
		app.list = []Assignable{}
	}

	return createAssignables(app.list), nil
}
