package patterns

import "errors"

type nextBuilder struct {
	isReverse bool
	group     Group
}

func createNextBuilder() NextBuilder {
	out := nextBuilder{
		isReverse: false,
		group:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *nextBuilder) Create() NextBuilder {
	return createNextBuilder()
}

// WithGroup adds a group to the builder
func (app *nextBuilder) WithGroup(group Group) NextBuilder {
	app.group = group
	return app
}

// IsReverse flags the builder as reverse
func (app *nextBuilder) IsReverse() NextBuilder {
	app.isReverse = true
	return app
}

// Now builds a new Next instance
func (app *nextBuilder) Now() (Next, error) {
	if app.group == nil {
		return nil, errors.New("the group is mandatory in order to build a Next instance")
	}

	return createNext(app.isReverse, app.group), nil
}
