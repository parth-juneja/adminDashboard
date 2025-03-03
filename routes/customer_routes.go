package routes

import (
	"adminDashboard/controllers"

	"github.com/gin-gonic/gin"
)

func SetupHealthCheckRoutes(router *gin.Engine) {
	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func SetupCustomerRoutes(router *gin.Engine, customerController *controllers.CustomerController) {
	customerGroup := router.Group("/api/v1/customers")
	{
		customerGroup.POST("/", customerController.CreateCustomer)
		customerGroup.GET("/:id", customerController.GetCustomerByID)
		customerGroup.GET("/", customerController.GetAllCustomers)
		customerGroup.PUT("/:id", customerController.UpdateCustomer)
		customerGroup.DELETE("/:id", customerController.DeleteCustomer)
	}
}
