package suites

type suite struct {
	name  string
	lines []Line
}

func createSuite(
	name string,
	lines []Line,
) Suite {
	out := suite{
		name:  name,
		lines: lines,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Lines returns the lines
func (obj *suite) Lines() []Line {
	return obj.lines
}
