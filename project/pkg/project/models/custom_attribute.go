package models

type CustomAttribute struct{
	ID       	  int       `json:"id"`
	ContactID	  int		`json:"contact_id"`
	Key     	  string    `json:"key"`
	Value 	  	  string	`json:"value"`
}