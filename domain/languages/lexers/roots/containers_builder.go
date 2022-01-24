package roots

import "errors"

type containersBuilder struct {
	list []Container
}

func createContainersBuilder() ContainersBuilder {
	out := containersBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *containersBuilder) Create() ContainersBuilder {
	return createContainersBuilder()
}

// WithList adds a list to the builder
func (app *containersBuilder) WithList(list []Container) ContainersBuilder {
	app.list = list
	return app
}

// Now builds a new Containers instance
func (app *containersBuilder) Now() (Containers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Container in order to build a Containers instance")
	}

	mp := map[string]Container{}
	for _, oneContainer := range app.list {
		name := oneContainer.Name()
		mp[name] = oneContainer
	}

	return createContainers(app.list, mp), nil
}
