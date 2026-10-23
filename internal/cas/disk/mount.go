package disk

import (
	"path/filepath"
	"qrix-runtime/internal/cas"
)

type TargetPath string

type Mount struct {
	targetPath TargetPath
	layout     Layout
}

func NewMount(
	targetPath TargetPath,
	layout Layout,
) *Mount {
	return &Mount{
		targetPath: targetPath,
		layout:     layout,
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
	rootPath := string(m.targetPath)
	relativePath := m.layout(hash)
	return filepath.Join(rootPath, relativePath)
}
