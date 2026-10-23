package cas

import (
	"encoding/binary"
	"encoding/hex"
)

// Hash is a unique 64-bit representation of content for usage in CAS.
type Hash uint64

func (h Hash) Bytes() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(h))
	return b
}

func (h Hash) HexString() string {
	return hex.EncodeToString(h.Bytes())
}
