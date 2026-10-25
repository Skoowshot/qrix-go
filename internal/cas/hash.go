package cas

import (
	"encoding/hex"
)

// Hash is a unique 128-bit representation of content for usage in CAS.
type Hash [16]byte

func (h Hash) Bytes() []byte {
	return h[:]
}
func (h Hash) HexString() string {
	return hex.EncodeToString(h.Bytes())
}
