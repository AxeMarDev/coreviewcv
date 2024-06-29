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

func GetProjects(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	companyId, _ := c.Get("company_id")

	fmt.Println(companyId)

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
