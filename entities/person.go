// main.go
package entities

import (
	"time"
)

// Person represents the JSON we’re going to serve.
type Person struct {
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	DateOfBirth time.Time `json:"date_of_birth"` // Go’s native date/time type
	Income      int       `json:"income"`
}
