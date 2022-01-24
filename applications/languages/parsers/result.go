package parsers

type result struct {
	first interface{}
	list  []interface{}
}

func createResult(
	first interface{},
	list []interface{},
) Result {
	out := result{
		first: first,
		list:  list,
	}

	return &out
}

// First returns the first element
func (obj *result) First() interface{} {
	return obj.first
}

// List returns the list of elements
func (obj *result) List() []interface{} {
	return obj.list
}

// IsSingle returns true if result is single, false otherwise
func (obj *result) IsSingle() bool {
	return len(obj.list) <= 1
}
