package patterns

type serieContent struct {
	list []Group
}

func createSerieContent(
	list []Group,
) SerieContent {
	out := serieContent{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *serieContent) List() []Group {
	return obj.list
}

// Length returns the serie length
func (obj *serieContent) Length() uint {
	return uint(len(obj.list))
}
