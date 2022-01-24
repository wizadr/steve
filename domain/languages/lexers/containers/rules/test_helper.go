package rules

// CreateElementWithConstantForTests creates an element with constant for tests
func CreateElementWithConstantForTests(constant string, code string) Element {
	content, err := NewContentBuilder().Create().WithConstant(constant).Now()
	if err != nil {
		panic(err)
	}

	ruleElement, err := NewElementBuilder().Create().WithContent(content).WithCode(code).Now()
	if err != nil {
		panic(err)
	}

	return ruleElement
}
