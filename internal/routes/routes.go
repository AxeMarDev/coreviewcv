package routes

import (
	"backendcv/internal/authenticator"
	"backendcv/internal/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {

	router.POST("/register", controllers.Register)
	router.POST("/employee-auth", controllers.EmployeeAuth)
	router.POST("/client-auth", controllers.ClientAuth)
	router.POST("/admin-auth", controllers.AdminAuth)

	router.POST("/admin-register", controllers.AdminRegister)

	router.POST("/internal/login")

	router.GET("/blog", controllers.GetBlogs)
	router.POST("/blog", authenticator.JWTAuthMiddleware(), controllers.AddBlog)
	router.DELETE("/blog/:blog_id", authenticator.JWTAuthMiddleware(), controllers.DeleteBlog)

	router.GET("/a/client", authenticator.JWTAuthMiddleware(), controllers.GetClients)
	router.POST("/a/client", authenticator.JWTAuthMiddleware(), controllers.AddClient)
	router.DELETE("/a/client", authenticator.JWTAuthMiddleware(), controllers.DeleteClient)

	router.GET("/a/client/:client_id", authenticator.JWTAuthMiddleware(), controllers.GetClient)

	router.PATCH("/a/project/:project_id/name", authenticator.JWTAuthMiddleware(), controllers.UpdateProjectName)
	router.GET("/a/project/:project_id/client", authenticator.JWTAuthMiddleware(), controllers.GetProjectClients)
	router.GET("/a/project/:project_id/employee", authenticator.JWTAuthMiddleware(), controllers.GetProjectEmployees)
	router.PUT("/a/project/:project_id/client/:client_id", authenticator.JWTAuthMiddleware(), controllers.AddClientToProject)
	router.PUT("/a/project/:project_id/employee/:employee_id", authenticator.JWTAuthMiddleware(), controllers.AddEmployeeToProject)
	router.DELETE("/a/project/:project_id/employee/:employee_id", authenticator.JWTAuthMiddleware(), controllers.DeleteEmployeeFromProject)
	router.DELETE("/a/project/:project_id/client/:client_id", authenticator.JWTAuthMiddleware(), controllers.DeleteClientsFromProject)

	router.GET("/a/project/:project_id", authenticator.JWTAuthMiddleware(), controllers.GetProject)
	router.GET("/a/project", authenticator.JWTAuthMiddleware(), controllers.GetProjects)
	router.POST("/a/project", authenticator.JWTAuthMiddleware(), controllers.AddProject)
	router.DELETE("/a/project", authenticator.JWTAuthMiddleware(), controllers.DeleteProject)

	router.GET("/a/employees", authenticator.JWTAuthMiddleware(), controllers.GetEmployees)
	router.POST("/a/employees", authenticator.JWTAuthMiddleware(), controllers.AddEmployees)

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "active"})
	})

	router.GET("/testing/echo/:message", func(c *gin.Context) {
		message := c.Param("message")
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

}
