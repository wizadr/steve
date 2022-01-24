package patterns

import "errors"

type discoveriesBuilder struct {
	list []Discovery
}

func createDiscoveriesBuilder() DiscoveriesBuilder {
	out := discoveriesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *discoveriesBuilder) Create() DiscoveriesBuilder {
	return createDiscoveriesBuilder()
}

// WithList adds a list to the builder
func (app *discoveriesBuilder) WithList(list []Discovery) DiscoveriesBuilder {
	app.list = list
	return app
}

// Now builds a new Discoveries instance
func (app *discoveriesBuilder) Now() (Discoveries, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Discovery in order to build a Discoveries instance")
	}

	mp := map[string]Discovery{}
	for _, oneDiscovery := range app.list {
		name := oneDiscovery.Container().Name()
		mp[name] = oneDiscovery
	}

	return createDiscoveries(app.list, mp), nil
}
