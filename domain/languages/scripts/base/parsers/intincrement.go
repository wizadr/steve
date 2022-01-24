package parsers

type intIncrement struct {
	variableName string
	increment    Value
}

func createIntIncrement(
	variableName string,
	increment Value,
) IntIncrement {
	out := intIncrement{
		variableName: variableName,
		increment:    increment,
	}

	return &out
}

// VariableName returns the variableName
func (obj *intIncrement) VariableName() string {
	return obj.variableName
}

// Increment returns the increment
func (obj *intIncrement) Increment() Value {
	return obj.increment
}
