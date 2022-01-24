package cardinality

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	specificAdapter := NewSpecificAdapter()
	builder := NewBuilder()
	nonZeroMultipleCharacter := "+"
	multipleCharacter := "*"
	optionalCharacter := "?"
	return createAdapter(
		specificAdapter,
		builder,
		nonZeroMultipleCharacter,
		multipleCharacter,
		optionalCharacter,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewSpecificAdapter creates a new specific adapter instance
func NewSpecificAdapter() SpecificAdapter {
	builder := NewBuilder()
	specificBuilder := NewSpecificBuilder()
	rangeBuilder := NewRangeBuilder()
	prefix := "{"
	suffix := "}"
	rangeSeparator := ","
	return createSpecificAdapter(
		builder,
		specificBuilder,
		rangeBuilder,
		prefix,
		suffix,
		rangeSeparator,
	)
}

// NewSpecificBuilder creates a new specific builder
func NewSpecificBuilder() SpecificBuilder {
	return createSpecificBuilder()
}

// NewRangeBuilder creates a new range builder
func NewRangeBuilder() RangeBuilder {
	return createRangeBuilder()
}

// Adapter represents a cardinality adapter
type Adapter interface {
	ToCardinality(content string) (Cardinality, error)
	ToPatternString(cardinality Cardinality) (string, error)
}

// Builder represents a cardinality builder
type Builder interface {
	Create() Builder
	IsOptional() Builder
	IsMandatory() Builder
	IsNonZeroMultiple() Builder
	IsZeroMultiple() Builder
	WithSpecific(specific Specific) Builder
	Now() (Cardinality, error)
}

// Cardinality represents the cardinality of an element
type Cardinality interface {
	IsValid(amount uint) bool
	Delimiter() (uint, *uint)
	IsOptional() bool
	IsMandatory() bool
	IsNonZeroMultiple() bool
	IsZeroMultiple() bool
	IsSpecific() bool
	Specific() Specific
}

// SpecificAdapter represents a specific adapter
type SpecificAdapter interface {
	ToSpecific(content string) (Specific, error)
	ToPatternString(specific Specific) (string, error)
}

// SpecificBuilder represents a specific cardinality builder
type SpecificBuilder interface {
	Create() SpecificBuilder
	WithAmount(amount uint) SpecificBuilder
	WithRange(rnge Range) SpecificBuilder
	Now() (Specific, error)
}

// Specific represents a specific cardinaltiy
type Specific interface {
	IsValid(amount uint) bool
	Delimiter() (uint, *uint)
	IsAmount() bool
	Amount() *uint
	IsRange() bool
	Range() Range
}

// RangeBuilder represents a range builder
type RangeBuilder interface {
	Create() RangeBuilder
	WithMinimum(min uint) RangeBuilder
	WithMaximum(max uint) RangeBuilder
	Now() (Range, error)
}

// Range represents a range
type Range interface {
	Min() uint
	HasMax() bool
	Max() *uint
}
