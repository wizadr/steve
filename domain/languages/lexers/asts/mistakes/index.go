package mistakes

type index struct {
	idx    uint
	line   uint
	column uint
}

func createIndex(
	idx uint,
	line uint,
	column uint,
) Index {
	out := index{
		idx:    idx,
		line:   line,
		column: column,
	}

	return &out
}

// Index returns the index
func (obj *index) Index() uint {
	return obj.idx
}

// Line returns the line
func (obj *index) Line() uint {
	return obj.line
}

// Column returns the column
func (obj *index) Column() uint {
	return obj.column
}
