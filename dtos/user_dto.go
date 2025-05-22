package dtos

type UserDTO struct {
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	DateOfBirth string  `json:"date_of_birth"`
	Income      float64 `json:"income"`
}
