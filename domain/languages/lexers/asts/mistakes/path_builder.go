package mistakes

import "errors"

type pathBuilder struct {
	rule    string
	token   string
	parents []string
}

func createPathBuilder() PathBuilder {
	out := pathBuilder{
		rule:    "",
		token:   "",
		parents: nil,
	}

	return &out
}

// Create initializes the builder
func (app *pathBuilder) Create() PathBuilder {
	return createPathBuilder()
}

// WithParents add parents to the builder
func (app *pathBuilder) WithParents(parents []string) PathBuilder {
	app.parents = parents
	return app
}

// WithRule add rule to the builder
func (app *pathBuilder) WithRule(rule string) PathBuilder {
	app.rule = rule
	return app
}

// WithToken add token to the builder
func (app *pathBuilder) WithToken(token string) PathBuilder {
	app.token = token
	return app
}

// Now builds a new Path instance
func (app *pathBuilder) Now() (Path, error) {
	var container PathContainer
	if app.rule != "" {
		container = createPathContainerWithRule(app.rule)
	}

	if app.token != "" {
		container = createPathContainerWithToken(app.token)
	}

	if container == nil {
		return nil, errors.New("the rule or token is mandatory in order to build a Path instance")
	}

	if app.parents != nil && len(app.parents) <= 0 {
		app.parents = nil
	}

	if app.parents != nil {
		return createPathWithParents(container, app.parents), nil
	}

	return createPath(container), nil
}
