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
	WithClear(clear string) PeerBuilder
	WithOnion(onion string) PeerBuilder
	WithPort(port uint) PeerBuilder
	Now() (Peer, error)
}

// Peer represents a peer
type Peer interface {
	Host() Host
	Port() uint
	String() string
}

// Host represents a peer host
type Host interface {
	IsClear() bool
	Clear() string
	IsOnion() bool
	Onion() string
}
