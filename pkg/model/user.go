package model

import "time"

type User struct {
	ID           int       `gorm:"primary_key",json:"id,omitempty"`
	FullName     string    `json:"full_name"`
	Message      string    `json:"message"`
	RegisteredOn time.Time `json:"registered_on,omitempty"`
}
