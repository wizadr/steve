package patterns

type discovery struct {
	index     uint
	content   string
	container Container
}

func createDiscovery(
	index uint,
	content string,
	container Container,
) Discovery {
	out := discovery{
		index:     index,
		content:   content,
		container: container,
	}

	return &out
}

// Index returns the index
func (obj *discovery) Index() uint {
	return obj.index
}

// Content returns the content
func (obj *discovery) Content() string {
	return obj.content
}

// Container returns the container
func (obj *discovery) Container() Container {
	return obj.container
}
