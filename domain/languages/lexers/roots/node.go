package roots

type node struct {
	name       string
	containers Containers
}

func createNode(
	name string,
	containers Containers,
) Node {
	out := node{
		name:       name,
		containers: containers,
	}

	return &out
}

// Content returns the content
func (obj *node) Content() string {
	return obj.containers.Content()
}

// Name returns the name
func (obj *node) Name() string {
	return obj.name
}

// Containers returns the containers
func (obj *node) Containers() Containers {
	return obj.containers
}

// Scan scans the node
func (obj *node) Scan(name string) (Root, Node, Container, string, error) {
	if obj.name == name {
		return nil, obj, nil, "", nil
	}

	return obj.containers.Scan(name)
}
