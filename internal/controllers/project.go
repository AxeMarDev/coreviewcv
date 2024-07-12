package controllers

import (
	"backendcv/internal/database"
	"backendcv/internal/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetProject(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	projectID := c.Param("project_id") // This retrieves the project_id from the URL path
	fmt.Println(projectID)

	row := db.QueryRow("SELECT id, name, company_id FROM project WHERE id = ($1)  ", projectID)

	var projects models.Project

	var p models.Project
	if err := row.Scan(&p.ID, &p.Name, &p.Companyid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan project"})
		fmt.Println("Failed to scan project")
		return
	}
	projects = p

	c.IndentedJSON(http.StatusOK, projects)

}
func GetProjects(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	companyId, _ := c.Get("company_id")

	rows, err := db.Query("SELECT id, name, company_id FROM project WHERE company_id = ($1) ORDER BY id ASC ", companyId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query projects"})
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not complete request"})
		}
	}(rows)

	var projects []models.Project

	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Companyid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan person"})
			return
		}
		projects = append(projects, p)
	}
	fmt.Println(projects)

	c.IndentedJSON(http.StatusOK, projects)

}

func AddProject(c *gin.Context) {
	var newProject models.Project

	var db *sql.DB
	db = database.Db

	companyId, _ := c.Get("company_id")

	fmt.Println(companyId)

	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert newPerson into the database
	query := `INSERT INTO project (name, company_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(query, newProject.Name, companyId).Scan(&id)

	if err != nil {
		log.Printf("Error while inserting new project: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new project"})
		return
	}

	// Return the new person as JSON
	c.JSON(http.StatusCreated, newProject)
}

func DeleteProject(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	id := c.Query("id")

	// Delete all entries in project_client referencing the project
	deleteProjectClientsQuery := `DELETE FROM project_client WHERE project_id = $1`
	if _, err := db.Exec(deleteProjectClientsQuery, id); err != nil {
		log.Printf("Error while deleting project clients: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project clients"})
		return
	}

	// Execute the delete query
	query := `DELETE FROM project WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error while deleting client: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking deletion result"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No project found with the provided ID"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "project deleted successfully"})

}

func UpdateProjectName(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	type parameters struct {
		NewName string `json:"newName"`
	}

	var urlParams parameters

	projectID := c.Param("project_id")

	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&urlParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error when binding name"})
		return
	}

	query := `UPDATE project SET  name = $1 WHERE id = $2`
	result, err := db.Exec(query, urlParams.NewName, projectID)

	if err != nil {
		log.Printf("Error while deleting person: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking deletion result"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No project found with the provided ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message updated successfully"})

}
