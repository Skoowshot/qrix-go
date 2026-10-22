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
			mount := NewMount(TargetPath(tempPath), 1, 2)
			if tt.setup != nil {
				tt.setup(t, mount)
			}

			content, err := mount.Get(tt.hash)
			if content != nil {
				defer content.Close()
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
	mismatchedHash := cas.Hash(54321)

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
			content:     byteContent(contentBytes),
			setup:       nil,
			expectedErr: nil,
		},
		{
			name:    "same content put twice successfully",
			hash:    validHash,
			content: byteContent(contentBytes),
			setup: func(t *testing.T, m cas.Mount) {
				if err := m.Put(validHash, byteContent(contentBytes)); err != nil {
					t.Fatalf("setup failed: %v", err)
				}
			},
			expectedErr: nil,
		},
		{
			name:        "has does not match content returns ErrHashMismatch",
			hash:        mismatchedHash,
			content:     byteContent(contentBytes),
			setup:       nil,
			expectedErr: cas.ErrHashMismatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempPath := t.TempDir()
			mount := NewMount(TargetPath(tempPath), 1, 2)
			if tt.setup != nil {
				tt.setup(t, mount)
			}

			err := mount.Put(tt.hash, tt.content)

			if errors.Is(err, tt.expectedErr) {
				t.Errorf("error = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}

func setupContentFromBytes(hash cas.Hash, contentBytes []byte) func(t *testing.T, m cas.Mount) {
	return func(t *testing.T, m cas.Mount) {
		err := m.Put(hash, byteContent(contentBytes))
		if err != nil {
			t.Fatalf("failed to setup test data: %v", err)
		}
	}
}

func validateErr(expectedErr error) func(t *testing.T, content cas.Content, err error) {
	return func(t *testing.T, content cas.Content, err error) {
		if !errors.Is(err, expectedErr) {
			t.Errorf("error = %v, expected %v", err, expectedErr)
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
			t.Errorf("data = %q, expected %q", gotBytes, expectedBytes)
			return
		}
		if len(gotBytes) != len(expectedBytes) {
			t.Errorf("size = %d, expected %d", len(gotBytes), len(expectedBytes))
			return
		}
	}
}

func byteContent(contentBytes []byte) cas.Content {
	return io.NopCloser(bytes.NewReader(contentBytes))
}
