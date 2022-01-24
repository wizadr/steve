package patterns

type group struct {
	name    string
	content GroupContent
}

func createGroup(
	name string,
	content GroupContent,
) Group {
	out := group{
		name:    name,
		content: content,
	}

	return &out
}

// Name returns the name
func (obj *group) Name() string {
	return obj.name
}

// Content returns the content
func (obj *group) Content() GroupContent {
	return obj.content
}
