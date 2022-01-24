package patterns

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	patternBuilder := NewPatternBuilder()
	choiceBuilder := NewChoiceBuilder()
	serieBuilder := NewSerieBuilder()
	groupsBuilder := NewGroupsBuilder()
	groupBuilder := NewGroupBuilder()
	resultBuilder := NewResultBuilder()
	nextBuilder := NewNextBuilder()
	discoveriesBuilder := NewDiscoveriesBuilder()
	discoveryBuilder := NewDiscoveryBuilder()
	containerBuilder := NewContainerBuilder()
	groupNamePattern := "[\\$]{1}[A-Z_]+"
	serieNamePattern := "[A-Z]{1}[A-Z_]*"
	reversePrefix := "@"
	patternNamePattern := "[a-z]{1}[a-z_]*"
	delimiter := ","
	prefix := ":"
	suffix := ";"
	stringDelimiter := "\""
	whiteSpacePattern := "[ \t\r\n]*"
	anythingExcept := "[^%s]*"
	escapeReplacementCode := "###"
	return createAdapter(
		builder,
		patternBuilder,
		choiceBuilder,
		serieBuilder,
		groupsBuilder,
		groupBuilder,
		resultBuilder,
		nextBuilder,
		discoveriesBuilder,
		discoveryBuilder,
		containerBuilder,
		groupNamePattern,
		serieNamePattern,
		reversePrefix,
		patternNamePattern,
		delimiter,
		prefix,
		suffix,
		stringDelimiter,
		whiteSpacePattern,
		anythingExcept,
		escapeReplacementCode,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewPatternBuilder creates a new pattern builder
func NewPatternBuilder() PatternBuilder {
	return createPatternBuilder()
}

// NewChoiceBuilder creates anew choice builder
func NewChoiceBuilder() ChoiceBuilder {
	return createChoiceBuilder()
}

// NewSerieBuilder creates a new serie builder
func NewSerieBuilder() SerieBuilder {
	return createSerieBuilder()
}

// NewGroupsBuilder creates a groups builder
func NewGroupsBuilder() GroupsBuilder {
	return createGroupsBuilder()
}

// NewGroupBuilder creates a group builder
func NewGroupBuilder() GroupBuilder {
	return createGroupBuilder()
}

// NewResultBuilder creates a new result builder
func NewResultBuilder() ResultBuilder {
	return createResultBuilder()
}

// NewNextBuilder creates anew next builder
func NewNextBuilder() NextBuilder {
	return createNextBuilder()
}

// NewDiscoveriesBuilder creates a new discoveries builder
func NewDiscoveriesBuilder() DiscoveriesBuilder {
	return createDiscoveriesBuilder()
}

// NewDiscoveryBuilder creates a new discovery builder
func NewDiscoveryBuilder() DiscoveryBuilder {
	return createDiscoveryBuilder()
}

// NewContainerBuilder returns the container builder
func NewContainerBuilder() ContainerBuilder {
	return createContainerBuilder()
}

// Adapter represents a patterns adapter
type Adapter interface {
	ToPatterns(content string) (Patterns, error)
	FromGroupToResult(group Group, content string) (Result, error)
	FromPatternToResult(pattern Pattern, content string) (Result, error)
	FromPatternsToResult(patterns Patterns, content string) (Result, error)
}

// Builder represents a patterns builder
type Builder interface {
	Create() Builder
	WithList(list []Pattern) Builder
	Now() (Patterns, error)
}

// Patterns represents patterns
type Patterns interface {
	All() []Pattern
	First() Pattern
	Find(name string) (Pattern, error)
}

// PatternBuilder represents a pattern builder
type PatternBuilder interface {
	Create() PatternBuilder
	WithName(name string) PatternBuilder
	WithList(list []Choice) PatternBuilder
	Now() (Pattern, error)
}

// Pattern represents a pattern
type Pattern interface {
	Name() string
	Content() PatternContent
}

// PatternContent represents a pattern content
type PatternContent interface {
	List() []Choice
	Length() uint
}

// ChoiceBuilder represents a pattern choice
type ChoiceBuilder interface {
	Create() ChoiceBuilder
	WithSerie(serie Serie) ChoiceBuilder
	IsReverse() ChoiceBuilder
	Now() (Choice, error)
}

// Choice represents a pattern choice
type Choice interface {
	Serie() Serie
	IsReverse() bool
}

// SerieBuilder represents a serie builder
type SerieBuilder interface {
	Create() SerieBuilder
	WithName(name string) SerieBuilder
	WithList(list []Group) SerieBuilder
	Now() (Serie, error)
}

// Serie represents a serie
type Serie interface {
	Name() string
	Content() SerieContent
}

// SerieContent represents a serie content
type SerieContent interface {
	List() []Group
	Length() uint
}

// GroupsBuilder represents a groups builder
type GroupsBuilder interface {
	Create() GroupsBuilder
	WithList(list []Group) GroupsBuilder
	Now() (Groups, error)
}

// Groups represents groups
type Groups interface {
	All() []Group
	Find(name string) (Group, error)
}

// GroupBuilder represents a group builder
type GroupBuilder interface {
	Create() GroupBuilder
	WithName(name string) GroupBuilder
	WithList(list []string) GroupBuilder
	Now() (Group, error)
}

// Group represents a group
type Group interface {
	Name() string
	Content() GroupContent
}

// GroupContent represents a group content
type GroupContent interface {
	List() []string
	First() string
	Length() uint
	Contains(content string) uint
}

// ResultBuilder represents a result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithInput(input string) ResultBuilder
	WithDiscoveries(discoveries Discoveries) ResultBuilder
	WithRemaining(remaining string) ResultBuilder
	WithNext(next Next) ResultBuilder
	Now() (Result, error)
}

// Result represents a result
type Result interface {
	Input() string
	Discoveries() Discoveries
	HasRemaining() bool
	Remaining() string
	HasNext() bool
	Next() Next
}

// NextBuilder represents the next builder
type NextBuilder interface {
	Create() NextBuilder
	WithGroup(group Group) NextBuilder
	IsReverse() NextBuilder
	Now() (Next, error)
}

// Next represents the next possible element
type Next interface {
	IsReverse() bool
	Group() Group
}

// DiscoveriesBuilder represents discoveries builder
type DiscoveriesBuilder interface {
	Create() DiscoveriesBuilder
	WithList(list []Discovery) DiscoveriesBuilder
	Now() (Discoveries, error)
}

// Discoveries represents discoveries
type Discoveries interface {
	Index() uint
	Content() string
	All() []Discovery
	Find(name string) (Discovery, error)
	Amount() uint
	IsAmount(amount uint) bool
	IsInRange(min uint, max uint) bool
	IsAtLeast(min uint) bool
}

// DiscoveryBuilder represents a discovery builder
type DiscoveryBuilder interface {
	Create() DiscoveryBuilder
	WithIndex(index uint) DiscoveryBuilder
	WithContent(content string) DiscoveryBuilder
	WithContainer(container Container) DiscoveryBuilder
	Now() (Discovery, error)
}

// Discovery represents a discovery
type Discovery interface {
	Index() uint
	Content() string
	Container() Container
}

// ContainerBuilder represents a container builder
type ContainerBuilder interface {
	Create() ContainerBuilder
	WithReverse(reverse Serie) ContainerBuilder
	WithGroup(group Group) ContainerBuilder
	Now() (Container, error)
}

// Container represents a container
type Container interface {
	Name() string
	IsReverse() bool
	Reverse() Serie
	IsGroup() bool
	Group() Group
}
