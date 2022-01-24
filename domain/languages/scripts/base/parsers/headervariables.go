package parsers

type headerVariables struct {
	list []HeaderVariable
}

func createHeaderVariables(
	list []HeaderVariable,
) HeaderVariables {
	out := headerVariables{
		list: list,
	}

	return &out
}

// All returns the variables
func (obj *headerVariables) All() []HeaderVariable {
	return obj.list
}
