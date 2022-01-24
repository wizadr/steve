package parsers

type assignables struct {
	list []Assignable
}

func createAssignables(
	list []Assignable,
) Assignables {
	out := assignables{
		list: list,
	}

	return &out
}

// IsEmpty returns true if the assignables is empty, false otherwise
func (obj *assignables) IsEmpty() bool {
	return len(obj.list) <= 0
}

// All returns the assignables
func (obj *assignables) All() []Assignable {
	return obj.list
}
