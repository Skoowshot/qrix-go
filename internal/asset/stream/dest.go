package stream

import (
	"io"
	"qrix-runtime/internal/asset"
)

var _ asset.Dest = (*Dest)(nil)

// Dest represents an implementation of asset.Dest that wraps an io.WriteCloser
// and carries a specific asset.DestType.
type Dest struct {
	destType asset.DestType
	writer   io.WriteCloser
}

// NewDest initializes and returns a new Dest instance.
func NewDest(destType asset.DestType, writer io.WriteCloser) *Dest {
	return &Dest{
		destType: destType,
		writer:   writer,
	}
}

// Type returns the destination type associated with this stream.
func (s *Dest) Type() asset.DestType {
	return s.destType
}

// Write writes data to the underlying writer stream.
func (s *Dest) Write(p []byte) (n int, err error) {
	return s.writer.Write(p)
}

// Close closes the underlying writer stream.
func (s *Dest) Close() error {
	return s.writer.Close()
}
