package interfaces

type Server interface {
	Notify() <-chan error
	Shutdown() error
}
