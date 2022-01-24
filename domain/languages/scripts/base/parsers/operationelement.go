package parsers

type operationElement struct {
	boolValue       *bool
	variableName    *string
	computableValue Value
	operation       Operation
	fnCall          OperationFuncCall
}

func createOperationElementWithBoolValue(
	boolValue *bool,
) OperationElement {
	return createOperationElementInternally(boolValue, nil, nil, nil, nil)
}

func createOperationElementWithVariableName(
	variableName *string,
) OperationElement {
	return createOperationElementInternally(nil, variableName, nil, nil, nil)
}

func createOperationElementWithComputableValue(
	computableValue Value,
) OperationElement {
	return createOperationElementInternally(nil, nil, computableValue, nil, nil)
}

func createOperationElementWithOperation(
	operation Operation,
) OperationElement {
	return createOperationElementInternally(nil, nil, nil, operation, nil)
}

func createOperationElementWithFuncCall(
	fnCall OperationFuncCall,
) OperationElement {
	return createOperationElementInternally(nil, nil, nil, nil, fnCall)
}

func createOperationElementInternally(
	boolValue *bool,
	variableName *string,
	computableValue Value,
	operation Operation,
	fnCall OperationFuncCall,
) OperationElement {
	out := operationElement{
		boolValue:       boolValue,
		variableName:    variableName,
		computableValue: computableValue,
		operation:       operation,
		fnCall:          fnCall,
	}

	return &out
}

// IsBoolValue returns true if there is a boolValue, false otherwise
func (obj *operationElement) IsBoolValue() bool {
	return obj.boolValue != nil
}

// BoolValue returns the boolValue, if any
func (obj *operationElement) BoolValue() *bool {
	return obj.boolValue
}

// IsVariableName returns true if there is a variableName, false otherwise
func (obj *operationElement) IsVariableName() bool {
	return obj.variableName != nil
}

// VariableName returns the variableName, if any
func (obj *operationElement) VariableName() *string {
	return obj.variableName
}

// IsComputableValue returns true if there is a computableValue, false otherwise
func (obj *operationElement) IsComputableValue() bool {
	return obj.computableValue != nil
}

// ComputableValue returns the computableValue, if any
func (obj *operationElement) ComputableValue() Value {
	return obj.computableValue
}

// IsOperation returns true if there is an operation, false otherwise
func (obj *operationElement) IsOperation() bool {
	return obj.operation != nil
}

// Operation returns the operation, if any
func (obj *operationElement) Operation() Operation {
	return obj.operation
}

// IsFuncCall returns true if there is a funcCall, false otherwise
func (obj *operationElement) IsFuncCall() bool {
	return obj.fnCall != nil
}

// FuncCall returns the funcCall, if any
func (obj *operationElement) FuncCall() OperationFuncCall {
	return obj.fnCall
}
