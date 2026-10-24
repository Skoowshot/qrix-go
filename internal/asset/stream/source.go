package stream

import (
	"io"
	"qrix-runtime/internal/asset"
)

var _ asset.Source = (*Source)(nil)

// Source represents an implementation of asset.Source that wraps an io.ReadCloser
// and carries a specific asset.SourceType.
type Source struct {
	sourceType asset.SourceType
	reader     io.ReadCloser
}

// NewSource initializes and returns a new Source instance.
func NewSource(assetType asset.SourceType, reader io.ReadCloser) *Source {
	return &Source{
		sourceType: assetType,
		reader:     reader,
	}
}

// Type returns the source type associated with this stream.
func (s *Source) Type() asset.SourceType {
	return s.sourceType
}

// Read reads data from the underlying reader stream.
func (s *Source) Read(p []byte) (n int, err error) {
	return s.reader.Read(p)
}

// Close closes the underlying reader stream.
func (s *Source) Close() error {
	return s.reader.Close()
}
