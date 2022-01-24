package cardinality

type cardinality struct {
	isOptional        bool
	isMandatory       bool
	isNonZeroMultiple bool
	isZeroMultiple    bool
	specific          Specific
}

func createCardinalityWithOptional() Cardinality {
	return createCardinalityInternally(true, false, false, false, nil)
}

func createCardinalityWithMandatory() Cardinality {
	return createCardinalityInternally(false, true, false, false, nil)
}

func createCardinalityWithNonZeroMultiple() Cardinality {
	return createCardinalityInternally(false, false, true, false, nil)
}

func createCardinalityWithZeroMultiple() Cardinality {
	return createCardinalityInternally(false, false, false, true, nil)
}

func createCardinalityWithSpecific(specific Specific) Cardinality {
	return createCardinalityInternally(false, false, false, false, specific)
}

func createCardinalityInternally(
	isOptional bool,
	isMandatory bool,
	isNonZeroMultiple bool,
	isZeroMultiple bool,
	specific Specific,
) Cardinality {
	out := cardinality{
		isOptional:        isOptional,
		isMandatory:       isMandatory,
		isNonZeroMultiple: isNonZeroMultiple,
		isZeroMultiple:    isZeroMultiple,
		specific:          specific,
	}

	return &out
}

// IsValid returns true if the cardinality is valid against the given amount, false otherwise
func (obj *cardinality) IsValid(amount uint) bool {
	min, pMax := obj.Delimiter()
	matchMin := amount >= min
	if pMax == nil {
		return matchMin
	}

	matchMax := amount <= *pMax
	return matchMin && matchMax
}

// Delimiter returns the delimiter
func (obj *cardinality) Delimiter() (uint, *uint) {
	// [0,1] = ?
	if obj.IsOptional() {
		amount := uint(1)
		return uint(0), &amount
	}

	// [1,1]
	if obj.IsMandatory() {
		amount := uint(1)
		return amount, &amount
	}

	// [1,inf] = +
	if obj.IsNonZeroMultiple() {
		return uint(1), nil
	}

	// [0,inf] = *
	if obj.IsZeroMultiple() {
		return uint(0), nil
	}

	return obj.specific.Delimiter()
}

// IsOptional returns true if there is an optional, false otherwise
func (obj *cardinality) IsOptional() bool {
	return obj.isOptional
}

// IsNonZeroMultiple returns true if there is a non-zero multiple, false otherwise
func (obj *cardinality) IsNonZeroMultiple() bool {
	return obj.isNonZeroMultiple
}

// IsNonZeroMultiple returns true if there is a zero multiple, false otherwise
func (obj *cardinality) IsZeroMultiple() bool {
	return obj.isZeroMultiple
}

// IsSpecific returns true if there is a specific cardinality, false otherwise
func (obj *cardinality) IsSpecific() bool {
	return obj.specific != nil
}

// Specific returns the specific cardinality, if any
func (obj *cardinality) Specific() Specific {
	return obj.specific
}

// IsMandatory returns true if the cardinality is mandatory, false otherwise
func (obj *cardinality) IsMandatory() bool {
	return obj.isMandatory
}
