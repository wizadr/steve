package mistakes

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type specifierDoNotMatch struct {
	containerName string
	cardinality   cardinality.Specific
	amount        uint
}

func createSpecifierDoNotMatch(
	containerName string,
	cardinality cardinality.Specific,
	amount uint,
) SpecifierDoNotMatch {
	out := specifierDoNotMatch{
		containerName: containerName,
		cardinality:   cardinality,
		amount:        amount,
	}

	return &out
}

// ContainerName returns the container name
func (obj *specifierDoNotMatch) ContainerName() string {
	return obj.containerName
}

// Cardinality returns the cardinality
func (obj *specifierDoNotMatch) Cardinality() cardinality.Specific {
	return obj.cardinality
}

// Amount returns the amount
func (obj *specifierDoNotMatch) Amount() uint {
	return obj.amount
}
