package links

import (
	"errors"
)

type linkBuilder struct {
	localToken     string
	referenceToken string
	include        Include
}

func createLinkBuilder() LinkBuilder {
	out := linkBuilder{
		localToken:     "",
		referenceToken: "",
		include:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder()
}

// WithLocalToken adds a local token to the builder
func (app *linkBuilder) WithLocalToken(local string) LinkBuilder {
	app.localToken = local
	return app
}

// WithReferenceToken adds a reference token to the builder
func (app *linkBuilder) WithReferenceToken(reference string) LinkBuilder {
	app.referenceToken = reference
	return app
}

// WithInclude adds an include to the builder
func (app *linkBuilder) WithInclude(include Include) LinkBuilder {
	app.include = include
	return app
}

// Now builds a new Reference instance
func (app *linkBuilder) Now() (Link, error) {
	if app.localToken == "" {
		return nil, errors.New("the local token is mandatory in order to build a Link instance")
	}

	if app.referenceToken == "" {
		return nil, errors.New("the reference token is mandatory in order to build a Link instance")
	}

	if app.include == nil {
		return nil, errors.New("the include instance is mandatory in order to build a Link instance")
	}

	return createReference(app.localToken, app.referenceToken, app.include), nil
}
