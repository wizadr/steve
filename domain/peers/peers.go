package peers

type peers struct {
	list []Peer
}

func createPeers(
	list []Peer,
) Peers {
	out := peers{
		list: list,
	}

	return &out
}

// List returns the list of peers
func (obj *peers) List() []Peer {
	return obj.list
}
