package cas

import (
	"bytes"
	"fmt"
	"io"
)

type Content interface {
	io.Closer
	io.Reader
}

func SilentClose(c Content) {
	_ = c.Close()
}

func StreamContent(r io.Reader) Content {
	if rc, ok := r.(io.ReadCloser); ok {
		return rc
	}

	return io.NopCloser(r)
}

func InMemoryContent(data []byte) Content {
	return io.NopCloser(bytes.NewReader(data))
}

func Buffer(content Content) (Content, error) {
	defer SilentClose(content)

	data, err := io.ReadAll(content)
	if err != nil {
		return nil, fmt.Errorf("to byte content: %w", err)
	}

	return InMemoryContent(data), nil
}
