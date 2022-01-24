package references

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/steve-care-software/steve/domain/languages/lexers/paths"
	"github.com/steve-care-software/steve/domain/languages/lexers/references/links"
)

type adapter struct {
	includeBuilder        links.IncludeBuilder
	pathsBuilder          paths.Builder
	linkBuilder           links.LinkBuilder
	linksBuilder          links.Builder
	builder               Builder
	includeKeyword        string
	quoteChar             string
	dotChar               string
	referenceKeyword      string
	replacementKeyword    string
	tokenPattern          string
	anythingExceptPattern string
	spacePattern          string
	pathSeparator         string
}

func createAdapter(
	includeBuilder links.IncludeBuilder,
	pathsBuilder paths.Builder,
	linkBuilder links.LinkBuilder,
	linksBuilder links.Builder,
	builder Builder,
	includeKeyword string,
	quoteChar string,
	dotChar string,
	referenceKeyword string,
	replacementKeyword string,
	tokenPattern string,
	anythingExceptPattern string,
	spacePattern string,
	pathSeparator string,
) Adapter {
	out := adapter{
		includeBuilder:        includeBuilder,
		pathsBuilder:          pathsBuilder,
		linkBuilder:           linkBuilder,
		linksBuilder:          linksBuilder,
		builder:               builder,
		includeKeyword:        includeKeyword,
		quoteChar:             quoteChar,
		dotChar:               dotChar,
		referenceKeyword:      referenceKeyword,
		replacementKeyword:    replacementKeyword,
		tokenPattern:          tokenPattern,
		anythingExceptPattern: anythingExceptPattern,
		spacePattern:          spacePattern,
		pathSeparator:         pathSeparator,
	}

	return &out
}

// ToReferences converts content to a References instance, if any
func (app *adapter) ToReferences(content string) (References, error) {
	includes, err := app.includes(content)
	if err != nil {
		return nil, err
	}

	refLinks, err := app.references(includes, content)
	if err != nil {
		return nil, err
	}

	replacements, err := app.replacements(includes, content)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithReferences(refLinks).
		WithReplacements(replacements).
		Now()
}

func (app *adapter) includes(content string) (map[string]links.Include, error) {
	anythingExceptQuote := fmt.Sprintf(app.anythingExceptPattern, app.quoteChar)
	patternStr := fmt.Sprintf(
		"%s?%s%s(%s)%s%s(%s)%s%s?",
		app.spacePattern,
		app.includeKeyword,
		app.spacePattern,
		app.tokenPattern,
		app.spacePattern,
		app.quoteChar,
		anythingExceptQuote,
		app.quoteChar,
		app.spacePattern,
	)

	includes := map[string]links.Include{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		pathStrings := strings.Split(oneMatch[2], app.pathSeparator)
		amount := len(pathStrings)
		name := oneMatch[1]
		if amount > 3 {
			str := fmt.Sprintf("the include (name: %s) must contain {2,3} (rules, tokens, channels (optional)) paths, %d paths provided", name, amount)
			return nil, errors.New(str)
		}

		pathsBuilder := app.pathsBuilder.Create().WithRules(pathStrings[0]).WithTokens(pathStrings[1])
		if amount > 2 {
			pathsBuilder.WithChannels(pathStrings[2])
		}

		paths, err := pathsBuilder.Now()
		if err != nil {
			return nil, err
		}

		ins, err := app.includeBuilder.Create().WithName(name).WithPaths(paths).Now()
		if err != nil {
			return nil, err
		}

		keyname := ins.Name()
		includes[keyname] = ins
	}

	return includes, nil
}

func (app *adapter) references(includes map[string]links.Include, content string) (links.Links, error) {
	patternStr := fmt.Sprintf(
		"%s?%s%s(%s)%s(%s)%s(%s)%s",
		app.spacePattern,
		app.referenceKeyword,
		app.spacePattern,
		app.tokenPattern,
		app.spacePattern,
		app.tokenPattern,
		app.dotChar,
		app.tokenPattern,
		app.spacePattern,
	)

	links := []links.Link{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		builder := app.linkBuilder.Create()
		if include, ok := includes[oneMatch[2]]; ok {
			builder.WithInclude(include)
		}

		ins, err := builder.WithLocalToken(oneMatch[1]).WithReferenceToken(oneMatch[3]).Now()
		if err != nil {
			return nil, err
		}

		links = append(links, ins)
	}

	if len(links) <= 0 {
		return nil, nil
	}

	return app.linksBuilder.Create().WithLinks(links).Now()
}

func (app *adapter) replacements(includes map[string]links.Include, content string) (links.Links, error) {
	patternStr := fmt.Sprintf(
		"%s?%s%s(%s)%s(%s)%s(%s)%s",
		app.spacePattern,
		app.replacementKeyword,
		app.spacePattern,
		app.tokenPattern,
		app.dotChar,
		app.tokenPattern,
		app.spacePattern,
		app.tokenPattern,
		app.spacePattern,
	)

	links := []links.Link{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		builder := app.linkBuilder.Create()
		if include, ok := includes[oneMatch[1]]; ok {
			builder.WithInclude(include)
		}

		ins, err := builder.WithLocalToken(oneMatch[3]).WithReferenceToken(oneMatch[2]).Now()
		if err != nil {
			return nil, err
		}

		links = append(links, ins)
	}

	if len(links) <= 0 {
		return nil, nil
	}

	return app.linksBuilder.Create().WithLinks(links).Now()
}
