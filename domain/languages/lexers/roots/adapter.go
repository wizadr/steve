package roots

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts"
)

type adapter struct {
	builder           Builder
	nodesBuilder      NodesBuilder
	nodeBuilder       NodeBuilder
	containersBuilder ContainersBuilder
	containerBuilder  ContainerBuilder
	tokenRoots        map[string]Root
}

func createAdapter(
	builder Builder,
	nodesBuilder NodesBuilder,
	nodeBuilder NodeBuilder,
	containersBuilder ContainersBuilder,
	containerBuilder ContainerBuilder,
) Adapter {
	out := adapter{
		builder:           builder,
		nodesBuilder:      nodesBuilder,
		nodeBuilder:       nodeBuilder,
		containersBuilder: containersBuilder,
		containerBuilder:  containerBuilder,
		tokenRoots:        nil,
	}

	return &out
}

func (app *adapter) init() {
	app.tokenRoots = map[string]Root{}
}

// ToRoot converts a root token to an instance
func (app *adapter) ToRoot(token asts.Token) (Root, error) {
	app.init()
	return app.toRoot(token)
}

func (app *adapter) toRoot(token asts.Token) (Root, error) {
	lineMatch := token.Match()
	if !lineMatch.HasMatches() {
		return nil, nil
	}

	// enter the sub tokens:
	matches := lineMatch.Matches()
	for _, oneMatch := range matches {
		if !oneMatch.IsToken() {
			continue
		}

		tokens := oneMatch.Token().Matches()
		for _, oneToken := range tokens {
			_, err := app.toRoot(oneToken)
			if err != nil {
				return nil, err
			}
		}
	}

	// create the containers map:
	containersMap := map[string][]Container{}
	for _, oneMatch := range matches {
		if oneMatch.IsToken() {
			tokenMatch := oneMatch.Token()
			tokenName := tokenMatch.Token()
			matchTokens := tokenMatch.Matches()
			for _, oneMatchToken := range matchTokens {
				root, err := app.toRootWithCache(oneMatchToken)
				if err != nil {
					return nil, err
				}

				content := oneMatchToken.Discovery()
				container, err := app.containerBuilder.Create().WithContent(content).WithName(tokenName).WithRoot(root).Now()
				if err != nil {
					return nil, err
				}

				containersMap[tokenName] = append(containersMap[tokenName], container)
			}

			continue
		}

		ruleMatch := oneMatch.Rule()
		ruleName := ruleMatch.Rule()
		content := ruleMatch.Result().Discoveries().Content()
		container, err := app.containerBuilder.Create().WithContent(content).WithName(ruleName).Now()
		if err != nil {
			return nil, err
		}

		containersMap[ruleName] = append(containersMap[ruleName], container)
	}

	nodesMap := map[string]Node{}
	for name, containerList := range containersMap {
		containers, err := app.containersBuilder.WithList(containerList).Now()
		if err != nil {
			return nil, err
		}

		node, err := app.nodeBuilder.Create().WithName(name).WithContainers(containers).Now()
		if err != nil {
			return nil, err
		}

		nodesMap[name] = node
	}

	nodes, err := app.nodesBuilder.Create().WithNodes(nodesMap).Now()
	if err != nil {
		return nil, err
	}

	name := token.Name()
	content := token.Discovery()
	ins, err := app.builder.Create().WithNodes(nodes).WithName(name).WithContent(content).Now()
	if err != nil {
		return nil, err
	}

	app.tokenRoots[content] = ins
	return ins, nil
}

func (app *adapter) toRootWithCache(token asts.Token) (Root, error) {
	content := token.Discovery()
	if ref, ok := app.tokenRoots[content]; ok {
		return ref, nil
	}

	return app.toRoot(token)
}
