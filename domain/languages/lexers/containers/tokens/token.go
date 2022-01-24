package tokens

type token struct {
	name      string
	block     Block
	testSuite TestSuite
}

func createToken(
	name string,
	block Block,
) Token {
	return createTokenInternally(name, block, nil)
}

func createTokenWithTestSuite(
	name string,
	block Block,
	testSuite TestSuite,
) Token {
	return createTokenInternally(name, block, testSuite)
}

func createTokenInternally(
	name string,
	block Block,
	testSuite TestSuite,
) Token {
	out := token{
		name:      name,
		block:     block,
		testSuite: testSuite,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Block returns the block
func (obj *token) Block() Block {
	return obj.block
}

// HasTestSuite returns true if there is a testSuite, false otherwise
func (obj *token) HasTestSuite() bool {
	return obj.testSuite != nil
}

// TestSuite returns the testSuite, if any
func (obj *token) TestSuite() TestSuite {
	return obj.testSuite
}
