package infra

type EventBus struct {}

func NewEventBus() *EventBus {
	return &EventBus{}
}

func (b *EventBus) Dispatch(event Event, handler EventHandler) error {
	// Do slack here for instance?
	handler.Handle(event.EventName())
	return nil
}
