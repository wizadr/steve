package roots

type container struct {
	content string
	name    string
	root    Root
}

func createContainer(
	content string,
	name string,
) Container {
	return createContainerInternally(content, name, nil)
}

func createContainerWithRoot(
	content string,
	name string,
	root Root,
) Container {
	return createContainerInternally(content, name, root)
}

func createContainerInternally(
	content string,
	name string,
	root Root,
) Container {
	out := container{
		content: content,
		name:    name,
		root:    root,
	}

	return &out
}

// Scan scans the container
func (obj *container) Scan(name string) (Root, Node, Container, string, error) {
	if obj.Name() == name {
		return nil, nil, obj, "", nil
	}

	if !obj.HasRoot() {
		return nil, nil, nil, "", nil
	}

	return obj.root.Nodes().Scan(name)
}

// Content returns the content
func (obj *container) Content() string {
	return obj.content
}

// Name returns the container name
func (obj *container) Name() string {
	return obj.name
}

// HasRoot returns true if there is a root, false otherwise
func (obj *container) HasRoot() bool {
	return obj.root != nil
}

// Root returns the root, if any
func (obj *container) Root() Root {
	return obj.root
}
