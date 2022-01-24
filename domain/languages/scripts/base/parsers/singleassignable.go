package parsers

type singleAssignable struct {
	declarableValue  Value
	operation        Operation
	extract          TotalListAssignable
	listFetchElement ListFetchElement
}

func createSingleAssignableWithDeclarableValue(
	declarableValue Value,
) SingleAssignable {
	return createSingleAssignableInternally(declarableValue, nil, nil, nil)
}

func createSingleAssignableWithOperation(
	operation Operation,
) SingleAssignable {
	return createSingleAssignableInternally(nil, operation, nil, nil)
}

func createSingleAssignableWithExtract(
	extract TotalListAssignable,
) SingleAssignable {
	return createSingleAssignableInternally(nil, nil, extract, nil)
}

func createSingleAssignableWithListFetchElement(
	listFetchElement ListFetchElement,
) SingleAssignable {
	return createSingleAssignableInternally(nil, nil, nil, listFetchElement)
}

func createSingleAssignableInternally(
	declarableValue Value,
	operation Operation,
	extract TotalListAssignable,
	listFetchElement ListFetchElement,
) SingleAssignable {
	out := singleAssignable{
		declarableValue:  declarableValue,
		operation:        operation,
		extract:          extract,
		listFetchElement: listFetchElement,
	}

	return &out
}

// IsDeclarableValue returns true if there is a declarableValue, false otherwise
func (obj *singleAssignable) IsDeclarableValue() bool {
	return obj.declarableValue != nil
}

// DeclarableValue returns the declarableValue, if any
func (obj *singleAssignable) DeclarableValue() Value {
	return obj.declarableValue
}

// IsOperation returns true if there is an operation, false otherwise
func (obj *singleAssignable) IsOperation() bool {
	return obj.operation != nil
}

// Operation returns the operation, if any
func (obj *singleAssignable) Operation() Operation {
	return obj.operation
}

// IsExtract returns true if there is an extract, false otherwise
func (obj *singleAssignable) IsExtract() bool {
	return obj.extract != nil
}

// Extract returns the extract, if any
func (obj *singleAssignable) Extract() TotalListAssignable {
	return obj.extract
}

// IsListFetchElement returns true if there is a listFetchElement, false otherwise
func (obj *singleAssignable) IsListFetchElement() bool {
	return obj.listFetchElement != nil
}

// ListFetchElement returns the listFetchElement, if any
func (obj *singleAssignable) ListFetchElement() ListFetchElement {
	return obj.listFetchElement
}
