package parsers

type assigneeNameFirsts struct {
	list []AssigneeNameFirst
}

func createAssigneeNameFirsts(
	list []AssigneeNameFirst,
) AssigneeNameFirsts {
	out := assigneeNameFirsts{
		list: list,
	}

	return &out
}

//  All returns the assignee names
func (obj *assigneeNameFirsts) All() []AssigneeNameFirst {
	return obj.list
}
