package app

type Server interface {
	Start() error
	GracefulStop()
	Stop()
}
