package parsers

type totalListAssignable struct {
	listValue ListValue
	slice     Slice
	appnd     Append
}

func createTotalListAssignableWithListValue(
	listValue ListValue,
) TotalListAssignable {
	return createTotalListAssignableInternally(listValue, nil, nil)
}

func createTotalListAssignableWithSlice(
	slice Slice,
) TotalListAssignable {
	return createTotalListAssignableInternally(nil, slice, nil)
}

func createTotalListAssignableWithAppend(
	appnd Append,
) TotalListAssignable {
	return createTotalListAssignableInternally(nil, nil, appnd)
}

func createTotalListAssignableInternally(
	listValue ListValue,
	slice Slice,
	appnd Append,
) TotalListAssignable {
	out := totalListAssignable{
		listValue: listValue,
		slice:     slice,
		appnd:     appnd,
	}

	return &out
}

// IsListValue returns true if there is a listValue, false otherwise
func (obj *totalListAssignable) IsListValue() bool {
	return obj.listValue != nil
}

// ListValue returns the listValue, if any
func (obj *totalListAssignable) ListValue() ListValue {
	return obj.listValue
}

// IsSlice returns true if there is a slice, false otherwise
func (obj *totalListAssignable) IsSlice() bool {
	return obj.slice != nil
}

// Slice returns the slice, if any
func (obj *totalListAssignable) Slice() Slice {
	return obj.slice
}

// IsAppend returns true if there is an append, false otherwise
func (obj *totalListAssignable) IsAppend() bool {
	return obj.appnd != nil
}

// Append returns the append, if any
func (obj *totalListAssignable) Append() Append {
	return obj.appnd
}
