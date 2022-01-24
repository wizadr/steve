package tokens

type instruction struct {
	isChannelSwitch bool
	element         Element
}

func createInstructionWithChannelSwitch() Instruction {
	return createInstructionInternally(true, nil)
}

func createInstructionWithElement(element Element) Instruction {
	return createInstructionInternally(false, element)
}

func createInstructionInternally(
	isChannelSwitch bool,
	element Element,
) Instruction {
	out := instruction{
		isChannelSwitch: isChannelSwitch,
		element:         element,
	}

	return &out
}

// IsChannelSwitch returns true if there is a channel switch, false otherwise
func (obj *instruction) IsChannelSwitch() bool {
	return obj.isChannelSwitch
}

// IsElement returns true if there is an elemnt, false otherwise
func (obj *instruction) IsElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *instruction) Element() Element {
	return obj.element
}
