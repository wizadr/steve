package tokens

type blockMatchContent struct {
	container   Match
	block       BlockMatch
	nextElement NextElement
}

func createBlockMatchContentWithContainerMatch(
	container Match,
) BlockMatchContent {
	return createBlockMatchContentInternally(container, nil, nil)
}

func createBlockMatchContentWithBlockMatch(
	block BlockMatch,
) BlockMatchContent {
	return createBlockMatchContentInternally(nil, block, nil)
}

func createBlockMatchContentWithNextElement(
	nextElement NextElement,
) BlockMatchContent {
	return createBlockMatchContentInternally(nil, nil, nextElement)
}

func createBlockMatchContentInternally(
	container Match,
	block BlockMatch,
	nextElement NextElement,
) BlockMatchContent {
	out := blockMatchContent{
		container:   container,
		block:       block,
		nextElement: nextElement,
	}

	return &out
}

// IsContainer returns true if there is container matches, false otherwise
func (obj *blockMatchContent) IsContainer() bool {
	return obj.container != nil
}

// Container returns the container matches, if any
func (obj *blockMatchContent) Container() Match {
	return obj.container
}

// IsBlock returns true if there is a BlockMatch,, false otherwise
func (obj *blockMatchContent) IsBlock() bool {
	return obj.block != nil
}

// Block returns the BlockMatch, if any
func (obj *blockMatchContent) Block() BlockMatch {
	return obj.block
}

// IsNextElement returns true if there is a next element, false otherwise
func (obj *blockMatchContent) IsNextElement() bool {
	return obj.nextElement != nil
}

// NextElement returns the next nextElement, if any
func (obj *blockMatchContent) NextElement() NextElement {
	return obj.nextElement
}
