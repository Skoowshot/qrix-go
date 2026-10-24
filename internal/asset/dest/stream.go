package dest

import (
	"io"
	"qrix-runtime/internal/asset"
)

var _ asset.Dest = (*Stream)(nil)

// Stream represents an implementation of asset.Dest that wraps an io.WriteCloser
// and carries a specific asset.DestType.
type Stream struct {
	destType asset.DestType
	writer   io.WriteCloser
}

// NewStream initializes and returns a new Stream instance.
func NewStream(destType asset.DestType, writer io.WriteCloser) *Stream {
	return &Stream{
		destType: destType,
		writer:   writer,
	}
}

// Type returns the destination type associated with this stream.
func (s *Stream) Type() asset.DestType {
	return s.destType
}

// Write writes data to the underlying writer stream.
func (s *Stream) Write(p []byte) (n int, err error) {
	return s.writer.Write(p)
}

// Close closes the underlying writer stream.
func (s *Stream) Close() error {
	return s.writer.Close()
}
