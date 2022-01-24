package parsers

type totalListValue struct {
	assignable   TotalListAssignable
	variableName string
}

func createTotalListValueWithAssignable(
	assignable TotalListAssignable,
) TotalListValue {
	return createTotalListValueInternally(assignable, "")
}

func createTotalListValueWithVariableName(
	variableName string,
) TotalListValue {
	return createTotalListValueInternally(nil, variableName)
}

func createTotalListValueInternally(
	assignable TotalListAssignable,
	variableName string,
) TotalListValue {
	out := totalListValue{
		assignable:   assignable,
		variableName: variableName,
	}

	return &out
}

// IsAssignable returns true if there is an assignable, false otherwise
func (obj *totalListValue) IsAssignable() bool {
	return obj.assignable != nil
}

// Assignable returns the assignable, if any
func (obj *totalListValue) Assignable() TotalListAssignable {
	return obj.assignable
}

// IsVariableName returns true if there is a variableName, false otherwise
func (obj *totalListValue) IsVariableName() bool {
	return obj.variableName != ""
}

// VariableName returns the variableName, if any
func (obj *totalListValue) VariableName() string {
	return obj.variableName
}
