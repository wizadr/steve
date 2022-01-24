package parsers

type direction struct {
	isBreak    bool
	isContinue bool
}

func createDirectionWithBreak() Direction {
	return createDirectionInternally(true, false)
}

func createDirectionWithContinue() Direction {
	return createDirectionInternally(false, true)
}

func createDirectionInternally(
	isBreak bool,
	isContinue bool,
) Direction {
	out := direction{
		isBreak:    isBreak,
		isContinue: isContinue,
	}

	return &out
}

// IsBreak returns true if there is a break, false otherwise
func (obj *direction) IsBreak() bool {
	return obj.isBreak
}

// IsContinue returns true if there is a continue, false otherwise
func (obj *direction) IsContinue() bool {
	return obj.isContinue
}
