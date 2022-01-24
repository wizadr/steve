package containers

type container struct {
	patterns string
	rules    string
	tokens   string
}

func createContainer(
	patterns string,
	rules string,
	tokens string,
) Container {
	out := container{
		patterns: patterns,
		rules:    rules,
		tokens:   tokens,
	}

	return &out
}

// Patterns returns the patterns
func (obj *container) Patterns() string {
	return obj.patterns
}

// Rules returns the rules
func (obj *container) Rules() string {
	return obj.rules
}

// Tokens returns the tokens
func (obj *container) Tokens() string {
	return obj.tokens
}
