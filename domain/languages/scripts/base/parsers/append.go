package parsers

type appnd struct {
	list       ListValue
	assignable SingleAssignable
}

func createAppend(
	list ListValue,
	assignable SingleAssignable,
) Append {
	out := appnd{
		list:       list,
		assignable: assignable,
	}

	return &out
}

// List returns the list
func (obj *appnd) List() ListValue {
	return obj.list
}

// Assignable returns the assignable
func (obj *appnd) Assignable() SingleAssignable {
	return obj.assignable
}
