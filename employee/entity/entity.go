package model

// Employee represents an employee entity
type Employee struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	// Add more fields as needed
}
