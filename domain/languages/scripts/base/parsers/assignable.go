package parsers

type assignable struct {
	list   TotalListAssignable
	single SingleAssignable
}

func createAssignableWithList(
	list TotalListAssignable,
) Assignable {
	return createAssignableInternally(list, nil)
}

func createAssignableWithSingle(
	single SingleAssignable,
) Assignable {
	return createAssignableInternally(nil, single)
}

func createAssignableInternally(
	list TotalListAssignable,
	single SingleAssignable,
) Assignable {
	out := assignable{
		list:   list,
		single: single,
	}

	return &out
}

// IsList returns true if there is a list, false otherwise
func (obj *assignable) IsList() bool {
	return obj.list != nil
}

// List returns the list, if any
func (obj *assignable) List() TotalListAssignable {
	return obj.list
}

// IsSingle returns true if there is a single, false otherwise
func (obj *assignable) IsSingle() bool {
	return obj.single != nil
}

// Single returns the single, if any
func (obj *assignable) Single() SingleAssignable {
	return obj.single
}
