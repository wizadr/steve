package lexers

import (
	"io/ioutil"
	"path/filepath"

	"github.com/steve-care-software/steve/domain/languages/lexers/asts"
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/mistakes"
	asts_paths "github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/rules"
	"github.com/steve-care-software/steve/domain/languages/lexers/containers/tokens"
	"github.com/steve-care-software/steve/domain/languages/lexers/contents"
	"github.com/steve-care-software/steve/domain/languages/lexers/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/references"
	"github.com/steve-care-software/steve/domain/languages/lexers/references/links"
	resources_containers "github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"
	resources_multiples "github.com/steve-care-software/steve/domain/languages/lexers/resources/multiples"
	resources_paths "github.com/steve-care-software/steve/domain/languages/lexers/resources/paths"
	resources_singles "github.com/steve-care-software/steve/domain/languages/lexers/resources/singles"
	"github.com/steve-care-software/steve/domain/languages/lexers/roots"
	"github.com/steve-care-software/steve/domain/languages/lexers/suites"
)

type application struct {
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
	builder                   Builder
	basePath                  paths.Paths
	replacementPaths          map[string]asts_paths.Path
}

func createApplication(
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
	builder Builder,
	basePath paths.Paths,
	replacementPaths map[string]asts_paths.Path,
) Application {
	out := application{
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
		builder:                   builder,
		basePath:                  basePath,
		replacementPaths:          replacementPaths,
	}

	return &out
}

// Root executes the root single resource and the content to produce the Root root
func (app *application) Root(rootResource resources_singles.Single, content contents.Content) (roots.Root, mistakes.Mistake, error) {
	ast, err := app.AST(rootResource, content)
	if err != nil {
		return nil, nil, err
	}

	if ast.IsMistake() {
		mistake := ast.Mistake()
		return nil, mistake, nil
	}

	rootToken := ast.Success()
	ins, err := app.rootAdapter.ToRoot(rootToken)
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}

// AST executes the root single resource and the content to produce the AST root
func (app *application) AST(rootResource resources_singles.Single, content contents.Content) (asts.AST, error) {
	multiple := rootResource.Multiple()
	container := multiple.Container()
	patternsScript := container.Patterns()
	rulesScript := container.Rules()
	tokensScript := container.Tokens()
	channelsScript := ""
	if multiple.HasChannels() {
		channelsScript = multiple.Channels()
	}

	rootTokenName := rootResource.RootToken()
	rootToken, tokens, channels, err := app.createRootToken(rootTokenName, patternsScript, rulesScript, tokensScript, channelsScript)
	if err != nil {
		return nil, err
	}

	astAdapterBuilder := app.astAdapterBuilder.Create().WithRoot(rootToken).WithTokens(tokens)
	if channels != nil {
		astAdapterBuilder.WithChannels(channels)
	}

	if content.IsPrefixLegal() {
		astAdapterBuilder.CanContainPrefix()
	}

	if content.IsSuffixLegal() {
		astAdapterBuilder.CanContainSuffix()
	}

	references, err := app.References(tokensScript)
	if err != nil {
		return nil, err
	}

	notFoundFunc, replacementFunc, err := app.createNotFoundAndReplacementFuncs(container, references, app.replacementPaths, false)
	if err != nil {
		return nil, err
	}

	if notFoundFunc != nil {
		astAdapterBuilder.WithTokenNotFoundFunc(notFoundFunc)
	}

	if replacementFunc != nil {
		astAdapterBuilder.WithTokenReplacementFunc(replacementFunc)
	}

	astAdapter, err := astAdapterBuilder.Now()
	if err != nil {
		return nil, err
	}

	cnt := content.Content()
	return astAdapter.ToAST(cnt)
}

// Path executes the path resource and the content to produce the Path AST root
func (app *application) Path(resource resources_paths.Path) (asts_paths.Path, error) {
	return app.path(resource, app.replacementPaths, false)
}

func (app *application) path(resource resources_paths.Path, replacementPaths map[string]asts_paths.Path, skipReplacement bool) (asts_paths.Path, error) {
	container := resource.Container()
	patternsScript := container.Patterns()
	rulesScript := container.Rules()
	tokensScript := container.Tokens()
	rootTokenName := resource.RootToken()
	rootToken, tokens, _, err := app.createRootToken(rootTokenName, patternsScript, rulesScript, tokensScript, "")
	if err != nil {
		return nil, err
	}

	references, err := app.References(tokensScript)
	if err != nil {
		return nil, err
	}

	pathAdapterBuilder := app.pathAdapterBuilder.Create().WithTokens(tokens)
	notFoundFunc, replacementFunc, err := app.createNotFoundAndReplacementFuncs(container, references, replacementPaths, skipReplacement)
	if err != nil {
		return nil, err
	}

	if notFoundFunc != nil {
		pathAdapterBuilder.WithTokenNotFoundFunc(notFoundFunc)
	}

	if replacementFunc != nil {
		pathAdapterBuilder.WithTokenReplacementFunc(replacementFunc)
	}

	pathAdapter, err := pathAdapterBuilder.Now()
	if err != nil {
		return nil, err
	}

	return pathAdapter.ToPath(rootToken)
}

// Tokens executes the container resource and the content to produce the Tokens root
func (app *application) Tokens(resource resources_containers.Container) (tokens.Tokens, error) {
	patternsContent := resource.Patterns()
	patterns, err := app.patternsAdapter.ToPatterns(patternsContent)
	if err != nil {
		return nil, err
	}

	rulesContent := resource.Rules()
	rules, err := app.rulesAdapter.Rules(rulesContent, patterns)
	if err != nil {
		return nil, err
	}

	tokenAdapter, err := app.tokensAdapterBuilder.Create().WithRules(rules).Now()
	if err != nil {
		return nil, err
	}

	tokensScript := resource.Tokens()
	return tokenAdapter.ToTokens(tokensScript)
}

// Tests executes the multiple resource and the content to produce the Suite roots
func (app *application) Tests(resource resources_multiples.Multiple) ([]suites.Suite, error) {
	container := resource.Container()
	patternsScript := container.Patterns()
	rulesScript := container.Rules()
	tokensScript := container.Tokens()
	channelsScript := ""
	if resource.HasChannels() {
		channelsScript = resource.Channels()
	}

	tokens, channels, err := app.createTokens(patternsScript, rulesScript, tokensScript, channelsScript)
	if err != nil {
		return nil, err
	}

	references, err := app.References(tokensScript)
	if err != nil {
		return nil, err
	}

	suiteApplicationBuilder := app.suiteApplicationBuilder.Create().WithTokens(tokens)
	if channels != nil {
		suiteApplicationBuilder.WithChannels(channels)
	}

	notFoundFunc, replacementFunc, err := app.createNotFoundAndReplacementFuncs(container, references, app.replacementPaths, false)
	if err != nil {
		return nil, err
	}

	if notFoundFunc != nil {
		suiteApplicationBuilder.WithTokenNotFoundFunc(notFoundFunc)
	}

	if replacementFunc != nil {
		suiteApplicationBuilder.WithTokenReplacementFunc(replacementFunc)
	}

	suiteApplication, err := suiteApplicationBuilder.Now()
	if err != nil {
		return nil, err
	}

	return suiteApplication.Execute()
}

// References executes the multiple resource and the content to produce the References root
func (app *application) References(tokensContent string) (references.References, error) {
	references, err := app.referencesAdapter.ToReferences(tokensContent)
	if err != nil {
		return nil, err
	}

	return references, nil
}

func (app *application) createRootToken(rootTokenName string, patternsScript string, rulesScript string, tokensScript string, channelsScript string) (tokens.Token, tokens.Tokens, tokens.Tokens, error) {
	tokens, channels, err := app.createTokens(patternsScript, rulesScript, tokensScript, channelsScript)
	if err != nil {
		return nil, nil, nil, err
	}

	rootToken, err := tokens.Find(rootTokenName)
	if err != nil {
		return nil, nil, nil, err
	}

	return rootToken, tokens, channels, nil
}

func (app *application) createTokens(patternsScript string, rulesScript string, tokensScript string, channelsScript string) (tokens.Tokens, tokens.Tokens, error) {
	patterns, err := app.patternsAdapter.ToPatterns(patternsScript)
	if err != nil {
		return nil, nil, err
	}

	rules, err := app.rulesAdapter.Rules(rulesScript, patterns)
	if err != nil {
		return nil, nil, err
	}

	tokenAdapter, err := app.tokensAdapterBuilder.Create().WithRules(rules).Now()
	if err != nil {
		return nil, nil, err
	}

	tokens, err := tokenAdapter.ToTokens(tokensScript)
	if err != nil {
		return nil, nil, err
	}

	if channelsScript != "" {
		channels, err := tokenAdapter.ToTokens(channelsScript)
		if err != nil {
			return nil, nil, err
		}

		return tokens, channels, nil
	}

	return tokens, nil, nil
}

func (app *application) createNotFoundAndReplacementFuncs(
	resContainer resources_containers.Container,
	references references.References,
	replacementPaths map[string]asts_paths.Path,
	skipReplacement bool,
) (asts_paths.FetchTokenNotFoundFn, asts_paths.FetchElementReplacementFn, error) {
	if !skipReplacement {
		replacement, err := app.fetchReplacementPaths(references, resContainer, replacementPaths)
		if err != nil {
			return nil, nil, err
		}

		replacementPaths = replacement
	}

	notFoundPaths, err := app.fetchNotFoundPaths(references, resContainer, replacementPaths)
	if err != nil {
		return nil, nil, err
	}

	var notFoundFn asts_paths.FetchTokenNotFoundFn
	if len(notFoundPaths) > 0 {
		notFoundFn = app.returnFetchTokenNotFoundFn(notFoundPaths)
	}

	var replacementFn asts_paths.FetchElementReplacementFn
	if len(replacementPaths) > 0 {
		replacementFn = app.returnFetchElementReplacementFn(replacementPaths)
	}

	return notFoundFn, replacementFn, nil
}

func (app *application) fetchReplacementPaths(
	references references.References,
	resContainer resources_containers.Container,
	replacementPaths map[string]asts_paths.Path,
) (map[string]asts_paths.Path, error) {
	if !references.HasReplacements() {
		return replacementPaths, nil
	}

	repList := references.Replacements().All()
	for _, oneReplacement := range repList {
		localTokenName := oneReplacement.LocalToken()
		resPath, err := app.resourcesPathBuilder.Create().WithRootToken(localTokenName).WithContainer(resContainer).Now()
		if err != nil {
			return nil, err
		}

		refPath, err := app.path(resPath, replacementPaths, true)
		if err != nil {
			return nil, err
		}

		refTokenName := oneReplacement.ReferenceToken()
		replacementPaths[refTokenName] = refPath
	}

	return replacementPaths, nil
}

func (app *application) fetchNotFoundPaths(
	references references.References,
	resContainer resources_containers.Container,
	replacementPaths map[string]asts_paths.Path,
) (map[string]asts_paths.Path, error) {
	notFoundPaths := map[string]asts_paths.Path{}
	if !references.HasReferences() {
		return notFoundPaths, nil
	}

	refList := references.References().All()
	for _, oneReference := range refList {
		path, err := app.linkToPathFromFile(oneReference, replacementPaths)
		if err != nil {
			return nil, err
		}

		localTokenName := oneReference.LocalToken()
		notFoundPaths[localTokenName] = path
	}

	return notFoundPaths, nil
}

func (app *application) linkToPathFromFile(
	reference links.Link,
	replacements map[string]asts_paths.Path,
) (asts_paths.Path, error) {
	applicationBuilder := app.builder.Create()
	if len(replacements) > 0 {
		applicationBuilder.WithReplacementPaths(replacements)
	}

	application, err := applicationBuilder.WithBasePaths(app.basePath).Now()
	if err != nil {
		return nil, err
	}

	includePaths := reference.Include().Paths()
	rulesPath := filepath.Join(app.basePath.Rules(), includePaths.Rules())
	rulesContent, err := ioutil.ReadFile(rulesPath)
	if err != nil {
		return nil, err
	}

	tokenPath := filepath.Join(app.basePath.Tokens(), includePaths.Tokens())
	tokenContent, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}

	resContainer, err := app.resourcesContainerBuilder.Create().WithRules(string(rulesContent)).WithTokens(string(tokenContent)).Now()
	if err != nil {
		return nil, err
	}

	rootToken := reference.ReferenceToken()
	resPath, err := app.resourcesPathBuilder.Create().WithRootToken(rootToken).WithContainer(resContainer).Now()
	if err != nil {
		return nil, err
	}

	refPath, err := application.Path(resPath)
	if err != nil {
		return nil, err
	}

	return app.replacePathElementName(refPath, reference)
}

func (app *application) replacePathElementName(
	refPath asts_paths.Path,
	reference links.Link,
) (asts_paths.Path, error) {
	localTokenName := reference.LocalToken()
	pathElement := refPath.Element()
	pathElementMustLines := pathElement.Must()
	newElementBuilder := app.astPathElementBuilder.Create().WithName(localTokenName).WithMust(pathElementMustLines)
	if pathElement.HasNot() {
		pathElementNotLines := pathElement.Not()
		newElementBuilder.WithNot(pathElementNotLines)
	}

	newElement, err := newElementBuilder.Now()
	if err != nil {
		return nil, err
	}

	dep := refPath.Dependencies()
	return app.astPathBuilder.Create().WithElement(newElement).WithDependencies(dep).Now()
}

func (app *application) returnFetchTokenNotFoundFn(
	mp map[string]asts_paths.Path,
) asts_paths.FetchTokenNotFoundFn {
	return func(name string) asts_paths.Path {
		if path, ok := mp[name]; ok {
			return path
		}

		return nil
	}
}

func (app *application) returnFetchElementReplacementFn(
	mp map[string]asts_paths.Path,
) asts_paths.FetchElementReplacementFn {
	return func(name string) (asts_paths.Element, []asts_paths.Element) {
		if path, ok := mp[name]; ok {
			return path.Element(), path.Dependencies().All()
		}

		return nil, nil
	}
}
