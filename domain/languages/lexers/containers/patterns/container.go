package patterns

type container struct {
	reverse Serie
	group   Group
}

func createContainerWithReverse(
	reverse Serie,
) Container {
	return createContainerInternally(reverse, nil)
}

func createContainerWithGroup(
	group Group,
) Container {
	return createContainerInternally(nil, group)
}

func createContainerInternally(
	reverse Serie,
	group Group,
) Container {
	out := container{
		reverse: reverse,
		group:   group,
	}

	return &out
}

// Name returns the name of the container
func (obj *container) Name() string {
	if obj.IsReverse() {
		return obj.reverse.Name()
	}

	return obj.group.Name()
}

// IsReverse returns true if there is a reverse, false otehrwise
func (obj *container) IsReverse() bool {
	return obj.reverse != nil
}

// Reverse returns the reverse, if any
func (obj *container) Reverse() Serie {
	return obj.reverse
}

// IsGroup returns true if there is a group, false otehrwise
func (obj *container) IsGroup() bool {
	return obj.group != nil
}

// Group returns the group, if any
func (obj *container) Group() Group {
	return obj.group
}
