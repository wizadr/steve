package parsers

type slice struct {
	list      ListValue
	delimiter SliceDelimiter
}

func createSlice(
	list ListValue,
	delimiter SliceDelimiter,
) Slice {
	out := slice{
		list:      list,
		delimiter: delimiter,
	}

	return &out
}

// List returns the list
func (obj *slice) List() ListValue {
	return obj.list
}

// Delimiter returns the delimiter
func (obj *slice) Delimiter() SliceDelimiter {
	return obj.delimiter
}
