package domain

import (
	"fmt"
	"time"
	"uuid"
)

type CookingClassID uuid.UUID

type CookingClass struct {
	ID         CookingClassID
	title      string
	timeSlot   *TimeSlot
	maxSeat    int
	price      *Price
	bookings   []*Booking
	fullEvents []*ClassBecameFullEvent
}

func NewCookingClass(title string, startTime, endTime time.Time, maxSeat int, amount int, currency string) (*CookingClass, error) {
	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}
	if maxSeat <= 0 {
		return nil, fmt.Errorf("maxSeat must be greater than 0")
	}
	// I've added these lines for timeslot because outsider shouldn't create insider's objects.
	timeSlot, err := NewTimeSlot(startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("invalid time slot: %v", err)
	}
	// the same reason above
	price, err := NewPrice(amount, currency)
	if err != nil {
		return nil, fmt.Errorf("invalid price: %v", err)
	}

	return &CookingClass{
		ID:       CookingClassID(uuid.New()),
		title:    title,
		timeSlot: timeSlot,
		maxSeat:  maxSeat,
		price:    price,
		bookings: []*Booking{},
	}, nil
}

func NewCookingClassID(id string) (CookingClassID, error) {
	if id == "" {
		return CookingClassID(uuid.New()), fmt.Errorf("id cannot be empty")
	}
	parsed, err := uuid.Parse(id)
	if err != nil {
		return CookingClassID(uuid.New()), fmt.Errorf("invalid id: %v", err)
	}
	return CookingClassID(parsed), nil
}

func (c *CookingClass) AddBooking(name, contact string) (*Booking, error) {
	participant, err := NewParticipant(name, contact)
	if err != nil {
		return nil, err
	}

	confirmedBookings := make([]*Booking, 0, len(c.bookings))
	for _, b := range c.bookings {
		if b.status == "CONFIRMED" {
			confirmedBookings = append(confirmedBookings, b)
		}
	}

	if isAlreadyParticipated(participant, c.bookings) {
		return nil, fmt.Errorf("participant is already registered for this cooking class")
	}

	// the id should be uuid therefore there is no need to check for duplicates
	booking := NewBooking(uuid.New(), participant)

	if len(confirmedBookings) < c.maxSeat {
		booking.status = "CONFIRMED"
		confirmedBookings = append(confirmedBookings, booking)
	}
	c.bookings = append(c.bookings, booking)

	if len(confirmedBookings) == c.maxSeat {
		fullEvent := NewClassBecameFull(c.ID)
		c.fullEvents = append(c.fullEvents, fullEvent)
	}
	return booking, nil
}

func (c *CookingClass) PopEvents() []*ClassBecameFullEvent {
	fullEvents := append([]*ClassBecameFullEvent{}, c.fullEvents...)
	c.fullEvents = []*ClassBecameFullEvent{}
	return fullEvents
}

func isAlreadyParticipated(participant *Participant, bookings []*Booking) bool {
	for _, b := range bookings {
		// Go compares struct values by comparing each field. If all fields are equal, the structs are considered equal.
		// But if we treat Participant as a entity and we want to compare them by their identity (e.g., a unique ID).
		if b.participant == participant {
			return true
		}
	}
	return false
}
