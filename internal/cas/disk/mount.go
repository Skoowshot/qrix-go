package disk

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"qrix-runtime/internal/cas"
)

const standardDirPerms os.FileMode = 0755
const putTemporaryFilePattern = "put-*"

type TargetPath string

// Mount is a physical disk implementation of cas.Mount, laying out files by given Layout at the given TargetPath
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

// Get retrieves the content from the storage by the given hash.
// Opens the file and returns the content as cas.StreamContent.
// The returned content should be closed after usage.
func (m *Mount) Get(hash cas.Hash) (cas.Content, error) {
	contentFilepath := m.contentPath(hash)
	reader, err := os.Open(contentFilepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("%w: %v", cas.ErrNotFound, err)
		}

		return nil, fmt.Errorf("open content file: %w", err)
	}

	return cas.StreamContent(reader), nil
}

// Has checks if the content exists by the given hash.
func (m *Mount) Has(hash cas.Hash) (bool, error) {
	contentFilepath := m.contentPath(hash)

	_, err := os.Stat(contentFilepath)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, fmt.Errorf("stat content file: %w", err)
}

// Put saves the content to the storage by the given hash.
// Uses atomic write using a temporary file, preventing data corruption.
func (m *Mount) Put(hash cas.Hash, content cas.Content) error {
	defer cas.SilentClose(content)

	contentFilepath := m.contentPath(hash)
	exists, err := m.Has(hash)
	if err != nil {
		return fmt.Errorf("check if content exists: %w", err)
	}
	if exists {
		return nil
	}

	dir := filepath.Dir(contentFilepath)
	if err := os.MkdirAll(dir, standardDirPerms); err != nil {
		return fmt.Errorf("create sharded directories: %w", err)
	}

	tmpFile, err := os.CreateTemp(dir, putTemporaryFilePattern)
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}

	tmpName := tmpFile.Name()

	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpName)
	}()

	if _, err := io.Copy(tmpFile, content); err != nil {
		return fmt.Errorf("copy content to temp: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("close temp file: %w", err)
	}

	if err := os.Rename(tmpName, contentFilepath); err != nil {
		return fmt.Errorf("rename temp file to final: %w", err)
	}

	return nil
}

func (m *Mount) contentPath(hash cas.Hash) string {
	rootPath := string(m.targetPath)
	relativePath := m.layout(hash)
	return filepath.Join(rootPath, relativePath)
}
