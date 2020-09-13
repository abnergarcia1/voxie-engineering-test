package models

import "time"

type Team struct{
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Contacts []Contact  `json:"contacts,omitempty"`
}