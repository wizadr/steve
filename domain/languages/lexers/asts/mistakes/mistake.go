package mistakes

type mistake struct {
	idx     Index
	path    Path
	content Content
}

func createMistake(
	idx Index,
	path Path,
	content Content,
) Mistake {
	out := mistake{
		idx:     idx,
		path:    path,
		content: content,
	}

	return &out
}

// Index returns the index
func (obj *mistake) Index() Index {
	return obj.idx
}

// Path returns the path
func (obj *mistake) Path() Path {
	return obj.path
}

// Content returns the content
func (obj *mistake) Content() Content {
	return obj.content
}
