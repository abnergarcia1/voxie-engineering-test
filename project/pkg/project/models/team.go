package models

type Team struct{
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Contacts []Contact  `json:"contacts,omitempty"`
}