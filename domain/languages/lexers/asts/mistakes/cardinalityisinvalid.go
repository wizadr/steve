package mistakes

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type cardinalityIsInvalid struct {
	cardinality cardinality.Cardinality
	amount      uint
}

func createCardinalityIsInvalid(
	cardinality cardinality.Cardinality,
	amount uint,
) CardinalityIsInvalid {
	out := cardinalityIsInvalid{
		cardinality: cardinality,
		amount:      amount,
	}

	return &out
}

// Cardinality returns the cardinality
func (obj *cardinalityIsInvalid) Cardinality() cardinality.Cardinality {
	return obj.cardinality
}

// Amount returns the amount
func (obj *cardinalityIsInvalid) Amount() uint {
	return obj.amount
}
