package mistakes

type containsPrefix struct {
	prefix string
}

func createContainsPrefix(
	prefix string,
) ContainsPrefix {
	out := containsPrefix{
		prefix: prefix,
	}

	return &out
}

// Prefix returns the prefix
func (obj *containsPrefix) Prefix() string {
	return obj.prefix
}
