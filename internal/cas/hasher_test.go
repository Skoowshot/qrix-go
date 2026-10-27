package cas

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/zeebo/xxh3"
)

func TestXXH3Hasher(t *testing.T) {
	hasher := XXH3Hasher()
	data := bytes.Repeat([]byte("hello world"), 64)
	expectedU128 := xxh3.Hash128(data)
	var expected Hash
	binary.BigEndian.PutUint64(expected[0:8], expectedU128.Hi)
	binary.BigEndian.PutUint64(expected[8:16], expectedU128.Lo)

	got, err := hasher(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if got != expected {
		t.Errorf("got %s, want %s", got.HexString(), expected.HexString())
	}
}
