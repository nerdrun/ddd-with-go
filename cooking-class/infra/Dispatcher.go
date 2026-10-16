package infra

type Event interface {
	Dispatch() error
}
