package domain

import "time"

type ClassBecameFullEvent struct {
	cookingClassID CookingClassID
	occurredAt     time.Time
}

func NewClassBecameFull(cookingClassID CookingClassID) *ClassBecameFullEvent {
	return &ClassBecameFullEvent{
		cookingClassID: cookingClassID,
		occurredAt:     time.Now(),
	}
}

func (e *ClassBecameFullEvent) EventName() string {
	return e.EventName()
}
