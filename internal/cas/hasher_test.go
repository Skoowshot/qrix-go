package cas

import (
	"bytes"
	"testing"

	"github.com/zeebo/xxh3"
)

func TestXXH3Hasher(t *testing.T) {
	hasher := XXH3Hasher()
	data := bytes.Repeat([]byte("hello world"), 64)
	expectedU128 := xxh3.Hash128(data)
	expected := uint128ToHash(expectedU128)

	got, err := hasher(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if got != expected {
		t.Errorf("got %s, want %s", got.HexString(), expected.HexString())
	}
}
