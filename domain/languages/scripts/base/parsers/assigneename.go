package parsers

type assigneeName struct {
	isSkip   bool
	variable string
}

func createAssigneeNameWithSkip() AssigneeName {
	return createAssigneeNameInternally(true, "")
}

func createAssigneeNameWithVariable(
	variable string,
) AssigneeName {
	return createAssigneeNameInternally(false, variable)
}

func createAssigneeNameInternally(
	isSkip bool,
	variable string,
) AssigneeName {
	out := assigneeName{
		isSkip:   isSkip,
		variable: variable,
	}

	return &out
}

// IsSkip returns true if skip, false otherwise
func (obj *assigneeName) IsSkip() bool {
	return obj.isSkip
}

// IsVariable returns true if variable, false otherwise
func (obj *assigneeName) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *assigneeName) Variable() string {
	return obj.variable
}
