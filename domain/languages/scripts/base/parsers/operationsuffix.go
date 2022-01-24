package parsers

type operationSuffix struct {
	operator uint16
	element  OperationElement
}

func createOperationSuffix(
	operator uint16,
	element OperationElement,
) OperationSuffix {
	out := operationSuffix{
		operator: operator,
		element:  element,
	}

	return &out
}

// Operator returns the operator
func (obj *operationSuffix) Operator() uint16 {
	return obj.operator
}

// Element returns the element
func (obj *operationSuffix) Element() OperationElement {
	return obj.element
}
