package tokens

import "errors"

type testSuiteBuilder struct {
	lines []string
}

func createTestSuiteBuilder() TestSuiteBuilder {
	out := testSuiteBuilder{
		lines: nil,
	}

	return &out
}

// Create initializes the builder
func (app *testSuiteBuilder) Create() TestSuiteBuilder {
	return createTestSuiteBuilder()
}

// WithLines add lines to the builder
func (app *testSuiteBuilder) WithLines(lines []string) TestSuiteBuilder {
	app.lines = lines
	return app
}

// Now builds a new TestSuite instance
func (app *testSuiteBuilder) Now() (TestSuite, error) {
	if app.lines != nil && len(app.lines) <= 0 {
		app.lines = nil
	}

	if app.lines == nil {
		return nil, errors.New("there must be at least 1 TestLine in order to build a TestSuite instance")
	}

	return createTestSuite(app.lines), nil
}
