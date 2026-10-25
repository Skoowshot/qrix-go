package peer

import "errors"

// ID is a unique runtime identifier of Peer.
type ID uint64

// Peer represents a connected participant in the runtime and holds
// its credentials and associated environments.
type Peer struct {
	ID           ID
	Token        Token
	Environments []Environment
}

// NewPeer constructs a new Peer instance with starting capacity
// for its environments.
func NewPeer(id ID, token Token, envCapacity int) *Peer {
	return &Peer{
		ID:           id,
		Token:        token,
		Environments: make([]Environment, 0, envCapacity),
	}
}

// Close gracefully closes all associated environments and resets the list.
// Any encountered errors are aggregated and returned via errors.Join.
func (p *Peer) Close() error {
	var errs []error

	for _, env := range p.Environments {
		if env == nil {
			continue
		}
		if err := env.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	p.Environments = nil
	return errors.Join(errs...)
}
