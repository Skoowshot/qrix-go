package cas

import (
	"encoding/binary"
	"encoding/hex"
)

type Hash uint64

func (h Hash) Bytes() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(h))
	return b
}

func (h Hash) HexString() string {
	return hex.EncodeToString(h.Bytes())
}
