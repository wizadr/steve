package contents

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithContent(content string) Builder
	IsPrefixLegal() Builder
	IsSuffixLegal() Builder
	Now() (Content, error)
}

// Content represents content
type Content interface {
	Content() string
	IsPrefixLegal() bool
	IsSuffixLegal() bool
}
