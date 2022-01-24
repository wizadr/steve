package lexers

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
	asts_paths "github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/contents"
	"github.com/steve-care-software/steve/domain/languages/lexers/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/references"
	resources_containers "github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"
	resources_multiples "github.com/steve-care-software/steve/domain/languages/lexers/resources/multiples"
	resources_paths "github.com/steve-care-software/steve/domain/languages/lexers/resources/paths"
	resources_singles "github.com/steve-care-software/steve/domain/languages/lexers/resources/singles"
	"github.com/steve-care-software/steve/domain/languages/lexers/roots"
	"github.com/steve-care-software/steve/domain/languages/lexers/suites"
)

// NewBuilder creates a new builder root
func NewBuilder() Builder {
	resourcesPathBuilder := resources_paths.NewBuilder()
	resourcesContainerBuilder := resources_containers.NewBuilder()
	referencesAdapter := references.NewAdapter()
	patternsAdapter := patterns.NewAdapter()
	rulesAdapter := rules.NewAdapter()
	tokensAdapterBuilder := tokens.NewAdapterBuilder()
	pathAdapterBuilder := asts_paths.NewAdapterBuilder()
	suiteApplicationBuilder := suites.NewApplicationBuilder()
	astAdapterBuilder := asts.NewAdapterBuilder()
	astPathBuilder := asts_paths.NewBuilder()
	astPathElementBuilder := asts_paths.NewElementBuilder()
	rootAdapter := roots.NewAdapter()
	return createBuilder(
		resourcesPathBuilder,
		resourcesContainerBuilder,
		referencesAdapter,
		patternsAdapter,
		rulesAdapter,
		tokensAdapterBuilder,
		pathAdapterBuilder,
		suiteApplicationBuilder,
		astAdapterBuilder,
		astPathBuilder,
		astPathElementBuilder,
		rootAdapter,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithBasePaths(paths paths.Paths) Builder
	WithReplacementPaths(replacementPaths map[string]asts_paths.Path) Builder
	Now() (Application, error)
}

// Application represents a lexer application
type Application interface {
	Root(rootResource resources_singles.Single, content contents.Content) (roots.Root, mistakes.Mistake, error)
	AST(rootResource resources_singles.Single, content contents.Content) (asts.AST, error)
	Path(resource resources_paths.Path) (asts_paths.Path, error)
	Tokens(resource resources_containers.Container) (tokens.Tokens, error)
	Tests(resource resources_multiples.Multiple) ([]suites.Suite, error)
	References(token string) (references.References, error)
}
