package cas

import "context"

// Mount represents CAS with different storage strategies.
type Mount interface {
	Get(ctx context.Context, hash Hash) (Content, error)
	Has(ctx context.Context, hash Hash) (bool, error)
	Put(ctx context.Context, hash Hash, content Content) error
}
