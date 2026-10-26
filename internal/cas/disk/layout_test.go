package disk

import (
	"path/filepath"
	"testing"

	"github.com/Skoowshot/qrix-runtime/internal/cas"
)

func TestFlatLayout(t *testing.T) {
	sampleHash := cas.Hash([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	hexStr := sampleHash.HexString()

	tests := []struct {
		name   string
		layout Layout
		hash   cas.Hash
		want   string
	}{
		{
			name:   "FlatLayout returns raw hex",
			layout: FlatLayout(),
			hash:   sampleHash,
			want:   hexStr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.layout(tt.hash)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestShardedLayout(t *testing.T) {
	sampleHash := cas.Hash([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	hexStr := sampleHash.HexString()

	tests := []struct {
		name   string
		layout Layout
		hash   cas.Hash
		want   string
	}{
		{
			name:   "ShardedLayout 2 levels",
			layout: ShardedLayout(2),
			hash:   sampleHash,
			want:   filepath.Join("01", "02", hexStr),
		},
		{
			name:   "Zero hash keeps leading zeros",
			layout: ShardedLayout(2),
			hash:   cas.Hash([16]byte{}),
			want:   filepath.Join("00", "00", "00000000000000000000000000000000"),
		},
		{
			name:   "Zero or negative levels fallback to flat",
			layout: ShardedLayout(0),
			hash:   sampleHash,
			want:   hexStr,
		},
		{
			name:   "Prefix longer than hash length falls back to flat",
			layout: ShardedLayout(20), // 40 > 32
			hash:   sampleHash,
			want:   hexStr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.layout(tt.hash)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
