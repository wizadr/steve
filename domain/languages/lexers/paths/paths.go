package paths

type paths struct {
	rules    string
	tokens   string
	channels string
}

func createPaths(
	rules string,
	tokens string,
) Paths {
	return createPathsInternally(rules, tokens, "")
}

func createPathsWithChanels(
	rules string,
	tokens string,
	channels string,
) Paths {
	return createPathsInternally(rules, tokens, channels)
}

func createPathsInternally(
	rules string,
	tokens string,
	channels string,
) Paths {
	out := paths{
		rules:    rules,
		tokens:   tokens,
		channels: channels,
	}

	return &out
}

// Rules returns the rules
func (obj *paths) Rules() string {
	return obj.rules
}

// Tokens returns the tokens
func (obj *paths) Tokens() string {
	return obj.tokens
}

// HasChannels returns true if there is channels, false otherwise
func (obj *paths) HasChannels() bool {
	return obj.channels != ""
}

// Channels returns the channels, if any
func (obj *paths) Channels() string {
	return obj.channels
}
