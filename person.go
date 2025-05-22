// main.go
package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Person represents the JSON we’re going to serve.
type Person struct {
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	DateOfBirth time.Time `json:"date_of_birth"` // Go’s native date/time type
	Income      int       `json:"income"`
}

func main() {
	r := gin.Default()

	// Handler for GET /person
	r.GET("/person", func(c *gin.Context) {
		// Fake data
		p := Person{
			Firstname:   "John",
			Lastname:    "Citizen",
			DateOfBirth: time.Date(2001, time.March, 3, 0, 0, 0, 0, time.UTC),
			Income:      50_000,
		}
		c.JSON(200, p) // Gin takes care of the marshaling & Content-Type
	})

	// listen on :8080 (default localhost)
	r.Run(":8080")
}
