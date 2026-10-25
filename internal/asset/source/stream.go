package source

import (
	"io"

	"github.com/Skoowshot/qrix-runtime/internal/asset"
)

var _ asset.Source = (*Stream)(nil)

// Stream represents an implementation of asset.Source that wraps an io.ReadCloser
// and carries a specific asset.SourceType.
type Stream struct {
	sourceType asset.SourceType
	reader     io.ReadCloser
}

// NewStream initializes and returns a new Stream instance.
func NewStream(assetType asset.SourceType, reader io.ReadCloser) *Stream {
	return &Stream{
		sourceType: assetType,
		reader:     reader,
	}
}

// Type returns the source type associated with this stream.
func (s *Stream) Type() asset.SourceType {
	return s.sourceType
}

// Read reads data from the underlying reader stream.
func (s *Stream) Read(p []byte) (n int, err error) {
	return s.reader.Read(p)
}

// Close closes the underlying reader stream.
func (s *Stream) Close() error {
	return s.reader.Close()
}
