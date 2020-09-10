package models

type Contact struct{
	ID        int       `json:"id,omitempty"`
	TeamID	  int		`json:"team_id"`
	Name      string    `json:"name"`
	Phone 	  string	`json:"phone,omitempty"`
	Email	  string	`json:"email,omitempty"`
	CreatedAt int `json:"created_at,omitempty"`
	UpdatedAt int `json:"updated_at,omitempty"`
	CustomAttributes []CustomAttribute `json:"custom_attributes,omitempty"`
}