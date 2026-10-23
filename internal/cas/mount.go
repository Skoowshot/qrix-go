package cas

// Mount represents CAS with different storage strategies.
type Mount interface {
	Get(hash Hash) (Content, error)
	Has(hash Hash) (bool, error)
	Put(hash Hash, content Content) error
}
