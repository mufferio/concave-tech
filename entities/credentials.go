package entities

import "time"

type Credential struct {
	ID           int
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}
