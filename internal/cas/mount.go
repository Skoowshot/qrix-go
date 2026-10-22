package cas

type Mount interface {
	Get(hash Hash) (Content, error)
	Put(hash Hash, content Content) error
}
