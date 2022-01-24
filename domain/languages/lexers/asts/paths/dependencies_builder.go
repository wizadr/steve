package paths

type dependenciesBuilder struct {
	list []Element
	mp   map[string]Element
}

func createDependenciesBuilder() DependenciesBuilder {
	out := dependenciesBuilder{
		list: nil,
		mp:   nil,
	}

	return &out
}

// Create initializes the dependencies
func (app *dependenciesBuilder) Create() DependenciesBuilder {
	return createDependenciesBuilder()
}

// WithDependencies add dependencies to the builder
func (app *dependenciesBuilder) WithDependencies(list []Element) DependenciesBuilder {
	app.list = list
	return app
}

// Now builds a new Dependencies instance
func (app *dependenciesBuilder) Now() (Dependencies, error) {
	if app.list == nil {
		app.list = []Element{}
	}

	mp := map[string]Element{}
	for _, oneElement := range app.list {
		name := oneElement.Name()
		mp[name] = oneElement
	}

	list := []Element{}
	for _, oneElement := range mp {
		list = append(list, oneElement)
	}

	return createDependencies(list, mp), nil
}
