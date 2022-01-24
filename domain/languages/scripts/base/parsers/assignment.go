package parsers

type assignment struct {
	assignee    Assignee
	assignables Assignables
}

func createAssignment(
	assignee Assignee,
	assignables Assignables,
) Assignment {
	out := assignment{
		assignee:    assignee,
		assignables: assignables,
	}

	return &out
}

// Assignee returns the assignee
func (obj *assignment) Assignee() Assignee {
	return obj.assignee
}

// Assignables returns the assignables
func (obj *assignment) Assignables() Assignables {
	return obj.assignables
}
