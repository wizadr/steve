package mistakes

import "errors"

type containsPrefixBuilder struct {
	prefix string
}

func createContainsPrefixBuilder() ContainsPrefixBuilder {
	out := containsPrefixBuilder{
		prefix: "",
	}

	return &out
}

// Create initializes the builder
func (app *containsPrefixBuilder) Create() ContainsPrefixBuilder {
	return createContainsPrefixBuilder()
}

// WithPrefix adds a prefix to the builder
func (app *containsPrefixBuilder) WithPrefix(prefix string) ContainsPrefixBuilder {
	app.prefix = prefix
	return app
}

// Now builds a new ContainsPrefix instance
func (app *containsPrefixBuilder) Now() (ContainsPrefix, error) {
	if app.prefix == "" {
		return nil, errors.New("the prefix is mandatory in order to build a ContainsPrefix instance")
	}

	return createContainsPrefix(app.prefix), nil
}
