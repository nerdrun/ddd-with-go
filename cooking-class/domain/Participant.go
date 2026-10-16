package domain

import "fmt"

type Participant struct {
	name    string
	contact string
}

func NewParticipant(name string, contact string) (*Participant, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if contact == "" {
		return nil, fmt.Errorf("contact cannot be empty")
	}

	return &Participant{
		name:    name,
		contact: contact,
	}, nil
}
