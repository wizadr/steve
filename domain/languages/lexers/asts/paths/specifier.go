package paths

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type specifier struct {
	containerName string
	cardinality   cardinality.Specific
}

func createSpecifier(
	containerName string,
	cardinality cardinality.Specific,
) Specifier {
	out := specifier{
		containerName: containerName,
		cardinality:   cardinality,
	}

	return &out
}

// ContainerName returns the container name
func (obj *specifier) ContainerName() string {
	return obj.containerName
}

// Cardinality returns the cardinality
func (obj *specifier) Cardinality() cardinality.Specific {
	return obj.cardinality
}
