package multiples

import "github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"

type multiple struct {
	container containers.Container
	channels  string
}

func createMultiple(
	container containers.Container,
) Multiple {
	return createMultipleInternally(container, "")
}

func createMultipleWithChannels(
	container containers.Container,
	channels string,
) Multiple {
	return createMultipleInternally(container, channels)
}

func createMultipleInternally(
	container containers.Container,
	channels string,
) Multiple {
	out := multiple{
		container: container,
		channels:  channels,
	}

	return &out
}

// Container returns the container
func (obj *multiple) Container() containers.Container {
	return obj.container
}

// HasChannels returns true if there is channels, false otherwise
func (obj *multiple) HasChannels() bool {
	return obj.channels != ""
}

// Channels returns the channels, if any
func (obj *multiple) Channels() string {
	return obj.channels
}
