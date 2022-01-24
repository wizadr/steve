package mistakes

type path struct {
	container PathContainer
	parents   []string
}

func createPath(
	container PathContainer,
) Path {
	return createPathInternally(container, nil)
}

func createPathWithParents(
	container PathContainer,
	parents []string,
) Path {
	return createPathInternally(container, parents)
}

func createPathInternally(
	container PathContainer,
	parents []string,
) Path {
	out := path{
		container: container,
		parents:   parents,
	}

	return &out
}

// Container returns the container
func (obj *path) Container() PathContainer {
	return obj.container
}

// HasParents returns true if there is parents, false otherwise
func (obj *path) HasParents() bool {
	return obj.parents != nil
}

// Parents retruns the parent tokens, if any
func (obj *path) Parents() []string {
	return obj.parents
}
