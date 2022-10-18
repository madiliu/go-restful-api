package main

import (
	"github.com/gin-gonic/gin"
	"user_assignment/api"
	"user_assignment/database"
	"user_assignment/middleware"
)

func main() {

	// Set up postgres connection
	postgres := database.SetUpPostgres()
	// Instantiates the user service
	queries := database.New(postgres.DB)
	userService := api.NewService(queries)

	router := gin.Default()
	// Instantiate the logger
	router.Use(middleware.LoggerToFile())
	// Register service handlers to the router
	userService.RegisterHandlers(router)

	// Start the router
	err := router.Run()
	if err != nil {
		return
	}
}
