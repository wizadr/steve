package tokens

import "errors"

type blockMatchBuilder struct {
	container     Match
	block         BlockMatch
	nextElement   NextElement
	channelPrefix LineMatch
}

func createBlockMatchBuilder() BlockMatchBuilder {
	out := blockMatchBuilder{
		container:     nil,
		block:         nil,
		nextElement:   nil,
		channelPrefix: nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockMatchBuilder) Create() BlockMatchBuilder {
	return createBlockMatchBuilder()
}

// WithContainer adds a container to the builder
func (app *blockMatchBuilder) WithContainer(container Match) BlockMatchBuilder {
	app.container = container
	return app
}

// WithBlock adds a block to the builder
func (app *blockMatchBuilder) WithBlock(block BlockMatch) BlockMatchBuilder {
	app.block = block
	return app
}

// WithNextElement adds a nextElement to the builder
func (app *blockMatchBuilder) WithNextElement(nextElement NextElement) BlockMatchBuilder {
	app.nextElement = nextElement
	return app
}

// WithChannelPrefix adds a channelPrefix to the builder
func (app *blockMatchBuilder) WithChannelPrefix(channelPrefix LineMatch) BlockMatchBuilder {
	app.channelPrefix = channelPrefix
	return app
}

// Now builds a new BlockMatch instance
func (app *blockMatchBuilder) Now() (BlockMatch, error) {
	if app.container != nil {
		content := createBlockMatchContentWithContainerMatch(app.container)
		if app.channelPrefix != nil {
			return createBlockMatchWithChannel(content, app.channelPrefix), nil
		}

		return createBlockMatch(content), nil
	}

	if app.block != nil {
		content := createBlockMatchContentWithBlockMatch(app.block)
		if app.channelPrefix != nil {
			return createBlockMatchWithChannel(content, app.channelPrefix), nil
		}

		return createBlockMatch(content), nil
	}

	if app.nextElement != nil {
		content := createBlockMatchContentWithNextElement(app.nextElement)
		if app.channelPrefix != nil {
			return createBlockMatchWithChannel(content, app.channelPrefix), nil
		}

		return createBlockMatch(content), nil
	}

	return nil, errors.New("the BlockMatch is invalid")
}
