package peers

// Builder represents a peers builder
type Builder interface {
	Create() Builder
	WithList(list []Peer) Builder
	Now() (Peers, error)
}

// Peers represents peers
type Peers interface {
	All() []Peer
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
