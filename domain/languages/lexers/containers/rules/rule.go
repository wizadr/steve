package rules

type rule struct {
	name    string
	code    string
	element Element
}

func createRule(
	name string,
	code string,
	element Element,
) Rule {
	out := rule{
		name:    name,
		code:    code,
		element: element,
	}

	return &out
}

// Name returns the name
func (obj *rule) Name() string {
	return obj.name
}

// Code returns the code
func (obj *rule) Code() string {
	return obj.code
}

// Element returns the element
func (obj *rule) Element() Element {
	return obj.element
}
