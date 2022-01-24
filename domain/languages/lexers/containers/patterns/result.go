package patterns

type result struct {
	input       string
	discoveries Discoveries
	remaining   string
	next        Next
}

func createResult(
	input string,
	discoveries Discoveries,
) Result {
	return createResultInternally(input, discoveries, "", nil)
}

func createResultWithRemaining(
	input string,
	discoveries Discoveries,
	remaining string,
) Result {
	return createResultInternally(input, discoveries, remaining, nil)
}

func createResultWithNext(
	input string,
	discoveries Discoveries,
	next Next,
) Result {
	return createResultInternally(input, discoveries, "", next)
}

func createResultWithRemainingAndNext(
	input string,
	discoveries Discoveries,
	remaining string,
	next Next,
) Result {
	return createResultInternally(input, discoveries, remaining, next)
}

func createResultInternally(
	input string,
	discoveries Discoveries,
	remaining string,
	next Next,
) Result {
	out := result{
		input:       input,
		discoveries: discoveries,
		remaining:   remaining,
		next:        next,
	}

	return &out
}

// Input returns the input
func (obj *result) Input() string {
	return obj.input
}

// Discoveries returns the discoveries
func (obj *result) Discoveries() Discoveries {
	return obj.discoveries
}

// HasRemaining returns true if there is remaining content, false otherwise
func (obj *result) HasRemaining() bool {
	return obj.remaining != ""
}

// Remaining returns the remaining content, if any
func (obj *result) Remaining() string {
	return obj.remaining
}

// HasNext returns true if there is a next, false otherwise
func (obj *result) HasNext() bool {
	return obj.next != nil
}

// Next returns the next, false otherwise
func (obj *result) Next() Next {
	return obj.next
}
