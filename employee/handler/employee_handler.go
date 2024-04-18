package handler

import (
	"emplyees-api/employee/service"

	"github.com/gofiber/fiber/v2"
)

// RegisterEmployeeHandler registers routes for employee CRUD operations
func RegisterEmployeeHandler(app *fiber.App, employeeService *service.EmployeeService) {
	// Define routes for creating a new employee
	app.Post("/employees", employeeService.CreateEmployee)

	// Define routes for fetching all employees
	app.Get("/employees", employeeService.GetAllEmployees)

	// Define routes for fetching a single employee by ID
	app.Get("/employees/:id", employeeService.GetEmployeeByID)

	// Define routes for updating an employee by ID
	app.Put("/employees/:id", employeeService.UpdateEmployee)

	// Define routes for deleting an employee by ID
	app.Delete("/employees/:id", employeeService.DeleteEmployee)
	//app.Get("/swagger/*", fiberSwagger.Handler)
}
