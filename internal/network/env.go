package network

import "io"

// Environment receives and sends messages.
type Environment interface {
	io.Closer
}
