package main

import (
	entity "emplyees-api/employee/entity"
	"emplyees-api/employee/handler"
	"emplyees-api/employee/repository"
	"emplyees-api/employee/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Establish database connection
	dsn := "postgresql://postgres:admin@localhost:5432/mydatabase?connect_timeout=300"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Automatically migrate the schema
	err = db.AutoMigrate(&entity.Employee{})
	if err != nil {
		fmt.Println("Error migrating schema:", err)
		return
	}

	// Create an instance of EmployeeService with the database connection
	employeeService := &service.EmployeeService{Repo: &repository.EmployeeRepository{DB: db}}

	// Create a new Fiber app
	app := fiber.New()

	// Register employee handler with the employeeService
	handler.RegisterEmployeeHandler(app, employeeService)

	// Serve Swagger UI
	app.Static("/swagger", "../swagger") // Assuming your Swagger UI files are in the "swagger" directory

	// Start the server
	err = app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
