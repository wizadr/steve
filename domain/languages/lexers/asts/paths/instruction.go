package paths

type instruction struct {
	isChannelSwitch bool
	container       Container
}

func createInstructionWithChannelSwitch() Instruction {
	return createInstructionInternally(true, nil)
}

func createInstructionWithContainer(container Container) Instruction {
	return createInstructionInternally(false, container)
}

func createInstructionInternally(
	isChannelSwitch bool,
	container Container,
) Instruction {
	out := instruction{
		isChannelSwitch: isChannelSwitch,
		container:       container,
	}

	return &out
}

// IsChannelSwitch returns true if there is a channel switch, false otherwise
func (obj *instruction) IsChannelSwitch() bool {
	return obj.isChannelSwitch
}

// IsContainer returns true if there is a container, false otherwise
func (obj *instruction) IsContainer() bool {
	return obj.container != nil
}

// Container retruns the container, if any
func (obj *instruction) Container() Container {
	return obj.container
}
