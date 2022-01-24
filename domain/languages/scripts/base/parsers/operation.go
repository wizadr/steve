package parsers

type operation struct {
	element OperationElement
	suffix  OperationSuffix
}

func createOperation(
	element OperationElement,
) Operation {
	return createOperationInternally(element, nil)
}

func createOperationWithSuffix(
	element OperationElement,
	suffix OperationSuffix,
) Operation {
	return createOperationInternally(element, suffix)
}

func createOperationInternally(
	element OperationElement,
	suffix OperationSuffix,
) Operation {
	out := operation{
		element: element,
		suffix:  suffix,
	}

	return &out
}

// Optimize returns the optimized vertion of the operation
func (obj *operation) Optimize() Operation {
	if !obj.HasSuffix() && obj.Element().IsOperation() {
		return obj.Element().Operation()
	}

	return obj
}

// Element returns the element
func (obj *operation) Element() OperationElement {
	return obj.element
}

// HasSuffix returns true if there is a suffix, false otherwise
func (obj *operation) HasSuffix() bool {
	return obj.suffix != nil
}

// Suffix returns the suffix
func (obj *operation) Suffix() OperationSuffix {
	return obj.suffix
}
