package lexers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts"
	asts_paths "github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/references"
	resources_containers "github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"
	resources_paths "github.com/steve-care-software/steve/domain/languages/lexers/resources/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/roots"
	"github.com/steve-care-software/steve/domain/languages/lexers/suites"
)

type builder struct {
	resourcesPathBuilder      resources_paths.Builder
	resourcesContainerBuilder resources_containers.Builder
	referencesAdapter         references.Adapter
	patternsAdapter           patterns.Adapter
	rulesAdapter              rules.Adapter
	tokensAdapterBuilder      tokens.AdapterBuilder
	pathAdapterBuilder        asts_paths.AdapterBuilder
	suiteApplicationBuilder   suites.ApplicationBuilder
	astAdapterBuilder         asts.AdapterBuilder
	astPathBuilder            asts_paths.Builder
	astPathElementBuilder     asts_paths.ElementBuilder
	rootAdapter               roots.Adapter
	basePath                  paths.Paths
	replacementPaths          map[string]asts_paths.Path
}

func createBuilder(
	resourcesPathBuilder resources_paths.Builder,
	resourcesContainerBuilder resources_containers.Builder,
	referencesAdapter references.Adapter,
	patternsAdapter patterns.Adapter,
	rulesAdapter rules.Adapter,
	tokensAdapterBuilder tokens.AdapterBuilder,
	pathAdapterBuilder asts_paths.AdapterBuilder,
	suiteApplicationBuilder suites.ApplicationBuilder,
	astAdapterBuilder asts.AdapterBuilder,
	astPathBuilder asts_paths.Builder,
	astPathElementBuilder asts_paths.ElementBuilder,
	rootAdapter roots.Adapter,
) Builder {
	out := builder{
		resourcesPathBuilder:      resourcesPathBuilder,
		resourcesContainerBuilder: resourcesContainerBuilder,
		referencesAdapter:         referencesAdapter,
		patternsAdapter:           patternsAdapter,
		rulesAdapter:              rulesAdapter,
		tokensAdapterBuilder:      tokensAdapterBuilder,
		pathAdapterBuilder:        pathAdapterBuilder,
		suiteApplicationBuilder:   suiteApplicationBuilder,
		astAdapterBuilder:         astAdapterBuilder,
		astPathBuilder:            astPathBuilder,
		astPathElementBuilder:     astPathElementBuilder,
		rootAdapter:               rootAdapter,
		basePath:                  nil,
		replacementPaths:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.resourcesPathBuilder,
		app.resourcesContainerBuilder,
		app.referencesAdapter,
		app.patternsAdapter,
		app.rulesAdapter,
		app.tokensAdapterBuilder,
		app.pathAdapterBuilder,
		app.suiteApplicationBuilder,
		app.astAdapterBuilder,
		app.astPathBuilder,
		app.astPathElementBuilder,
		app.rootAdapter,
	)
}

// WithBasePaths adds a base textpaths to the builder
func (app *builder) WithBasePaths(paths paths.Paths) Builder {
	app.basePath = paths
	return app
}

// WithReplacementPaths add replacement paths to the builder
func (app *builder) WithReplacementPaths(replacementPaths map[string]asts_paths.Path) Builder {
	app.replacementPaths = replacementPaths
	return app
}

// Now builds a new Application root
func (app *builder) Now() (Application, error) {
	if app.basePath == nil {
		return nil, errors.New("the base Paths is mandatory in order to build an Application root")
	}

	if app.replacementPaths == nil {
		app.replacementPaths = map[string]asts_paths.Path{}
	}

	return createApplication(
		app.resourcesPathBuilder,
		app.resourcesContainerBuilder,
		app.referencesAdapter,
		app.patternsAdapter,
		app.rulesAdapter,
		app.tokensAdapterBuilder,
		app.pathAdapterBuilder,
		app.suiteApplicationBuilder,
		app.astAdapterBuilder,
		app.astPathBuilder,
		app.astPathElementBuilder,
		app.rootAdapter,
		app,
		app.basePath,
		app.replacementPaths,
	), nil
}
