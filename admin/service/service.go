package service

import (
	"emplyees-api/admin/repository"
)

type AdminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
}

// Implement admin service methods here
