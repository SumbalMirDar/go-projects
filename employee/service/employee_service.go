package service

import (
	entity "emplyees-api/employee/entity"
	repository "emplyees-api/employee/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

// CreateEmployee creates a new employee
func (service *EmployeeService) CreateEmployee(ctx *fiber.Ctx) error {
	// Parse request body to extract employee data
	var employee entity.Employee // Use the Employee type from the entity package
	if err := ctx.BodyParser(&employee); err != nil {
		return err
	}

	// Call repository method to save employee data to the database
	if err := service.Repo.CreateEmployee(&employee); err != nil {
		// Handle error
		return err
	}

	// Return success response
	return ctx.SendString("Employee created successfully")
}

// GetAllEmployees fetches all employees
func (service *EmployeeService) GetAllEmployees(ctx *fiber.Ctx) error {
	// Call the GetAllEmployees method of the repository
	employees, err := service.Repo.GetAllEmployees()
	if err != nil {
		// Handle the error
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching employees")

	}

	// Return the employees as JSON response
	return ctx.JSON(employees)
}

// GetEmployeeByID fetches a single employee by ID
func (service *EmployeeService) GetEmployeeByID(ctx *fiber.Ctx) error {
	// Parse the employee ID from the request params
	id := ctx.Params("id")

	// Convert the ID string to uint
	employeeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// If an error occurs while parsing the ID, return a bad request response
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid employee ID")
	}

	// Call the GetEmployeeByID method of the repository to fetch the employee by ID
	employee, err := service.Repo.GetEmployeeByID(uint(employeeID))
	if err != nil {
		// If an error occurs while fetching the employee, return an employee server error response
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching employees")
	}

	// If the employee is not found, return a not found response
	if employee == nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Employee not found")
	}

	// Return the employee as a JSON response
	return ctx.JSON(employee)
}

// UpdateEmployee updates an existing employee
func (service *EmployeeService) UpdateEmployee(ctx *fiber.Ctx) error {
	// Parse the employee ID from the request params
	id := ctx.Params("id")

	// Convert the ID string to uint
	employeeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// If an error occurs while parsing the ID, return a bad request response
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid employee ID")
	}

	// Fetch the existing employee record from the database by ID
	existingEmployee, err := service.Repo.GetEmployeeByID(uint(employeeID))
	if err != nil {
		// If an error occurs while fetching the employee, return an employee server error response
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching employees")
	}

	// Parse the updated employee data from the request body
	var updatedEmployee entity.Employee
	if err := ctx.BodyParser(&updatedEmployee); err != nil {
		// If an error occurs while parsing the request body, return a bad request response
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	// Update only the fields that are provided in the request body
	if updatedEmployee.Name != "" {
		existingEmployee.Name = updatedEmployee.Name
	}
	if updatedEmployee.Email != "" {
		existingEmployee.Email = updatedEmployee.Email
	}
	// Update other fields as needed...

	// Save the changes back to the database
	if err := service.Repo.UpdateEmployee(existingEmployee); err != nil {
		// If an error occurs while updating the employee, return an employee server error response
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching employees")
	}

	// Return the updated employee as a JSON response
	return ctx.JSON(existingEmployee)
}

// DeleteEmployee deletes an existing employee
func (service *EmployeeService) DeleteEmployee(ctx *fiber.Ctx) error {
	// Parse the employee ID from the request params
	id := ctx.Params("id")

	// Convert the ID string to uint
	employeeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// If an error occurs while parsing the ID, return a bad request response
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid employee ID")
	}

	// Call the DeleteEmployee method of the repository to delete the employee by ID
	if err := service.Repo.DeleteEmployee(uint(employeeID)); err != nil {
		// If an error occurs while deleting the employee, return an employee server error response
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching employees")
	}

	// Return a success response
	return ctx.SendString("Employee deleted successfully")
}
