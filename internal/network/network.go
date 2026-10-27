package network

// ID is a unique runtime identifier of Network.
type ID uint64

// Network handles peer storage and communication.
type Network struct {
	peers map[PeerID]Peer
}

// New creates a new Network.
func New() *Network {
	return &Network{
		peers: make(map[PeerID]Peer),
	}
}
