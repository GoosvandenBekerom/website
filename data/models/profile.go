package models

import "time"

type Profile struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Picture     []byte    `json:"picture,omitempty"`
}
