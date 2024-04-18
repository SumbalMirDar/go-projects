package model

// Employee represents an employee entity
type Admin struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password_hash string `json:"Password_hash"`
	// Add more fields as needed
}
