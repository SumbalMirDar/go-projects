package middleware

import (
	"emplyees-api/admin/repository" // Import admin repository package
	"net/http"
)

func AuthenticateAdmin(next http.Handler, repo repository.AdminRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse email and password from the request (assuming they're sent in the request body or headers)
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Call the repository function to retrieve admin credentials by email
		admin, err := repo.GetAdminByEmail(email)
		if err != nil {
			// Handle error (e.g., database error)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Check if admin with the provided email exists and if the password matches
		if admin != nil && admin.PasswordMatches(password) {
			// If authentication is successful, proceed to the next handler
			next.ServeHTTP(w, r)
		} else {
			// If authentication fails, return 401 Unauthorized response
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

// Example function to check if the request contains valid authentication token
func isValidToken(r *http.Request) bool {
	// Implement token validation logic here
	return true // Placeholder logic, replace with actual implementation
}
