package peers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type peerBuilder struct {
	delimiter string
	host      string
	port      *uint
	str       string
}

func createPeerBuilder(
	delimiter string,
) PeerBuilder {
	out := peerBuilder{
		delimiter: delimiter,
		host:      "",
		port:      nil,
		str:       "",
	}

	return &out
}

// Create initializes the builder
func (app *peerBuilder) Create() PeerBuilder {
	return createPeerBuilder(
		app.delimiter,
	)
}

// WithHost adds a host to the builder
func (app *peerBuilder) WithHost(host string) PeerBuilder {
	app.host = host
	return app
}

// WithPort adds a port to the builder
func (app *peerBuilder) WithPort(port uint) PeerBuilder {
	app.port = &port
	return app
}

// WithString adds a string to the builder
func (app *peerBuilder) WithString(str string) PeerBuilder {
	app.str = str
	return app
}

// Now builds a new Peer instance
func (app *peerBuilder) Now() (Peer, error) {
	if app.str != "" {
		sections := strings.Split(app.str, app.delimiter)
		if len(sections) != 2 {
			str := fmt.Sprintf("the delimiter (%s) was expected to be contained once in the peer string: '%s'", app.delimiter, app.str)
			return nil, errors.New(str)
		}

		portInt, err := strconv.Atoi(sections[1])
		if err != nil {
			return nil, err
		}

		port := uint(portInt)
		app.port = &port
		app.host = sections[0]
	}

	if app.host == "" {
		return nil, errors.New("the host is mandatory in order to build a Peer instance")
	}

	if app.port == nil {
		return nil, errors.New("the port is mandatory in order to build a Peer instance")
	}

	return createPeer(app.host, *app.port, app.delimiter), nil
}
