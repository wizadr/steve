package parsers

type assigneeNames struct {
	list []AssigneeName
}

func createAssigneeNames(
	list []AssigneeName,
) AssigneeNames {
	out := assigneeNames{
		list: list,
	}

	return &out
}

//  All returns the assignee names
func (obj *assigneeNames) All() []AssigneeName {
	return obj.list
}
