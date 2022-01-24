package patterns

type serie struct {
	name    string
	content SerieContent
}

func createSerie(
	name string,
	content SerieContent,
) Serie {
	out := serie{
		name:    name,
		content: content,
	}

	return &out
}

// Name returns the name
func (obj *serie) Name() string {
	return obj.name
}

// Content returns the content
func (obj *serie) Content() SerieContent {
	return obj.content
}
