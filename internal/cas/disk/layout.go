package disk

import (
	"path/filepath"

	"github.com/Skoowshot/qrix-runtime/internal/cas"
)

const shardingLevelSize = 2

// Layout defines how the contents are aligned inside of disk.Mount.
type Layout func(hash cas.Hash) string

// FlatLayout aligns the contents flat, all files in the same directory.
func FlatLayout() Layout {
	return func(hash cas.Hash) string {
		return hash.HexString()
	}
}

// ShardedLayout shards the contents by defined levels, files split into multiple directories.
func ShardedLayout(shardingLevels int) Layout {
	return func(hash cas.Hash) string {
		hexHash := hash.HexString()
		if shardingLevels <= 0 {
			return hexHash
		}

		prefixLen := shardingLevels * shardingLevelSize
		if len(hexHash) < prefixLen {
			return hexHash
		}

		parts := make([]string, 0, shardingLevels+1)
		for i := 0; i < prefixLen; i += shardingLevelSize {
			parts = append(parts, hexHash[i:i+shardingLevelSize])
		}
		parts = append(parts, hexHash)

		return filepath.Join(parts...)
	}
}
