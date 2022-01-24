package tokens

type testSuite struct {
	lines []string
}

func createTestSuite(
	lines []string,
) TestSuite {
	out := testSuite{
		lines: lines,
	}

	return &out
}

// Lines returns the test lines
func (obj *testSuite) Lines() []string {
	return obj.lines
}
