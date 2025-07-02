package entities

import "time"

type Profile struct {
	ID           int
	CredentialID int
	FirstName    string
	LastName     string
	DateOfBirth  time.Time
	Income       float64
}
