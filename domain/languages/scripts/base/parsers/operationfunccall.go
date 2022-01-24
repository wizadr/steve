package parsers

type operationFuncCall struct {
	isNull Operation
}

func createOperationFuncCallWithIsNull(
	isNull Operation,
) OperationFuncCall {
	return createOperationFuncCallInternally(isNull)
}

func createOperationFuncCallInternally(
	isNull Operation,
) OperationFuncCall {
	out := operationFuncCall{
		isNull: isNull,
	}

	return &out
}

// IsNull returns true if the func call is isNull, false otherwise
func (obj *operationFuncCall) IsNull() bool {
	return obj.isNull != nil
}

// Null returns the isNull operation
func (obj *operationFuncCall) Null() Operation {
	return obj.isNull
}
