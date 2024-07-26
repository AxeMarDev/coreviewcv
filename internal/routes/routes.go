package routes

import (
	"backendcv/internal/authenticator"
	"backendcv/internal/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {

	version := router.Group("/web/v1")
	{
		version.GET("/blog", controllers.GetBlogs)
		version.POST("/register", controllers.Register)
		version.POST("/employee-auth", controllers.EmployeeAuth)
		version.POST("/client-auth", controllers.ClientAuth)
		version.POST("/admin-auth", controllers.AdminAuth)
		version.POST("/admin-register", controllers.AdminRegister)

		// THE FOLLOWING ROUTES ARE GROUP BECAUSE THEY ARE ONLY ACCESSIBLE BY "CONTRACTORS". DENOTED BY "/a"
		{
			/*
			 * TO ACCESS AND MODIFY PROJECT RESOURCES THAT BELONG TO SOME COMPANY. CAN ONLY
			 * BE ACCESSED BY "CONTRACTOR" ACCOUNTS ( EMPLOYEE IN DB ). DENOTED BY "/project"
			 */
			version.GET("/a/project", authenticator.JWTAuthMiddleware(), controllers.GetProjects)
			version.POST("/a/project", authenticator.JWTAuthMiddleware(), controllers.AddProject)
			version.DELETE("/a/project", authenticator.JWTAuthMiddleware(), controllers.DeleteProject)

			/*
			 * TO ACCESS AND MODIFY CLIENT RESOURCES THAT BELONG TO SOME COMPANY. CAN ONLY
			 * BE ACCESSED BY "CONTRACTOR" ACCOUNTS ( EMPLOYEE IN DB ). DENOTED BY "/client"
			 */
			version.GET("/a/client", authenticator.JWTAuthMiddleware(), controllers.GetClients)
			version.POST("/a/client", authenticator.JWTAuthMiddleware(), controllers.AddClient)
			version.DELETE("/a/client", authenticator.JWTAuthMiddleware(), controllers.DeleteClient)

			/*
			 * TO ACCESS AND MODIFY EMPLOYEE RESOURCES THAT BELONG TO SOME COMPANY. CAN ONLY
			 * BE ACCESSED BY "CONTRACTOR" ACCOUNTS ( EMPLOYEE IN DB ). DENOTED BY "/employees"
			 */
			version.GET("/a/employees", authenticator.JWTAuthMiddleware(), controllers.GetEmployees)
			version.POST("/a/employees", authenticator.JWTAuthMiddleware(), controllers.AddEmployees)

			// THE FOLLOWING PERTAIN TO A PROJECT RESOURCE AND THE RESOURCES THAT ARE CONNECTED TO A PROJECT
			{
				// files are a in file table
				version.POST("/a/project/:project_id/files", authenticator.JWTAuthMiddleware(), controllers.AddFile)
				version.GET("/a/project/:project_id/files", authenticator.JWTAuthMiddleware(), controllers.GetFiles)

				version.GET("/a/project/:project_id/files/:file_id", authenticator.JWTAuthMiddleware(), controllers.GetFile)
				version.PATCH("/a/project/:project_id/name", authenticator.JWTAuthMiddleware(), controllers.UpdateProjectName)
				version.GET("/a/project/:project_id/client", authenticator.JWTAuthMiddleware(), controllers.GetProjectClients)
				version.GET("/a/project/:project_id/employee", authenticator.JWTAuthMiddleware(), controllers.GetProjectEmployees)
				version.PUT("/a/project/:project_id/client/:client_id", authenticator.JWTAuthMiddleware(), controllers.AddClientToProject)
				version.PUT("/a/project/:project_id/employee/:employee_id", authenticator.JWTAuthMiddleware(), controllers.AddEmployeeToProject)
				version.DELETE("/a/project/:project_id/employee/:employee_id", authenticator.JWTAuthMiddleware(), controllers.DeleteEmployeeFromProject)
				version.DELETE("/a/project/:project_id/client/:client_id", authenticator.JWTAuthMiddleware(), controllers.DeleteClientsFromProject)
				version.GET("/a/project/:project_id", authenticator.JWTAuthMiddleware(), controllers.GetProject)
			}

			// THE FOLLOWING PERTAIN TO A CLIENT RESOURCE AND THE RESOURCES THAT ARE CONNECT TO A CLIENT
			{
				version.GET("/a/client/:client_id", authenticator.JWTAuthMiddleware(), controllers.GetClient)
			}

		}

		router.POST("/blog", authenticator.JWTAuthMiddleware(), controllers.AddBlog)
		router.DELETE("/blog/:blog_id", authenticator.JWTAuthMiddleware(), controllers.DeleteBlog)
	}

	router.POST("/internal/login")
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "active"})
	})

	router.GET("/testing/echo/:message", func(c *gin.Context) {
		message := c.Param("message")
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

}
