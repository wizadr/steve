package paths

type element struct {
	name string
	must []Line
	not  []Line
}

func createElement(
	name string,
	must []Line,
) Element {
	return createElementInternally(name, must, nil)
}

func createElementwithNotContainers(
	name string,
	must []Line,
	not []Line,
) Element {
	return createElementInternally(name, must, not)
}

func createElementInternally(
	name string,
	must []Line,
	not []Line,
) Element {
	out := element{
		name: name,
		must: must,
		not:  not,
	}

	return &out
}

// Name returns the name
func (obj *element) Name() string {
	return obj.name
}

// Must returns the must lines
func (obj *element) Must() []Line {
	return obj.must
}

// HasNot returns true if there is not lines, false otherwise
func (obj *element) HasNot() bool {
	return obj.not != nil
}

// Not returns the not lines, if any
func (obj *element) Not() []Line {
	return obj.not
}
