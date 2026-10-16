package domain

import "time"

type ClassBecameFullEvent struct {
	eventName string
	cookingClassID CookingClassID
	occurredAt     time.Time
}

func NewClassBecameFull(cookingClassID CookingClassID) *ClassBecameFullEvent {
	return &ClassBecameFullEvent{
		eventName:      "ClassBecameFull",
		cookingClassID: cookingClassID,
		occurredAt:     time.Now(),
	}
}

func (e *ClassBecameFullEvent) EventName() string {
	return e.eventName
}
