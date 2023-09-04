package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guerraglucas/ginAPIwithGoroutine/handlers"
)

// InitRoutes initializes the routes for the application
func InitRoutes() *gin.Engine {
	router := gin.Default()

	// Add routes here
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Up And Running!",
		})
	})

	router.GET("/get-something", handlers.GetSomething)

	return router
}
