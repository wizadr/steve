package links

import "github.com/steve-care-software/steve/domain/languages/lexers/paths"

type include struct {
	name  string
	paths paths.Paths
}

func createInclude(
	name string,
	paths paths.Paths,
) Include {
	out := include{
		name:  name,
		paths: paths,
	}

	return &out
}

// Name returns the name
func (obj *include) Name() string {
	return obj.name
}

// Paths returns the paths
func (obj *include) Paths() paths.Paths {
	return obj.paths
}
