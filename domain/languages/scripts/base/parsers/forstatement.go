package parsers

type forStatement struct {
	index    AssigneeName
	value    AssigneeName
	iterable TotalListValue
}

func createForStatement(
	index AssigneeName,
	value AssigneeName,
	iterable TotalListValue,
) ForStatement {
	out := forStatement{
		index:    index,
		value:    value,
		iterable: iterable,
	}

	return &out
}

// Index returns the index
func (obj *forStatement) Index() AssigneeName {
	return obj.index
}

// Value returns the value
func (obj *forStatement) Value() AssigneeName {
	return obj.value
}

// Iterable returns the iterable
func (obj *forStatement) Iterable() TotalListValue {
	return obj.iterable
}
