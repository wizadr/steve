package patterns

type pattern struct {
	name    string
	content PatternContent
}

func createPattern(
	name string,
	content PatternContent,
) Pattern {
	out := pattern{
		name:    name,
		content: content,
	}

	return &out
}

// Name returns the name
func (obj *pattern) Name() string {
	return obj.name
}

// Content returns the content
func (obj *pattern) Content() PatternContent {
	return obj.content
}
