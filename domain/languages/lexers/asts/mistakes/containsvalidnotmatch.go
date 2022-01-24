package mistakes

type containsValidNotMatch struct {
	line uint
}

func createContainsValidNotMatch(
	line uint,
) ContainsValidNotMatch {
	out := containsValidNotMatch{
		line: line,
	}

	return &out
}

// Line returns the line
func (obj *containsValidNotMatch) Line() uint {
	return obj.line
}
