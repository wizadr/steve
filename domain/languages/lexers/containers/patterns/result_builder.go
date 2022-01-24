package patterns

import "errors"

type resultBuilder struct {
	input       string
	discoveries Discoveries
	remaining   string
	next        Next
}

func createResultBuilder() ResultBuilder {
	out := resultBuilder{
		input:       "",
		discoveries: nil,
		remaining:   "",
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *resultBuilder) Create() ResultBuilder {
	return createResultBuilder()
}

// WithInput adds the input to the builder
func (app *resultBuilder) WithInput(input string) ResultBuilder {
	app.input = input
	return app
}

// WithDiscoveries add discoveries to the builder
func (app *resultBuilder) WithDiscoveries(discoveries Discoveries) ResultBuilder {
	app.discoveries = discoveries
	return app
}

// WithRemaining adds a remaining to the builder
func (app *resultBuilder) WithRemaining(remaining string) ResultBuilder {
	app.remaining = remaining
	return app
}

// WithNext adds a next to the builder
func (app *resultBuilder) WithNext(next Next) ResultBuilder {
	app.next = next
	return app
}

// Now builds a new Result instance
func (app *resultBuilder) Now() (Result, error) {
	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Result instance")
	}

	if app.discoveries == nil {
		return nil, errors.New("the discoveries is mandatory in order to build a Result instance")
	}

	if app.remaining != "" && app.next != nil {
		return createResultWithRemainingAndNext(app.input, app.discoveries, app.remaining, app.next), nil
	}

	if app.remaining != "" {
		return createResultWithRemaining(app.input, app.discoveries, app.remaining), nil
	}

	if app.next != nil {
		return createResultWithNext(app.input, app.discoveries, app.next), nil
	}

	return createResult(app.input, app.discoveries), nil
}
