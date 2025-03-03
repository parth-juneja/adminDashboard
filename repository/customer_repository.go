package repository

import (
	"adminDashboard/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

// Create a new customer
func (r *CustomerRepository) CreateCustomer(customer *models.Customer) error {
	return r.DB.Create(customer).Error
}

// Get a customer by ID
func (r *CustomerRepository) GetCustomerByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	err := r.DB.First(&customer, id).Error
	return &customer, err
}

// Get all customers
func (r *CustomerRepository) GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.DB.Find(&customers).Error
	return customers, err
}

// Update customer
func (r *CustomerRepository) UpdateCustomer(customer *models.Customer) error {
	return r.DB.Save(customer).Error
}

// Delete customer
func (r *CustomerRepository) DeleteCustomer(id uint) error {
	return r.DB.Delete(&models.Customer{}, id).Error
}
