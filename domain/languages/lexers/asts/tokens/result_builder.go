package tokens

import "github.com/steve-care-software/steve/domain/languages/lexers/asts/results"

type resultBuilder struct {
	list []results.Result
}

func createResultBuilder() ResultBuilder {
	out := resultBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *resultBuilder) Create() ResultBuilder {
	return createResultBuilder()
}

// WithResults add results to the builder
func (app *resultBuilder) WithResults(results []results.Result) ResultBuilder {
	app.list = results
	return app
}

// Now builds a new Result instance
func (app *resultBuilder) Now() (Result, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list != nil {
		return createResultWithList(app.list), nil
	}

	return createResult(), nil
}
