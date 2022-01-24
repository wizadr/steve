package patterns

type next struct {
	isReverse bool
	group     Group
}

func createNext(
	isReverse bool,
	group Group,
) Next {
	out := next{
		isReverse: isReverse,
		group:     group,
	}

	return &out
}

// IsReverse returns true if reverse, false otherwise
func (obj *next) IsReverse() bool {
	return obj.isReverse
}

// Group returns the group
func (obj *next) Group() Group {
	return obj.group
}
