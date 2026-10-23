package disk

import (
	"path/filepath"
	"qrix-runtime/internal/cas"
)

const shardingLevelSize = 2

type Layout func(hash cas.Hash) string

func FlatLayout() Layout {
	return func(hash cas.Hash) string {
		return hash.HexString()
	}
}

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
