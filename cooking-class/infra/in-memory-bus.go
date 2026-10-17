package infra

type InMemoryBus struct {}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{}
}

func (b *InMemoryBus) Subscribe(eventName string, handler func(event Event) error) {
	// Implement the logic to subscribe the handler to the event name
}

func (b *InMemoryBus) Dispatch(event Event) error {
	// Implement the logic to dispatch the event to the appropriate handlers
	return nil
}