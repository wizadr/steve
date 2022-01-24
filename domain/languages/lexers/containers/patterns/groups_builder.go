package patterns

import "errors"

type groupsBuilder struct {
	list []Group
}

func createGroupsBuilder() GroupsBuilder {
	out := groupsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *groupsBuilder) Create() GroupsBuilder {
	return createGroupsBuilder()
}

// WithList adds a list to the builder
func (app *groupsBuilder) WithList(list []Group) GroupsBuilder {
	app.list = list
	return app
}

// Now builds a new Groups instance
func (app *groupsBuilder) Now() (Groups, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Group in order to build a Groups instance")
	}

	mp := map[string]Group{}
	for _, oneGroup := range app.list {
		name := oneGroup.Name()
		mp[name] = oneGroup
	}

	return createGroups(app.list, mp), nil
}
