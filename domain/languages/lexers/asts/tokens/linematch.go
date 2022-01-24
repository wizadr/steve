package tokens

type lineMatch struct {
	index   uint
	matches Matches
}

func createLineMatch(
	index uint,
	matches Matches,
) LineMatch {
	out := lineMatch{
		index:   index,
		matches: matches,
	}

	return &out
}

// Index returns the index
func (obj *lineMatch) Index() uint {
	return obj.index
}

// Matches returns the matches
func (obj *lineMatch) Matches() Matches {
	return obj.matches
}
