package domain

type Dispatcher interface {
	Dispatch(event Event) error
}
