package models

import "time"

type Contact struct{
	ID        int       `json:"id,omitempty"`
	TeamID	  int		`json:"team_id"`
	Name      string    `json:"name"`
	Phone 	  string	`json:"phone"`
	Email	  string	`json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CustomAttributes []CustomAttribute `json:"custom_attributes,omitempty"`
}