package parsers

type assignee struct {
	first AssigneeNameFirsts
	name  AssigneeNames
}

func createAssignee(
	name AssigneeNames,
) Assignee {
	return createAssigneeInternally(nil, name)
}

func createAssigneeWithFirst(
	first AssigneeNameFirsts,
) Assignee {
	return createAssigneeInternally(first, nil)
}

func createAssigneeInternally(
	first AssigneeNameFirsts,
	name AssigneeNames,
) Assignee {
	out := assignee{
		first: first,
		name:  name,
	}

	return &out
}

// IsFirst returns true if assigneeNameFirst is valid, false otherwiose
func (obj *assignee) IsFirst() bool {
	return obj.first != nil
}

// First returns the assigneeNameFirst, if any
func (obj *assignee) First() AssigneeNameFirsts {
	return obj.first
}

// IsName returns true if assigneeName is valid, false otherwiose
func (obj *assignee) IsName() bool {
	return obj.name != nil
}

// Name returns the assigneeName, if any
func (obj *assignee) Name() AssigneeNames {
	return obj.name
}
