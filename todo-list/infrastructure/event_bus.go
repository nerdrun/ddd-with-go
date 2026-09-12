package infrastructure

import (
	"context"
	"todo/domain"
)

type EventHandler func(ctx context.Context, event domain.Event) error

type InMemoryEventDispatcher struct {
	handlers map[string][]EventHandler
}

func NewInMemoryEventDispatcher() *InMemoryEventDispatcher {
	return &InMemoryEventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (b *InMemoryEventDispatcher) Subscribe(eventName string, handler EventHandler) {
	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *InMemoryEventDispatcher) Dispatch(ctx context.Context, event domain.Event) error {
	handlers, exists := b.handlers[event.EventName()]
	if !exists {
		return nil
	}

	for _, handler := range handlers {
		if err := handler(ctx, event); err != nil {
			return err
		}
	}
	return nil
}
