package disk

import (
	"path/filepath"
	"testing"

	"github.com/Skoowshot/qrix-runtime/internal/cas"
)

func TestFlatLayout(t *testing.T) {
	const sampleHash = cas.Hash(0x0102030405060708)
	const hexStr = "0102030405060708"

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
	const sampleHash = cas.Hash(0x0102030405060708)
	const hexStr = "0102030405060708"

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
			hash:   cas.Hash(0),
			want:   filepath.Join("00", "00", "0000000000000000"),
		},
		{
			name:   "Zero or negative levels fallback to flat",
			layout: ShardedLayout(0),
			hash:   sampleHash,
			want:   hexStr,
		},
		{
			name:   "Prefix longer than hash length falls back to flat",
			layout: ShardedLayout(10), // 20 > 16
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
