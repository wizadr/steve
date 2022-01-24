package patterns

type patternContent struct {
	list []Choice
}

func createPatternContent(
	list []Choice,
) PatternContent {
	out := patternContent{
		list: list,
	}

	return &out
}

// List returns the list of choices
func (obj *patternContent) List() []Choice {
	return obj.list
}

// Length returns the amount of choices in the Pattern content
func (obj *patternContent) Length() uint {
	return uint(len(obj.list))
}
