package tokens

type skipMatch struct {
	content SkipMatchContent
}

func createSkipMatch(
	content SkipMatchContent,
) SkipMatch {
	out := skipMatch{
		content: content,
	}

	return &out
}

// Name returns the name
func (obj *skipMatch) Name() string {
	if obj.content.IsRule() {
		return obj.content.Rule().Rule().Base().Name()
	}

	return obj.content.Token().Path().Element().Name()
}

// Input returns the input
func (obj *skipMatch) Input() string {
	if obj.content.IsRule() {
		return obj.content.Rule().Result().Path().Base().Code()
	}

	return obj.content.Token().Result().Input()
}

// Content returns the content
func (obj *skipMatch) Content() SkipMatchContent {
	return obj.content
}
