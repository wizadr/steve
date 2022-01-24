package parsers

type listValue struct {
	assignables  Assignables
	variableName string
}

func createListValueWithAssignables(
	assignables Assignables,
) ListValue {
	return createListValueInternally(assignables, "")
}

func createListValueWithVariableName(
	variableName string,
) ListValue {
	return createListValueInternally(nil, variableName)
}

func createListValueInternally(
	assignables Assignables,
	variableName string,
) ListValue {
	out := listValue{
		assignables:  assignables,
		variableName: variableName,
	}

	return &out
}

// IsAssignables returns true if there is assignables, false otherwise
func (obj *listValue) IsAssignables() bool {
	return obj.assignables != nil
}

// Assignables returns the assignables, if any
func (obj *listValue) Assignables() Assignables {
	return obj.assignables
}

// IsVariableName returns true if there is a variableName, false otherwise
func (obj *listValue) IsVariableName() bool {
	return obj.variableName != ""
}

// VariableName returns the variableName, if any
func (obj *listValue) VariableName() string {
	return obj.variableName
}
