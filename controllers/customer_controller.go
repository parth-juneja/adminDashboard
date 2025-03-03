package controllers

import (
	"net/http"
	"strconv"
	"time"

	"adminDashboard/models"
	"adminDashboard/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Service *services.CustomerService
}

func NewCustomerController(service *services.CustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

// Create Customer
func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer.StartDate = time.Now()
	if err := c.Service.CreateCustomer(&customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, customer)
}

// Get Customer by ID
func (c *CustomerController) GetCustomerByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	customer, err := c.Service.GetCustomerByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

// Get All Customers
func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {
	customers, err := c.Service.GetAllCustomers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

// Update Customer
func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer.ID = uint(id)
	if err := c.Service.UpdateCustomer(&customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

// Delete Customer
func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	if err := c.Service.DeleteCustomer(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
