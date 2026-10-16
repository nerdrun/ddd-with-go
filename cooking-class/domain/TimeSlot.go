package domain

import (
	"fmt"
	"time"
)

type TimeSlot struct {
	StartTime time.Time
	EndTime   time.Time
}

func NewTimeSlot(startTime, endTime time.Time) (*TimeSlot, error) {
	if startTime.After(endTime) {
		return nil, fmt.Errorf("startTime must be before endTime")
	}
	return &TimeSlot{
		StartTime: startTime,
		EndTime:   endTime,
	}, nil
}
