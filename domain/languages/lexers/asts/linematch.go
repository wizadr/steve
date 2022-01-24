package asts

type lineMatch struct {
	index   uint
	matches []Match
}

func createLineMatch(
	index uint,
) LineMatch {
	return createLineMatchInternally(index, nil)
}

func createLineMatchWithMatches(
	index uint,
	matches []Match,
) LineMatch {
	return createLineMatchInternally(index, matches)
}

func createLineMatchInternally(
	index uint,
	matches []Match,
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

// HasMatches returns true if there is matches, false otherwise
func (obj *lineMatch) HasMatches() bool {
	return obj.matches != nil
}

// Matches returns the matches
func (obj *lineMatch) Matches() []Match {
	return obj.matches
}
