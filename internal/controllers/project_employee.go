package controllers

import (
	"backendcv/internal/database"
	"backendcv/internal/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
	"net/http"
)

func DeleteEmployeeFromProject(c *gin.Context) {
	var db *sql.DB
	db = database.Db

	employeeID := c.Param("employee_id")
	projectID := c.Param("project_id")

	query := `DELETE FROM project_employee WHERE employee_id = $1 AND project_id = $2`
	result, err := db.Exec(query, employeeID, projectID)
	if err != nil {
		log.Printf("Error while deleting client: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee<->project relation where employee"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking deletion result"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No client found with the provided ID"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "client deleted successfully"})

}

func AddEmployeeToProject(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	employeeID := c.Param("employee_id")
	projectID := c.Param("project_id")

	// Insert clientID and projectID into the database
	query := `INSERT INTO project_employee (project_id, employee_id) VALUES ($1, $2)`
	_, err := db.Exec(query, projectID, employeeID)

	if err != nil {
		log.Printf("Error while inserting new relation (Employee <-> Project): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting new relation (Employee <-> Project)"})
		return
	}

	// Return success message as JSON
	c.JSON(http.StatusCreated, gin.H{"success": "Inserted new relation (Employee <-> Project)"})
}

func GetProjectEmployees(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	projectID := c.Param("project_id")

	rows, err := db.Query("SELECT  employee_id FROM project_employee WHERE project_id = ($1) ORDER BY employee_id ASC ", projectID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query employee <-> project relation"})
		return
	}

	var employeeIds []int // Assuming client_id is of type int in your database

	for rows.Next() {
		var employeeID int
		if err := rows.Scan(&employeeID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan employee_id"})
			return
		}
		employeeIds = append(employeeIds, employeeID)
	}

	query := ` SELECT id, name FROM employee  WHERE id = ANY($1) `

	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}
	defer stmt.Close()

	// Execute the query with array of client IDs
	rows, err = stmt.Query(pq.Array(employeeIds))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}
	defer rows.Close()

	// Process the rows and populate the clients slice
	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.ID, &employee.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan emplotee"})
			return
		}
		employees = append(employees, employee)
	}

	c.JSON(http.StatusOK, employees)
}
