package peers

import "fmt"

type peer struct {
	host string
	port uint
}

func createPeer(
	host string,
	port uint,
) Peer {
	out := peer{
		host: host,
		port: port,
	}

	return &out
}

// Host returns the host
func (obj *peer) Host() string {
	return obj.host
}

// Port returns the port
func (obj *peer) Port() uint {
	return obj.port
}

// String returns the string representation of the peer
func (obj *peer) String() string {
	return fmt.Sprintf("%s%d", obj.host, obj.port)
}
