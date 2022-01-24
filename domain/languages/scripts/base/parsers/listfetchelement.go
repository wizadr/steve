package parsers

type listFetchElement struct {
	list  TotalListValue
	index IndexValue
}

func createListFetchElement(
	list TotalListValue,
	index IndexValue,
) ListFetchElement {
	out := listFetchElement{
		list:  list,
		index: index,
	}

	return &out
}

// List returns the list
func (obj *listFetchElement) List() TotalListValue {
	return obj.list
}

// Index returns the index
func (obj *listFetchElement) Index() IndexValue {
	return obj.index
}
