package controllers

import (
	"backendcv/internal/database"
	"backendcv/internal/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
	"net/http"
)

func DeleteClientsFromProject(c *gin.Context) {
	var db *sql.DB
	db = database.Db

	clientID := c.Param("client_id")
	projectID := c.Param("project_id")

	query := `DELETE FROM project_client WHERE client_id = $1 AND project_id = $2`
	result, err := db.Exec(query, clientID, projectID)
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

func AddClientToProject(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	clientID := c.Param("client_id")
	projectID := c.Param("project_id")

	// Insert clientID and projectID into the database
	query := `INSERT INTO project_client (project_id, client_id) VALUES ($1, $2)`
	_, err := db.Exec(query, projectID, clientID)

	if err != nil {
		log.Printf("Error while inserting new relation (Client <-> Project): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting new relation (Client <-> Project)"})
		return
	}

	// Return success message as JSON
	c.JSON(http.StatusCreated, gin.H{"success": "Inserted new relation (Client <-> Project)"})
}

func GetProjectClients(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	fmt.Println("1")

	projectID := c.Param("project_id")

	// how can i get the table clients using client_id from this query??
	rows, err := db.Query("SELECT client_id FROM project_client WHERE project_id = ($1) ORDER BY client_id ASC ", projectID)

	fmt.Println("2")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query clients-projects relation"})
		return
	}

	var clientIDs []int // Assuming client_id is of type int in your database

	for rows.Next() {
		var clientID int
		if err := rows.Scan(&clientID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan client_id"})
			return
		}
		clientIDs = append(clientIDs, clientID)
	}

	// Prepare the query to fetch client details based on client IDs
	query := ` SELECT id, name FROM client  WHERE id = ANY($1) `

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}
	defer stmt.Close()

	// Execute the query with array of client IDs
	rows, err = stmt.Query(pq.Array(clientIDs))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	defer rows.Close()

	// Process the rows and populate the clients slice
	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan client"})
			return
		}
		clients = append(clients, client)
	}

	c.JSON(http.StatusOK, clients)
}
