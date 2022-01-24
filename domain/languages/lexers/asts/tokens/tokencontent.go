package tokens

type tokenContent struct {
	must LineMatch
	not  LineMatch
}

func createTokenContent(
	must LineMatch,
) TokenContent {
	return createTokenContentInternally(must, nil)
}

func createTokenContentWithNot(
	must LineMatch,
	not LineMatch,
) TokenContent {
	return createTokenContentInternally(must, not)
}

func createTokenContentInternally(
	must LineMatch,
	not LineMatch,
) TokenContent {
	out := tokenContent{
		must: must,
		not:  not,
	}

	return &out
}

// IsValid returns true if the token is valid, false otherwise
func (obj *tokenContent) IsValid() bool {
	if obj.HasNot() {
		return !obj.Not().Matches().IsValid()
	}

	return obj.must.Matches().IsValid()
}

// IsExact returns true if the token is exact, false otherwise
func (obj *tokenContent) IsExact() bool {
	if obj.HasNot() && obj.Not().Matches().IsValid() {
		return false
	}

	return obj.must.Matches().IsExact()
}

// Discoveries returns the discoveries
func (obj *tokenContent) Discoveries() string {
	if obj.HasNot() && obj.Not().Matches().IsValid() {
		return ""
	}

	return obj.must.Matches().Discoveries()
}

// Must returns the must line match
func (obj *tokenContent) Must() LineMatch {
	return obj.must
}

// HasNot returns true if there is a not line match, false otherwise
func (obj *tokenContent) HasNot() bool {
	return obj.not != nil
}

// Not returns the not line match, if any
func (obj *tokenContent) Not() LineMatch {
	return obj.not
}
