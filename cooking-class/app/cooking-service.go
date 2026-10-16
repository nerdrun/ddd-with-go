package app

import (
	"cooking-class/domain"
	"time"
)

type CookingService struct{}

func NewCookingService() *CookingService {
	return &CookingService{}
}

func (s *CookingService) AddBooking(cookingClassID string) error {
	// this should be fetched from repository, but for simplicity, we create a new one here
	cookingClass, err := domain.NewCookingClass("Italian Cooking", time.Now(), time.Now().Add(2*time.Hour), 10, 1000, "USD")
	if err != nil {
		return err
	}
	_, err = cookingClass.AddBooking("John Doe", "john.doe@example.com")
	if err != nil {
		return err
	}

	// I don't know how to pop the events here
	for _, event := range cookingClass.PopEvents() {
		event.EventName()
	}

	return nil
}
