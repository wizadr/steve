package links

import "errors"

type builder struct {
	list []Link
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLinks add links to the builder
func (app *builder) WithLinks(links []Link) Builder {
	app.list = links
	return app
}

// Now builds a new References instance
func (app *builder) Now() (Links, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Reference instance in order to build a References instance")
	}

	mpLocalToken := map[string]Link{}
	for _, oneLink := range app.list {
		keyname := oneLink.LocalToken()
		mpLocalToken[keyname] = oneLink
	}

	mpRefToken := map[string]map[string]Link{}
	for _, oneLink := range app.list {
		includeName := oneLink.Include().Name()
		if _, ok := mpRefToken[includeName]; !ok {
			mpRefToken[includeName] = map[string]Link{}
		}

		refToken := oneLink.ReferenceToken()
		mpRefToken[includeName][refToken] = oneLink
	}

	return createLinks(app.list, mpLocalToken, mpRefToken), nil
}
