package roots

type root struct {
	content string
	name    string
	nodes   Nodes
}

func createRoot(
	content string,
	name string,
	nodes Nodes,
) Root {
	out := root{
		content: content,
		name:    name,
		nodes:   nodes,
	}

	return &out
}

// Content returns the content
func (obj *root) Content() string {
	return obj.content
}

// Name returns the name
func (obj *root) Name() string {
	return obj.name
}

// Nodes returns the nodes
func (obj *root) Nodes() Nodes {
	return obj.nodes
}

// Scan scans the root
func (obj *root) Scan(name string) (Root, Node, Container, string, error) {
	if obj.Name() == name {
		return obj, nil, nil, "", nil
	}

	return obj.nodes.Scan(name)
}
