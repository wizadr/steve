package contents

import "errors"

type builder struct {
	content       string
	isPrefixLegal bool
	isSuffixLegal bool
}

func createBuilder() Builder {
	out := builder{
		content:       "",
		isPrefixLegal: false,
		isSuffixLegal: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content string) Builder {
	app.content = content
	return app
}

// IsPrefixLegal flags the prefix as legal in the builder
func (app *builder) IsPrefixLegal() Builder {
	app.isPrefixLegal = true
	return app
}

// IsSuffixLegal flags the suffix as legal in the builder
func (app *builder) IsSuffixLegal() Builder {
	app.isSuffixLegal = true
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.content == "" {
		return nil, errors.New("the content is mandatory in order to build a Content instance")
	}

	if app.isPrefixLegal && app.isSuffixLegal {
		return createContentWithPrefixAndSuffix(app.content), nil
	}

	if app.isPrefixLegal {
		return createContentWithPrefix(app.content), nil
	}

	if app.isSuffixLegal {
		return createContentWithSuffix(app.content), nil
	}

	return createContent(app.content), nil
}
