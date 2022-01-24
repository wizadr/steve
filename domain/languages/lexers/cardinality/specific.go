package cardinality

type specific struct {
	amount *uint
	rnge   Range
}

func createSpecificWithAmount(
	amount *uint,
) Specific {
	return createSpecificInternally(amount, nil)
}

func createSpecificWithRange(
	rnge Range,
) Specific {
	return createSpecificInternally(nil, rnge)
}

func createSpecificInternally(
	amount *uint,
	rnge Range,
) Specific {
	out := specific{
		amount: amount,
		rnge:   rnge,
	}

	return &out
}

// IsValid returns true if the cardinality is valid against the given amount, false otherwise
func (obj *specific) IsValid(amount uint) bool {
	min, pMax := obj.Delimiter()
	matchMin := amount >= min
	if pMax == nil {
		return matchMin
	}

	matchMax := amount <= *pMax
	return matchMin && matchMax
}

// Delimiter returns the delimiter
func (obj *specific) Delimiter() (uint, *uint) {
	if obj.IsAmount() {
		pAmount := obj.Amount()
		return *pAmount, pAmount
	}

	min := obj.rnge.Min()
	return min, obj.rnge.Max()
}

// IsAmount returns true if there is an amount, false otherwise
func (obj *specific) IsAmount() bool {
	return obj.amount != nil
}

// Amount returns the amount, if any
func (obj *specific) Amount() *uint {
	return obj.amount
}

// IsRange returns true if there is a range, false otherwise
func (obj *specific) IsRange() bool {
	return obj.rnge != nil
}

// Range returns the range, if any
func (obj *specific) Range() Range {
	return obj.rnge
}
