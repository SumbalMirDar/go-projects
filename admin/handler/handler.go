package handler

import (
	repository "emplyees-api/admin/repository"
	middleware "emplyees-api/pkg/middleware"
	"net/http"
)

var yourAdminRepository repository.AdminRepository

func YourAdminHandler(w http.ResponseWriter, r *http.Request) {
	// Your admin handler logic here

	// Define the next handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler or perform additional logic
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Admin Handler called successfully"))
	})

	// Call AuthenticateAdmin middleware passing the next handler and the repository
	middleware.AuthenticateAdmin(nextHandler, yourAdminRepository)
}
