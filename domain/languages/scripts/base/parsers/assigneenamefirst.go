package parsers

type assigneeNameFirst struct {
	declaration  VariableDeclaration
	assigneeName AssigneeName
}

func createAssigneeNameFirstWithDeclaration(
	declaration VariableDeclaration,
) AssigneeNameFirst {
	return createAssigneeNameFirstInternally(declaration, nil)
}

func createAssigneeNameFirstWithAssigneeName(
	assigneeName AssigneeName,
) AssigneeNameFirst {
	return createAssigneeNameFirstInternally(nil, assigneeName)
}

func createAssigneeNameFirstInternally(
	declaration VariableDeclaration,
	assigneeName AssigneeName,
) AssigneeNameFirst {
	out := assigneeNameFirst{
		declaration:  declaration,
		assigneeName: assigneeName,
	}

	return &out
}

// IsDeclaration returns true if there is a declaration, false otherwise
func (obj *assigneeNameFirst) IsDeclaration() bool {
	return obj.declaration != nil
}

// Declaration returns the declaration, if any
func (obj *assigneeNameFirst) Declaration() VariableDeclaration {
	return obj.declaration
}

// IsAssigneeName returns true if there is an assigneeName, false otherwise
func (obj *assigneeNameFirst) IsAssigneeName() bool {
	return obj.assigneeName != nil
}

// AssigneeName returns the assigneeName, if any
func (obj *assigneeNameFirst) AssigneeName() AssigneeName {
	return obj.assigneeName
}
