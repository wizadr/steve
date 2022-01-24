package peers

import "fmt"

type peer struct {
	host      string
	port      uint
	delimiter string
}

func createPeer(
	host string,
	port uint,
	delimiter string,
) Peer {
	out := peer{
		host:      host,
		port:      port,
		delimiter: delimiter,
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
	return fmt.Sprintf("%s%s%d", obj.host, obj.delimiter, obj.port)
}
