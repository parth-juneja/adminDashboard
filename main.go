package main

import (
	"adminDashboard/config"
	"adminDashboard/controllers"
	"adminDashboard/repository"
	"adminDashboard/routes"
	"adminDashboard/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Default().Println("INFO: Starting the application...")
	config.ConnectDatabase()

	log.Default().Println("INFO: Database connected successfully...")

	customerRepo := repository.NewCustomerRepository(config.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	router := gin.Default()
	routes.SetupHealthCheckRoutes(router)
	routes.SetupCustomerRoutes(router, customerController)

	router.Run(":8080")
}
