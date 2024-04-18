package repository

import (
	model "emplyees-api/employee/entity"
	"fmt"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

// CreateEmployee creates a new employee record in the database
func (repo *EmployeeRepository) CreateEmployee(employee *model.Employee) error {
	// Implement logic to create an employee record in the database
	if err := repo.DB.Create(employee).Error; err != nil {
		return err
	}
	return nil
}

// GetAllEmployees fetches all employee records from the database
func (repo *EmployeeRepository) GetAllEmployees() ([]*model.Employee, error) {
	var employees []*model.Employee
	if err := repo.DB.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// GetEmployeeByID fetches a single employee record from the database by ID
func (repo *EmployeeRepository) GetEmployeeByID(id uint) (*model.Employee, error) {
	// Declare a variable to store the fetched employee
	var employee model.Employee

	// Query the database to fetch the employee record by ID
	if err := repo.DB.First(&employee, id).Error; err != nil {
		// If an error occurs during the query, return the error
		return nil, err
	}

	// If no error occurs, return the fetched employee
	return &employee, nil
}

// UpdateEmployee updates an existing employee record in the database
func (repo *EmployeeRepository) UpdateEmployee(employee *model.Employee) error {
	// Check if the employee with the given ID exists in the database
	existingEmployee := model.Employee{}
	if err := repo.DB.First(&existingEmployee, employee.ID).Error; err != nil {
		// If the employee does not exist, return an error
		return fmt.Errorf("employee with ID %d not found", employee.ID)
	}

	// Update the existing employee record with the new data
	if err := repo.DB.Model(&existingEmployee).Updates(employee).Error; err != nil {
		// If an error occurs while updating the record, return the error
		return err
	}

	// Return nil to indicate success
	return nil
}

// DeleteEmployee deletes an existing employee record from the database by ID
func (repo *EmployeeRepository) DeleteEmployee(id uint) error {
	// Find the employee record by ID
	employee := &model.Employee{}
	result := repo.DB.First(employee, id)
	if result.Error != nil {
		// If the record is not found, return an error
		return result.Error
	}

	// Delete the employee record
	if err := repo.DB.Delete(employee).Error; err != nil {
		// If an error occurs while deleting the record, return the error
		return err
	}

	return nil
}
