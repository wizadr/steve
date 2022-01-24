package suites

type line struct {
	index     uint
	isSuccess bool
}

func createLine(
	index uint,
	isSuccess bool,
) Line {
	out := line{
		index:     index,
		isSuccess: isSuccess,
	}

	return &out
}

// Index returns the index
func (obj *line) Index() uint {
	return obj.index
}

// IsSuccess returns true if the line is a success, false otherwise
func (obj *line) IsSuccess() bool {
	return obj.isSuccess
}
