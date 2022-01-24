package peers

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewPeerBuilder returns a peer builder
func NewPeerBuilder(delimiter string) PeerBuilder {
	return createPeerBuilder(delimiter)
}

// Builder represents a peers builder
type Builder interface {
	Create() Builder
	WithList(list []Peer) Builder
	Now() (Peers, error)
}

// Peers represents peers
type Peers interface {
	List() []Peer
}

// PeerBuilder represents a peer builder
type PeerBuilder interface {
	Create() PeerBuilder
	WithHost(host string) PeerBuilder
	WithPort(port uint) PeerBuilder
	WithString(str string) PeerBuilder
	Now() (Peer, error)
}

// Peer represents a peer
type Peer interface {
	Host() string
	Port() uint
	String() string
}
