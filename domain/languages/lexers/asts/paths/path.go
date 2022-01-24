package paths

type path struct {
	element Element
	dep     Dependencies
}

func createPath(
	element Element,
	dep Dependencies,
) Path {
	out := path{
		element: element,
		dep:     dep,
	}

	return &out
}

// Element returns the element
func (obj *path) Element() Element {
	return obj.element
}

// Dependencies returns the dependencies
func (obj *path) Dependencies() Dependencies {
	return obj.dep
}

// CombinedElements returns the current element combined with the dependencies elements
func (obj *path) CombinedElements() []Element {
	dep := obj.dep.All()
	return append(dep, obj.element)
}
