package cas

import (
	"bytes"
	"fmt"
	"io"
)

// Content represents data inside of CAS.
// It can be read as io.Reader and requires to be closed as being io.Closer.
type Content interface {
	io.Closer
	io.Reader
}

// SilentClose closes the Content and ignores the error.
func SilentClose(c Content) {
	_ = c.Close()
}

// StreamContent creates a new streaming Content from the given io.Reader.
func StreamContent(r io.Reader) Content {
	if rc, ok := r.(io.ReadCloser); ok {
		return rc
	}

	return io.NopCloser(r)
}

// InMemoryContent creates a new in-memory Content from the given byte data.
func InMemoryContent(data []byte) Content {
	return io.NopCloser(bytes.NewReader(data))
}

// Buffer converts any Content to InMemoryContent, buffering its data for fast usage.
func Buffer(content Content) (Content, error) {
	defer SilentClose(content)

	data, err := io.ReadAll(content)
	if err != nil {
		return nil, fmt.Errorf("to byte content: %w", err)
	}

	return InMemoryContent(data), nil
}
