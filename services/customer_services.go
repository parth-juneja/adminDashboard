package services

import (
	"adminDashboard/models"
	"adminDashboard/repository"
)

type CustomerService struct {
	Repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{Repo: repo}
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) error {
	return s.Repo.CreateCustomer(customer)
}

func (s *CustomerService) GetCustomerByID(id uint) (*models.Customer, error) {
	return s.Repo.GetCustomerByID(id)
}

func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
	return s.Repo.GetAllCustomers()
}

func (s *CustomerService) UpdateCustomer(customer *models.Customer) error {
	return s.Repo.UpdateCustomer(customer)
}

func (s *CustomerService) DeleteCustomer(id uint) error {
	return s.Repo.DeleteCustomer(id)
}
