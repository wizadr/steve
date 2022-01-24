package roots

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	builder := NewBuilder()
	nodesBuilder := NewNodesBuilder()
	nodeBuilder := NewNodeBuilder()
	containersBuilder := NewContainersBuilder()
	containerBuilder := NewContainerBuilder()
	return createAdapter(
		builder,
		nodesBuilder,
		nodeBuilder,
		containersBuilder,
		containerBuilder,
	)
}

// NewBuilder creates a new instance builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewNodesBuilder creates a new nodes builder
func NewNodesBuilder() NodesBuilder {
	return createNodesBuilder()
}

// NewNodeBuilder creates a new node builder
func NewNodeBuilder() NodeBuilder {
	return createNodeBuilder()
}

// NewContainersBuilder creates a new containers builder
func NewContainersBuilder() ContainersBuilder {
	return createContainersBuilder()
}

// NewContainerBuilder creates a new container builder
func NewContainerBuilder() ContainerBuilder {
	return createContainerBuilder()
}

// Adapter represents an instances adapter
type Adapter interface {
	ToRoot(rootToken asts.Token) (Root, error)
}

// Builder represents a root token builder
type Builder interface {
	Create() Builder
	WithContent(content string) Builder
	WithName(name string) Builder
	WithNodes(nodes Nodes) Builder
	Now() (Root, error)
}

// Root represents a root token
type Root interface {
	Content() string
	Name() string
	Nodes() Nodes
	Scan(name string) (Root, Node, Container, string, error)
}

// NodesBuilder represents a nodes builder
type NodesBuilder interface {
	Create() NodesBuilder
	WithNodes(nodes map[string]Node) NodesBuilder
	Now() (Nodes, error)
}

// Nodes represents nodes
type Nodes interface {
	All() map[string]Node
	Amount() uint
	Fetch(name string) (Node, error)
	Scan(name string) (Root, Node, Container, string, error)
	Exists(name string) bool
}

// NodeBuilder represents a node builder
type NodeBuilder interface {
	Create() NodeBuilder
	WithName(name string) NodeBuilder
	WithContainers(containers Containers) NodeBuilder
	Now() (Node, error)
}

// Node represents a container node
type Node interface {
	Name() string
	Content() string
	Containers() Containers
	Scan(name string) (Root, Node, Container, string, error)
}

// ContainersBuilder represents containers builder
type ContainersBuilder interface {
	Create() ContainersBuilder
	WithList(list []Container) ContainersBuilder
	Now() (Containers, error)
}

// Containers represents containers
type Containers interface {
	Content() string
	List() []Container
	Fetch(name string) (Container, error)
	Scan(name string) (Root, Node, Container, string, error)
}

// ContainerBuilder represents a container builder
type ContainerBuilder interface {
	Create() ContainerBuilder
	WithContent(content string) ContainerBuilder
	WithName(name string) ContainerBuilder
	WithRoot(root Root) ContainerBuilder
	Now() (Container, error)
}

// Container represents a container
type Container interface {
	Scan(name string) (Root, Node, Container, string, error)
	Content() string
	Name() string
	HasRoot() bool
	Root() Root
}
