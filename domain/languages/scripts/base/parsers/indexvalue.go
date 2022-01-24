package parsers

type indexValue struct {
	value        Value
	variableName string
}

func createIndexValueWithValue(
	value Value,
) IndexValue {
	return createIndexValueInternally(value, "")
}

func createIndexValueWithVariableName(
	variableName string,
) IndexValue {
	return createIndexValueInternally(nil, variableName)
}

func createIndexValueInternally(
	value Value,
	variableName string,
) IndexValue {
	out := indexValue{
		value:        value,
		variableName: variableName,
	}

	return &out
}

// IsValue returns true if there is a value, false otherwise
func (obj *indexValue) IsValue() bool {
	return obj.value != nil
}

// Value retruns the value, if any
func (obj *indexValue) Value() Value {
	return obj.value
}

// IsVariableName returns true if there is a variableName, false otherwise
func (obj *indexValue) IsVariableName() bool {
	return obj.variableName != ""
}

// VariableName retruns the variableName, if any
func (obj *indexValue) VariableName() string {
	return obj.variableName
}
