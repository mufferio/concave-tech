package entities

import "time"

type User struct {
	ID          int
	Firstname   string
	Lastname    string
	DateOfBirth time.Time
	Income      float64
	Credential  Credential
	Profile     Profile
}
