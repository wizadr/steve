package rules

import "errors"

type ruleBuilder struct {
	name    string
	code    string
	element Element
}

func createRuleBuilder() RuleBuilder {
	out := ruleBuilder{
		name:    "",
		code:    "",
		element: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleBuilder) Create() RuleBuilder {
	return createRuleBuilder()
}

// WithName adds a name to the builder
func (app *ruleBuilder) WithName(name string) RuleBuilder {
	app.name = name
	return app
}

// WithCode adds a code to the builder
func (app *ruleBuilder) WithCode(code string) RuleBuilder {
	app.code = code
	return app
}

// WithElement add an element to the builder
func (app *ruleBuilder) WithElement(element Element) RuleBuilder {
	app.element = element
	return app
}

// Now builds a new Rule instance
func (app *ruleBuilder) Now() (Rule, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Rule instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a Rule instance")
	}

	if app.element == nil {
		return nil, errors.New("the Element is mandatory in order to build an Element instance")
	}

	return createRule(app.name, app.code, app.element), nil
}
