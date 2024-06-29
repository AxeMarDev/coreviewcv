package routes

import (
	"backendcv/internal/authenticator"
	"backendcv/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// company
	router.POST("/register", controllers.Register)

	router.POST("/employee-auth", controllers.EmployeeAuth)
	router.POST("/client-auth", controllers.ClientAuth)

	router.GET("/client", authenticator.JWTAuthMiddleware(), controllers.GetClients)
	router.POST("/client", authenticator.JWTAuthMiddleware(), controllers.AddClient)
	router.DELETE("/client", authenticator.JWTAuthMiddleware(), controllers.DeleteClient)

	router.GET("/project", authenticator.JWTAuthMiddleware(), controllers.GetProjects)
	router.POST("/project", authenticator.JWTAuthMiddleware(), controllers.AddProject)

	router.GET("/employees", authenticator.JWTAuthMiddleware(), controllers.GetEmployees)
	router.POST("/employees", authenticator.JWTAuthMiddleware(), controllers.AddEmployees)

}
