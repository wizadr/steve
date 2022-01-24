package parsers

import "errors"

type listFetchElementBuilder struct {
	list  TotalListValue
	index IndexValue
}

func createListfetchElementBuilder() ListFetchElementBuilder {
	out := listFetchElementBuilder{
		list:  nil,
		index: nil,
	}

	return &out
}

// Create initializes the builder
func (app *listFetchElementBuilder) Create() ListFetchElementBuilder {
	return createListfetchElementBuilder()
}

// WithList adds a list to the builder
func (app *listFetchElementBuilder) WithList(list TotalListValue) ListFetchElementBuilder {
	app.list = list
	return app
}

// WithIndex adds an index to the builder
func (app *listFetchElementBuilder) WithIndex(index IndexValue) ListFetchElementBuilder {
	app.index = index
	return app
}

// Now builds a new ListFetchElement instance
func (app *listFetchElementBuilder) Now() (ListFetchElement, error) {
	if app.list == nil {
		return nil, errors.New("the list is mandatory in order to build a ListFetchElement instance")
	}

	if app.index == nil {
		return nil, errors.New("the index is mandatory in order tobuild a ListFetchElement instance")
	}

	return createListFetchElement(app.list, app.index), nil
}
