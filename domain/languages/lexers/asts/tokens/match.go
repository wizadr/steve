package tokens

type match struct {
	content MatchContent
}

func createMatch(
	content MatchContent,
) Match {
	out := match{
		content: content,
	}

	return &out
}

// Input returns the input
func (obj *match) Input() string {
	if obj.content.IsRule() {
		return obj.content.Rule().Rule().Base().Code()
	}

	if obj.content.IsToken() {
		return obj.content.Token().Result().Input()
	}

	return obj.content.Skip().Input()
}

// Discoveries returns the matched discoveries
func (obj *match) Discoveries() string {
	if obj.content.IsRule() {
		return obj.content.Rule().Result().Result().Discoveries()
	}

	if obj.content.IsToken() {
		return obj.content.Token().Result().Discoveries()
	}

	// skip has always no discoveries:
	return ""
}

// IsValid returns true if the match is valid, false otherwise
func (obj *match) IsValid() bool {
	if obj.content.IsRule() {
		return obj.content.Rule().IsValid()
	}

	if obj.content.IsToken() {
		return obj.content.Token().IsValid()
	}

	// skip is always valid:
	return true
}

// IsExact returns true if the match is exact, false otherwise
func (obj *match) IsExact() bool {
	if obj.content.IsRule() {
		return obj.content.Rule().IsExact()
	}

	if obj.content.IsToken() {
		return obj.content.Token().IsExact()
	}

	// skip is always exact:
	return true
}

// Content returns the match content
func (obj *match) Content() MatchContent {
	return obj.content
}
