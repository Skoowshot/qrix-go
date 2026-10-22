package disk

import (
	"path/filepath"
	"qrix-runtime/internal/cas"
)

type TargetPath string

type Mount struct {
	targetPath        TargetPath
	shardingLevels    int
	shardingLevelSize int
}

func NewMount(
	targetPath TargetPath,
	shardingLevels int,
	shardingLevelSize int,
) *Mount {
	return &Mount{
		targetPath:        targetPath,
		shardingLevels:    shardingLevels,
		shardingLevelSize: shardingLevelSize,
	}
}

func (m *Mount) Get(hash cas.Hash) (cas.Content, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mount) Put(hash cas.Hash, content cas.Content) error {
	//TODO implement me
	panic("implement me")
}

func (m *Mount) contentPath(hash cas.Hash) string {
	relativePath := shardedPath(hash, m.shardingLevels, m.shardingLevelSize)
	return filepath.Join(string(m.targetPath), relativePath)
}

func shardedPath(hash cas.Hash, shardingLevels int, levelSize int) string {
	hexHash := hash.HexString()
	if shardingLevels <= 0 || levelSize <= 0 {
		return hexHash
	}

	prefixLen := shardingLevels * levelSize
	if len(hexHash) < prefixLen {
		return hexHash
	}

	parts := make([]string, 0, shardingLevels+1)
	for i := 0; i < prefixLen; i += levelSize {
		parts = append(parts, hexHash[i:i+levelSize])
	}
	parts = append(parts, hexHash)

	return filepath.Join(parts...)
}
