package tokens

import "strings"

type blockMatch struct {
	content       BlockMatchContent
	channelPrefix LineMatch
}

func createBlockMatch(
	content BlockMatchContent,
) BlockMatch {
	return createBlockMatchInternally(content, nil)
}

func createBlockMatchWithChannel(
	content BlockMatchContent,
	channelPrefix LineMatch,
) BlockMatch {
	return createBlockMatchInternally(content, channelPrefix)
}

func createBlockMatchInternally(
	content BlockMatchContent,
	channelPrefix LineMatch,
) BlockMatch {
	out := blockMatch{
		content:       content,
		channelPrefix: channelPrefix,
	}

	return &out
}

// Input returns the input
func (obj *blockMatch) Input() string {
	prefix := ""
	if obj.HasChannelPrefix() {
		prefix = obj.ChannelPrefix().Matches().Input()
	}

	content := obj.Content()
	if content.IsContainer() {
		return strings.Join([]string{
			prefix,
			content.Container().Input(),
		}, "")
	}

	return strings.Join([]string{
		prefix,
		content.Block().Input(),
	}, "")
}

// IsValid returns true if the blockMatch is valid, false otherwise
func (obj *blockMatch) IsValid() bool {
	content := obj.Content()
	if content.IsContainer() {
		return content.Container().IsValid()
	}

	if content.IsBlock() {
		return content.Block().IsValid()
	}

	// the match is invalid if the BlockMatch contains a NextElement instance:
	return false
}

// IsExact returns true if the blockMatch is exact, false otherwise
func (obj *blockMatch) IsExact() bool {
	content := obj.Content()
	if content.IsContainer() {
		return content.Container().IsExact()
	}

	if content.IsBlock() {
		return content.Block().IsExact()
	}

	// the match is not exact if the BlockMatch contains a NextElement instance:
	return false
}

// Discoveries returns the matched discoveries
func (obj *blockMatch) Discoveries() string {
	prefix := ""
	if obj.HasChannelPrefix() {
		prefix = obj.ChannelPrefix().Matches().Discoveries()
	}

	content := obj.Content()
	if content.IsContainer() {
		return strings.Join([]string{
			prefix,
			content.Container().Discoveries(),
		}, "")
	}

	if content.IsBlock() {
		return strings.Join([]string{
			prefix,
			content.Block().Discoveries(),
		}, "")
	}

	return prefix
}

// Content returns the content
func (obj *blockMatch) Content() BlockMatchContent {
	return obj.content
}

// HasChannelPrefix returns true if there is a channelPrefix, false otherwise
func (obj *blockMatch) HasChannelPrefix() bool {
	return obj.channelPrefix != nil
}

// ChannelPrefix returns the channel prefix, if any
func (obj *blockMatch) ChannelPrefix() LineMatch {
	return obj.channelPrefix
}
