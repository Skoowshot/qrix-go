package cas

import (
	"encoding/binary"
	"io"

	"github.com/zeebo/xxh3"
)

const streamBufSize = 64 * 1024

type Hasher func(r io.Reader) (Hash, error)

func XXH3Hasher() Hasher {
	return func(r io.Reader) (Hash, error) {
		h := xxh3.New()
		buf := make([]byte, streamBufSize)
		if _, err := io.CopyBuffer(h, r, buf); err != nil {
			return Hash{}, err
		}

		u128 := h.Sum128()
		return uint128ToHash(u128), nil
	}
}

func uint128ToHash(u128 xxh3.Uint128) Hash {
	var out Hash
	binary.BigEndian.PutUint64(out[0:8], u128.Hi)
	binary.BigEndian.PutUint64(out[8:16], u128.Lo)
	return out
}
