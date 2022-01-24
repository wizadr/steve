package tokens

import "errors"

type matchesBuilder struct {
	list []BlockMatch
}

func createMatchesBuilder() MatchesBuilder {
	out := matchesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *matchesBuilder) Create() MatchesBuilder {
	return createMatchesBuilder()
}

// WithList add matches to the builder
func (app *matchesBuilder) WithList(matches []BlockMatch) MatchesBuilder {
	app.list = matches
	return app
}

// Now builds a new Matches instance
func (app *matchesBuilder) Now() (Matches, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 BlockMatch in order to build a Matches instance")
	}

	return createMatches(app.list), nil
}
