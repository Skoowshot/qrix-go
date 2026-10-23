package disk

import (
	"bytes"
	"errors"
	"io"
	"qrix-runtime/internal/cas"
	"testing"
)

func TestMount_Get(t *testing.T) {
	tests := []struct {
		name     string
		hash     cas.Hash
		setup    func(t *testing.T, m cas.Mount)
		validate func(t *testing.T, content cas.Content, err error)
	}{
		{
			name:     "existing content is returned successfully",
			hash:     cas.Hash(12345),
			setup:    setupContentFromBytes(cas.Hash(12345), []byte("hello world")),
			validate: validateContentBytes([]byte("hello world")),
		},
		{
			name:     "missing content returns ErrNotFound",
			hash:     cas.Hash(12345),
			setup:    nil,
			validate: validateErr(cas.ErrNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempPath := t.TempDir()
			mount := NewMount(TargetPath(tempPath), FlatLayout())
			if tt.setup != nil {
				tt.setup(t, mount)
			}

			content, err := mount.Get(tt.hash)
			if content != nil {
				defer cas.SilentClose(content)
			}

			if tt.validate == nil {
				return
			}

			tt.validate(t, content, err)
		})
	}
}

func TestMount_Put(t *testing.T) {
	contentBytes := []byte("hello world")
	validHash := cas.Hash(12345)

	tests := []struct {
		name        string
		hash        cas.Hash
		content     cas.Content
		setup       func(t *testing.T, m cas.Mount)
		expectedErr error
	}{
		{
			name:        "new content put successfully",
			hash:        validHash,
			content:     cas.InMemoryContent(contentBytes),
			setup:       nil,
			expectedErr: nil,
		},
		{
			name:    "same content put twice successfully",
			hash:    validHash,
			content: cas.InMemoryContent(contentBytes),
			setup: func(t *testing.T, m cas.Mount) {
				if err := m.Put(validHash, cas.InMemoryContent(contentBytes)); err != nil {
					t.Fatalf("setup failed: %v", err)
				}
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempPath := t.TempDir()
			mount := NewMount(TargetPath(tempPath), FlatLayout())
			if tt.setup != nil {
				tt.setup(t, mount)
			}

			err := mount.Put(tt.hash, tt.content)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("error = %v, want %v", err, tt.expectedErr)
			}
		})
	}
}

func setupContentFromBytes(hash cas.Hash, contentBytes []byte) func(t *testing.T, m cas.Mount) {
	return func(t *testing.T, m cas.Mount) {
		err := m.Put(hash, cas.InMemoryContent(contentBytes))
		if err != nil {
			t.Fatalf("failed to setup test data: %v", err)
		}
	}
}

func validateErr(expectedErr error) func(t *testing.T, content cas.Content, err error) {
	return func(t *testing.T, content cas.Content, err error) {
		if !errors.Is(err, expectedErr) {
			t.Errorf("error = %v, want %v", err, expectedErr)
		}
	}
}

func validateContentBytes(expectedBytes []byte) func(t *testing.T, content cas.Content, err error) {
	return func(t *testing.T, content cas.Content, err error) {
		gotBytes, err := io.ReadAll(content)

		if err != nil {
			t.Fatalf("failed to validate content bytes: %v", err)
			return
		}
		if !bytes.Equal(gotBytes, expectedBytes) {
			t.Errorf("data = %q, want %q", gotBytes, expectedBytes)
			return
		}
		if len(gotBytes) != len(expectedBytes) {
			t.Errorf("size = %d, want %d", len(gotBytes), len(expectedBytes))
			return
		}
	}
}
