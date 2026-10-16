package domain

import "uuid"

type BookingID uuid.UUID

type Booking struct {
	ID          BookingID
	participant *Participant
	status      string
}

func NewBooking(id uuid.UUID, participant *Participant) *Booking {
	return &Booking{
		ID:          BookingID(id),
		participant: participant,
		status:      "WAITLISTED",
	}
}

func (b *Booking) Status() string {
	return b.status
}
