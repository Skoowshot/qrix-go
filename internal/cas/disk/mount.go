package disk

import "qrix-runtime/internal/cas"

type Mount struct{}

func NewMount() *Mount {
	return &Mount{}
}

func (m *Mount) Get(hash cas.Hash) (cas.Content, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mount) Put(hash cas.Hash, content cas.Content) error {
	//TODO implement me
	panic("implement me")
}
