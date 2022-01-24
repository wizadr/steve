package patterns

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type adapter struct {
	builder               Builder
	patternBuilder        PatternBuilder
	choiceBuilder         ChoiceBuilder
	serieBuilder          SerieBuilder
	groupsBuilder         GroupsBuilder
	groupBuilder          GroupBuilder
	resultBuilder         ResultBuilder
	nextBuilder           NextBuilder
	discoveriesBuilder    DiscoveriesBuilder
	discoveryBuilder      DiscoveryBuilder
	containerBuilder      ContainerBuilder
	groupNamePattern      string
	serieNamePattern      string
	reversePrefix         string
	patternNamePattern    string
	delimiter             string
	prefix                string
	suffix                string
	stringDelimiter       string
	whiteSpacePattern     string
	anythingExcept        string
	escapeReplacementCode string
}

func createAdapter(
	builder Builder,
	patternBuilder PatternBuilder,
	choiceBuilder ChoiceBuilder,
	serieBuilder SerieBuilder,
	groupsBuilder GroupsBuilder,
	groupBuilder GroupBuilder,
	resultBuilder ResultBuilder,
	nextBuilder NextBuilder,
	discoveriesBuilder DiscoveriesBuilder,
	discoveryBuilder DiscoveryBuilder,
	containerBuilder ContainerBuilder,
	groupNamePattern string,
	serieNamePattern string,
	reversePrefix string,
	patternNamePattern string,
	delimiter string,
	prefix string,
	suffix string,
	stringDelimiter string,
	whiteSpacePattern string,
	anythingExcept string,
	escapeReplacementCode string,
) Adapter {
	out := adapter{
		builder:               builder,
		patternBuilder:        patternBuilder,
		choiceBuilder:         choiceBuilder,
		serieBuilder:          serieBuilder,
		groupsBuilder:         groupsBuilder,
		groupBuilder:          groupBuilder,
		resultBuilder:         resultBuilder,
		nextBuilder:           nextBuilder,
		discoveriesBuilder:    discoveriesBuilder,
		discoveryBuilder:      discoveryBuilder,
		containerBuilder:      containerBuilder,
		groupNamePattern:      groupNamePattern,
		serieNamePattern:      serieNamePattern,
		reversePrefix:         reversePrefix,
		patternNamePattern:    patternNamePattern,
		delimiter:             delimiter,
		prefix:                prefix,
		suffix:                suffix,
		stringDelimiter:       stringDelimiter,
		whiteSpacePattern:     whiteSpacePattern,
		anythingExcept:        anythingExcept,
		escapeReplacementCode: escapeReplacementCode,
	}

	return &out
}

// ToPatterns converts content to a Patterns instance
func (app *adapter) ToPatterns(content string) (Patterns, error) {
	groups, err := app.findGroups(content)
	if err != nil {
		return nil, err
	}

	series, err := app.findSeries(groups, content)
	if err != nil {
		return nil, err
	}

	return app.findPatterns(series, content)
}

// FromGroupToResult converts a group to result
func (app *adapter) FromGroupToResult(group Group, content string) (Result, error) {
	amount := len(content)
	for i := 0; i < amount; i++ {
		subContent := content[i:]
		matchedAmount, err := app.matchesGroup(group, subContent, false)
		if err != nil {
			return nil, err
		}

		if matchedAmount == 0 {
			continue
		}

		container, err := app.containerBuilder.Create().WithGroup(group).Now()
		if err != nil {
			return nil, err
		}

		discovery, err := app.discoveryBuilder.Create().WithContainer(container).WithContent(content[i : i+int(matchedAmount)]).WithIndex(uint(i)).Now()
		if err != nil {
			return nil, err
		}

		discoveries, err := app.discoveriesBuilder.Create().WithList([]Discovery{
			discovery,
		}).Now()
		if err != nil {
			return nil, err
		}

		builder := app.resultBuilder.Create().WithDiscoveries(discoveries).WithInput(content)
		if matchedAmount < uint(len(content)) {
			builder.WithRemaining(content[matchedAmount:])
		}

		ins, err := builder.Now()
		if err != nil {
			return nil, err
		}

		return ins, nil
	}

	str := fmt.Sprintf("the group (name: %s) could not be any match in the given content: '%s'", group.Name(), content)
	return nil, errors.New(str)
}

// FromPatternToResult converts patterns with content to a Result instance
func (app *adapter) FromPatternToResult(pattern Pattern, content string) (Result, error) {
	amount := len(content)
	for i := 0; i < amount; i++ {
		subContent := content[i:]
		res, err := app.matchPattern(pattern, subContent, uint(i))
		if err != nil {
			return nil, err
		}

		if res == nil {
			continue
		}

		return res, nil
	}

	return nil, errors.New("the pattern could not find any match in the given content")
}

// FromPatternsToResult converts patterns + content to a Result instance
func (app *adapter) FromPatternsToResult(patterns Patterns, content string) (Result, error) {
	index := int(^uint(0) >> 1)
	var res Result

	list := patterns.All()
	for _, oneElement := range list {
		ins, err := app.FromPatternToResult(oneElement, content)
		if err != nil {
			continue
		}

		disIndex := int(ins.Discoveries().Index())
		if disIndex < index {
			index = disIndex
			res = ins
		}
	}

	if res == nil {
		return nil, errors.New("the patterns could not find any match in the given content")
	}

	return res, nil
}

func (app *adapter) matchPattern(pattern Pattern, content string, nextIndex uint) (Result, error) {
	var next Next
	remaining := content
	discoveriesList := []Discovery{}
	choices := pattern.Content().List()
	for idx, oneChoice := range choices {
		currentIndex := uint(0)
		if idx <= 0 {
			currentIndex = nextIndex
		}

		isMatch, matchedDiscoveries, nextGroup, remainingContent, err := app.matchChoice(oneChoice, remaining, currentIndex)
		if err != nil {
			return nil, err
		}

		remaining = remainingContent
		if !isMatch && nextGroup != nil {
			nextBuilder := app.nextBuilder.Create().WithGroup(nextGroup)
			if oneChoice.IsReverse() {
				nextBuilder.IsReverse()
			}

			nxt, err := nextBuilder.Now()
			if err != nil {
				return nil, err
			}

			next = nxt
			break
		}

		if !isMatch {
			break
		}

		discoveriesList = append(discoveriesList, matchedDiscoveries...)
	}

	if len(discoveriesList) <= 0 {
		return nil, nil
	}

	discoveries, err := app.discoveriesBuilder.Create().WithList(discoveriesList).Now()
	if err != nil {
		return nil, err
	}

	builder := app.resultBuilder.Create().WithDiscoveries(discoveries).WithInput(content)
	if next != nil {
		builder.WithNext(next)
	}

	if remaining != "" {
		builder.WithRemaining(remaining)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (app *adapter) matchChoice(choice Choice, content string, index uint) (bool, []Discovery, Group, string, error) {
	serie := choice.Serie()
	if choice.IsReverse() {
		return app.matchReversedSerie(serie, content, index)
	}

	return app.matchSerie(serie, content, index)
}

func (app *adapter) matchSerie(serie Serie, content string, index uint) (bool, []Discovery, Group, string, error) {
	matchedDiscoveriesList := []Discovery{}
	remaining := content
	groups := serie.Content().List()
	for _, oneGroup := range groups {
		if remaining == "" {
			return false, matchedDiscoveriesList, oneGroup, "", nil
		}

		amount, err := app.matchesGroup(oneGroup, remaining, false)
		if err != nil {
			return false, nil, nil, "", err
		}

		if amount <= 0 {
			return false, matchedDiscoveriesList, oneGroup, remaining, nil
		}

		container, err := app.containerBuilder.Create().WithGroup(oneGroup).Now()
		if err != nil {
			return false, nil, nil, "", nil
		}

		discovered := remaining[:amount]
		discovery, err := app.discoveryBuilder.Create().WithContainer(container).WithContent(discovered).WithIndex(index).Now()
		if err != nil {
			return false, nil, nil, "", nil
		}

		matchedDiscoveriesList = append(matchedDiscoveriesList, discovery)
		remaining = remaining[amount:]
	}

	return true, matchedDiscoveriesList, nil, remaining, nil
}

func (app *adapter) matchReversedSerie(serie Serie, content string, index uint) (bool, []Discovery, Group, string, error) {
	groups := serie.Content().List()
	amount, err := app.matchesGroup(groups[0], content, true)
	if err != nil {
		return false, nil, nil, "", err
	}

	if amount <= 0 {
		return false, nil, nil, "", nil
	}

	container, err := app.containerBuilder.Create().WithReverse(serie).Now()
	if err != nil {
		return false, nil, nil, "", err
	}

	discovered := content[:amount]
	discovery, err := app.discoveryBuilder.Create().WithContainer(container).WithContent(discovered).WithIndex(index).Now()
	if err != nil {
		return false, nil, nil, "", err
	}

	return true, []Discovery{
		discovery,
	}, groups[0], content[amount:], nil
}

func (app *adapter) matchesGroup(group Group, content string, isReverse bool) (uint, error) {
	if !isReverse {
		prefixLength := group.Content().Contains(content)
		return prefixLength, nil
	}

	// we are in reverse:
	amount := uint(0)
	remaining := content
	for {
		if len(remaining) <= 0 {
			break
		}

		prefixLength := group.Content().Contains(remaining)
		isBreak := prefixLength <= 0
		if isBreak {
			amount++
			remaining = remaining[1:]
			continue
		}

		break
	}

	return amount, nil
}

func (app *adapter) findPatterns(series map[string]Serie, content string) (Patterns, error) {
	anythingExceptEnd := fmt.Sprintf(app.anythingExcept, app.suffix)
	patternStr := fmt.Sprintf(
		"(%s)%s?%s%s?(%s)%s?%s",
		app.patternNamePattern,
		app.whiteSpacePattern,
		app.prefix,
		app.whiteSpacePattern,
		anythingExceptEnd,
		app.whiteSpacePattern,
		app.suffix,
	)

	patterns := []Pattern{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		choicePatternStr := fmt.Sprintf("(%s)?%s", app.reversePrefix, app.serieNamePattern)
		choiceNamePattern := regexp.MustCompile(choicePatternStr)
		choiceNames := choiceNamePattern.FindAllString(oneMatch[2], -1)
		if len(choiceNames) <= 0 {
			continue
		}

		choices := []Choice{}
		for _, oneChoiceName := range choiceNames {
			isReverse := strings.HasPrefix(oneChoiceName, app.reversePrefix)
			if isReverse {
				oneChoiceName = oneChoiceName[1:]
			}

			if _, ok := series[oneChoiceName]; !ok {
				str := fmt.Sprintf("the serie (name: %s) referenced in pattern (name: %s) does not exists", oneChoiceName, oneMatch[1])
				return nil, errors.New(str)
			}

			choiceBuilder := app.choiceBuilder.Create()
			if ins, ok := series[oneChoiceName]; ok {
				choiceBuilder.WithSerie(ins)
			}

			if isReverse {
				choiceBuilder.IsReverse()
			}

			choice, err := choiceBuilder.Now()
			if err != nil {
				return nil, err
			}

			choices = append(choices, choice)
		}

		pattern, err := app.patternBuilder.Create().WithList(choices).WithName(oneMatch[1]).Now()
		if err != nil {
			return nil, err
		}

		patterns = append(patterns, pattern)
	}

	return app.builder.Create().WithList(patterns).Now()
}

func (app *adapter) findSeries(groups Groups, content string) (map[string]Serie, error) {
	anythingExceptEnd := fmt.Sprintf(app.anythingExcept, app.suffix)
	patternStr := fmt.Sprintf(
		"(%s)%s?%s%s?(%s)%s?%s",
		app.serieNamePattern,
		app.whiteSpacePattern,
		app.prefix,
		app.whiteSpacePattern,
		anythingExceptEnd,
		app.whiteSpacePattern,
		app.suffix,
	)

	out := map[string]Serie{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		groupNamePattern := regexp.MustCompile(app.groupNamePattern)
		groupNameMatches := groupNamePattern.FindAllString(oneMatch[2], -1)
		if len(groupNameMatches) <= 0 {
			continue
		}

		list := []Group{}
		for _, oneName := range groupNameMatches {
			ins, err := groups.Find(oneName)
			if err != nil {
				str := fmt.Sprintf("the group (name: %s) referenced in serie (name: %s) may not exists: %s", oneName, oneMatch[1], err.Error())
				return nil, errors.New(str)
			}

			list = append(list, ins)
		}

		serie, err := app.serieBuilder.Create().WithList(list).WithName(oneMatch[1]).Now()
		if err != nil {
			return nil, err
		}

		name := serie.Name()
		out[name] = serie
	}

	return out, nil
}

func (app *adapter) findGroups(content string) (Groups, error) {
	anythingExceptEnd := fmt.Sprintf(app.anythingExcept, app.suffix)
	patternStr := fmt.Sprintf(
		"(%s)%s?%s%s?%s(%s)%s%s?%s",
		app.groupNamePattern,
		app.whiteSpacePattern,
		app.prefix,
		app.whiteSpacePattern,
		app.stringDelimiter,
		anythingExceptEnd,
		app.stringDelimiter,
		app.whiteSpacePattern,
		app.suffix,
	)

	escaped := fmt.Sprintf("%s%s", "\\", app.stringDelimiter)
	content = strings.Replace(content, escaped, app.escapeReplacementCode, -1)

	list := []Group{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		characters := []string{}
		options := strings.Split(oneMatch[2], app.delimiter)
		for _, oneOption := range options {

			if oneOption == app.escapeReplacementCode {
				oneOption = app.stringDelimiter
			}

			characters = append(characters, oneOption)
		}

		group, err := app.groupBuilder.Create().WithName(oneMatch[1]).WithList(characters).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, group)
	}

	return app.groupsBuilder.Create().WithList(list).Now()
}
