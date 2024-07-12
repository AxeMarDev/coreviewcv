package routes

import (
	"backendcv/internal/authenticator"
	"backendcv/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	/*

		routes that start with /c/ are accessible by clients, routes that start with /a/ are accessible by admins

	*/
	router.POST("/register", controllers.Register)

	router.POST("/employee-auth", controllers.EmployeeAuth)
	router.POST("/client-auth", controllers.ClientAuth)

	router.GET("/a/client", authenticator.JWTAuthMiddleware(), controllers.GetClients)
	router.POST("/a/client", authenticator.JWTAuthMiddleware(), controllers.AddClient)
	router.DELETE("/a/client", authenticator.JWTAuthMiddleware(), controllers.DeleteClient)

	router.GET("/a/project/:project_id/client", authenticator.JWTAuthMiddleware(), controllers.GetProjectClients)
	router.GET("/a/project/:project_id/employee", authenticator.JWTAuthMiddleware(), controllers.GetProjectEmployees)
	router.PUT("/a/project/:project_id/client/:client_id", authenticator.JWTAuthMiddleware(), controllers.AddClientToProject)
	router.PUT("/a/project/:project_id/employee/:employee_id", authenticator.JWTAuthMiddleware(), controllers.AddEmployeeToProject)
	router.DELETE("/a/project/:project_id/employee/:employee_id", authenticator.JWTAuthMiddleware(), controllers.DeleteEmployeeFromProject)

	router.GET("/a/project/:project_id", authenticator.JWTAuthMiddleware(), controllers.GetProject)
	router.GET("/a/project", authenticator.JWTAuthMiddleware(), controllers.GetProjects)
	router.POST("/a/project", authenticator.JWTAuthMiddleware(), controllers.AddProject)
	router.DELETE("/a/project", authenticator.JWTAuthMiddleware(), controllers.DeleteProject)

	router.GET("/a/employees", authenticator.JWTAuthMiddleware(), controllers.GetEmployees)
	router.POST("/a/employees", authenticator.JWTAuthMiddleware(), controllers.AddEmployees)

}
