package cas

import (
	"io"
)

type Content interface {
	io.Closer
	io.Reader
}
